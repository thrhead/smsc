package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host           string
	Port           int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

type Server struct {
	cfg    Config
	log    *logrus.Logger
	router *gin.Engine
	srv    *http.Server
}

func New(cfg Config, log *logrus.Logger) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// CORS middleware ekle
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	s := &Server{
		cfg:    cfg,
		log:    log,
		router: router,
	}

	s.setupRoutes()
	return s
}

func (s *Server) Start() error {
	// Set up routes
	s.setupRoutes()

	// Configure HTTP server
	s.srv = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		Handler:        s.router,
		ReadTimeout:    s.cfg.ReadTimeout,
		WriteTimeout:   s.cfg.WriteTimeout,
		MaxHeaderBytes: s.cfg.MaxHeaderBytes,
	}

	// Start server
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Errorf("HTTP server error: %v", err)
		}
	}()

	s.log.Infof("API server started on %s:%d", s.cfg.Host, s.cfg.Port)
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}
	return nil
}

func (s *Server) setupRoutes() {
	// Health check
	s.router.GET("/health", s.healthCheck)

	// API v1 routes
	v1 := s.router.Group("/api/v1")
	{
		// Message endpoints
		messages := v1.Group("/messages")
		{
			messages.POST("/send", s.sendMessage)
			messages.GET("/status/:id", s.getMessageStatus)
			messages.GET("/list", s.listMessages)
		}

		// Operator endpoints
		operators := v1.Group("/operators")
		{
			operators.GET("/", s.listOperators)
			operators.POST("/", s.addOperator)
			operators.PUT("/:id", s.updateOperator)
			operators.DELETE("/:id", s.deleteOperator)
		}

		// Routing endpoints
		routing := v1.Group("/routing")
		{
			routing.GET("/rules", s.listRoutingRules)
			routing.POST("/rules", s.addRoutingRule)
			routing.PUT("/rules/:id", s.updateRoutingRule)
			routing.DELETE("/rules/:id", s.deleteRoutingRule)
		}

		// System endpoints
		system := v1.Group("/system")
		{
			system.GET("/status", s.getSystemStatus)
			system.GET("/metrics", s.getMetrics)
		}
	}
}

// Handler implementations

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Server) sendMessage(c *gin.Context) {
	// TODO: Implement message sending
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) getMessageStatus(c *gin.Context) {
	// TODO: Implement message status retrieval
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) listMessages(c *gin.Context) {
	// TODO: Implement message listing
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) listOperators(c *gin.Context) {
	operators := []map[string]interface{}{
		{
			"id":       1,
			"name":     "Operator 1",
			"priority": 1,
			"weight":   100,
			"maxTps":   1000,
			"status":   "active",
		},
		{
			"id":       2,
			"name":     "Operator 2",
			"priority": 2,
			"weight":   50,
			"maxTps":   500,
			"status":   "active",
		},
	}
	c.JSON(http.StatusOK, operators)
}

func (s *Server) addOperator(c *gin.Context) {
	var operator struct {
		Name     string `json:"name" binding:"required"`
		Priority int    `json:"priority" binding:"required"`
		Weight   int    `json:"weight" binding:"required"`
		MaxTps   int    `json:"maxTps" binding:"required"`
	}

	if err := c.ShouldBindJSON(&operator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the received operator data
	s.log.WithFields(logrus.Fields{
		"name":     operator.Name,
		"priority": operator.Priority,
		"weight":   operator.Weight,
		"maxTps":   operator.MaxTps,
	}).Info("Adding new operator")

	response := map[string]interface{}{
		"id":       time.Now().UnixNano(),
		"name":     operator.Name,
		"priority": operator.Priority,
		"weight":   operator.Weight,
		"maxTps":   operator.MaxTps,
		"status":   "active",
	}

	c.JSON(http.StatusCreated, response)
}

func (s *Server) updateOperator(c *gin.Context) {
	id := c.Param("id")
	var operator struct {
		Name     string `json:"name" binding:"required"`
		Priority int    `json:"priority" binding:"required"`
		Weight   int    `json:"weight" binding:"required"`
		MaxTps   int    `json:"maxTps" binding:"required"`
	}

	if err := c.ShouldBindJSON(&operator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"id":       id,
		"name":     operator.Name,
		"priority": operator.Priority,
		"weight":   operator.Weight,
		"maxTps":   operator.MaxTps,
		"status":   "active",
	}

	c.JSON(http.StatusOK, response)
}

func (s *Server) deleteOperator(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Operator %s deleted successfully", id)})
}

func (s *Server) listRoutingRules(c *gin.Context) {
	// TODO: Implement routing rule listing
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) addRoutingRule(c *gin.Context) {
	// TODO: Implement routing rule addition
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) updateRoutingRule(c *gin.Context) {
	// TODO: Implement routing rule update
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) deleteRoutingRule(c *gin.Context) {
	// TODO: Implement routing rule deletion
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) getSystemStatus(c *gin.Context) {
	// TODO: Implement system status retrieval
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Server) getMetrics(c *gin.Context) {
	// TODO: Implement metrics retrieval
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
} 