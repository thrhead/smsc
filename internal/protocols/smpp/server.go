package smpp

import (
	"fmt"
	"net"
	"sync"

	"github.com/sirupsen/logrus"
	"smsc/internal/config"
)

type Server struct {
	cfg    config.SMPPConfig
	log    *logrus.Logger
	ln     net.Listener
	tlsLn  net.Listener
	mu     sync.Mutex
	active bool
}

func New(cfg config.SMPPConfig, log *logrus.Logger) *Server {
	return &Server{
		cfg:    cfg,
		log:    log,
		active: false,
	}
}

func (s *Server) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.active {
		return fmt.Errorf("SMPP server is already running")
	}

	// Start non-TLS listener
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to start SMPP listener: %w", err)
	}
	s.ln = ln

	// Start TLS listener if configured
	if s.cfg.TLSPort > 0 {
		tlsAddr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.TLSPort)
		tlsLn, err := net.Listen("tcp", tlsAddr)
		if err != nil {
			s.ln.Close()
			return fmt.Errorf("failed to start SMPP TLS listener: %w", err)
		}
		s.tlsLn = tlsLn
	}

	s.active = true
	s.log.Infof("SMPP server started on %s (TLS: %s)", addr, fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.TLSPort))

	// Start accepting connections
	go s.acceptConnections()

	return nil
}

func (s *Server) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.active {
		return nil
	}

	if err := s.ln.Close(); err != nil {
		return fmt.Errorf("failed to close SMPP listener: %w", err)
	}

	if s.tlsLn != nil {
		if err := s.tlsLn.Close(); err != nil {
			return fmt.Errorf("failed to close SMPP TLS listener: %w", err)
		}
	}

	s.active = false
	s.log.Info("SMPP server stopped")
	return nil
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			if !s.active {
				return
			}
			s.log.Errorf("Failed to accept SMPP connection: %v", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// TODO: Implement SMPP protocol handling
	// 1. Read PDU
	// 2. Parse PDU
	// 3. Handle different PDU types (bind, submit_sm, etc.)
	// 4. Send response PDUs
	// 5. Implement session management
	// 6. Handle authentication
	// 7. Implement rate limiting
	// 8. Handle message routing
} 