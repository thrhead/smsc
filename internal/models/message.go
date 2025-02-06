package models

import (
	"time"
)

// MessageStatus represents the current status of a message
type MessageStatus string

const (
	StatusPending    MessageStatus = "pending"
	StatusSent       MessageStatus = "sent"
	StatusDelivered  MessageStatus = "delivered"
	StatusFailed     MessageStatus = "failed"
	StatusExpired    MessageStatus = "expired"
	StatusRejected   MessageStatus = "rejected"
	StatusScheduled  MessageStatus = "scheduled"
)

// Message represents an SMS message in the system
type Message struct {
	ID              int64         `json:"id" db:"id"`
	Sender          string        `json:"sender" db:"sender"`
	Recipient       string        `json:"recipient" db:"recipient"`
	Content         string        `json:"content" db:"content"`
	Status          MessageStatus `json:"status" db:"status"`
	Priority        int           `json:"priority" db:"priority"`
	ValidityPeriod  time.Duration `json:"validity_period" db:"validity_period"`
	ScheduledTime   *time.Time    `json:"scheduled_time,omitempty" db:"scheduled_time"`
	CreatedAt       time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" db:"updated_at"`
	SentAt          *time.Time    `json:"sent_at,omitempty" db:"sent_at"`
	DeliveredAt     *time.Time    `json:"delivered_at,omitempty" db:"delivered_at"`
	OperatorID      string        `json:"operator_id" db:"operator_id"`
	MessageID       string        `json:"message_id" db:"message_id"`
	RetryCount      int           `json:"retry_count" db:"retry_count"`
	LastError       string        `json:"last_error" db:"last_error"`
	ClientID        string        `json:"client_id" db:"client_id"`
	CampaignID      *string       `json:"campaign_id,omitempty" db:"campaign_id"`
	DeliveryReport  string        `json:"delivery_report" db:"delivery_report"`
	Encoding        string        `json:"encoding" db:"encoding"`
	ProtocolID      int           `json:"protocol_id" db:"protocol_id"`
	ESMClass        int           `json:"esm_class" db:"esm_class"`
	DataCoding      int           `json:"data_coding" db:"data_coding"`
	SourceTON       int           `json:"source_ton" db:"source_ton"`
	SourceNPI       int           `json:"source_npi" db:"source_npi"`
	DestinationTON  int           `json:"destination_ton" db:"destination_ton"`
	DestinationNPI  int           `json:"destination_npi" db:"destination_npi"`
	ServiceType     string        `json:"service_type" db:"service_type"`
	BillingInfo     string        `json:"billing_info" db:"billing_info"`
	Cost            float64       `json:"cost" db:"cost"`
}

// NewMessage creates a new Message with default values
func NewMessage(sender, recipient, content string) *Message {
	now := time.Now()
	return &Message{
		Sender:         sender,
		Recipient:      recipient,
		Content:        content,
		Status:         StatusPending,
		Priority:       1,
		ValidityPeriod: 24 * time.Hour,
		CreatedAt:      now,
		UpdatedAt:      now,
		RetryCount:     0,
		Encoding:       "GSM",
		DataCoding:     0,
		SourceTON:      1,
		SourceNPI:      1,
		DestinationTON: 1,
		DestinationNPI: 1,
	}
}

// IsExpired checks if the message has expired based on its validity period
func (m *Message) IsExpired() bool {
	return time.Since(m.CreatedAt) > m.ValidityPeriod
}

// CanRetry determines if the message can be retried based on configuration
func (m *Message) CanRetry(maxRetries int) bool {
	return m.RetryCount < maxRetries && !m.IsExpired()
}

// UpdateStatus updates the message status and related timestamps
func (m *Message) UpdateStatus(status MessageStatus) {
	m.Status = status
	m.UpdatedAt = time.Now()

	switch status {
	case StatusSent:
		now := time.Now()
		m.SentAt = &now
	case StatusDelivered:
		now := time.Now()
		m.DeliveredAt = &now
	}
} 