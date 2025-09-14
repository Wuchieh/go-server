package server

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wuchieh/wtype"
)

type Config struct {
	Addr           string        `mapstructure:"addr"`
	ReadTimeout    time.Duration `mapstructure:"read_timeout"`
	WriteTimeout   time.Duration `mapstructure:"write_timeout"`
	IdleTimeout    time.Duration `mapstructure:"idle_timeout"`
	MaxHeaderBytes int           `mapstructure:"max_header_bytes"`
}

type Server struct {
	*gin.Engine

	server  *http.Server
	mu      sync.RWMutex
	running bool
}

// Run 啟動伺服器
func (s *Server) Run(ctx context.Context, addr string) error {
	return s.RunWithConfig(ctx, &Config{Addr: addr})
}

// RunWithConfig 啟動伺服器
func (s *Server) RunWithConfig(ctx context.Context, config *Config) error {
	if config == nil {
		config = &Config{}
	}

	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return ErrServerRunning
	}
	s.running = true
	s.mu.Unlock()

	// 使用 fallback 設定預設值
	s.server = &http.Server{
		Addr:           wtype.Fallback(config.Addr, ":8080"),
		Handler:        s,
		ReadTimeout:    wtype.Fallback(config.ReadTimeout, 30*time.Second),
		WriteTimeout:   wtype.Fallback(config.WriteTimeout, 30*time.Second),
		IdleTimeout:    wtype.Fallback(config.IdleTimeout, 60*time.Second),
		MaxHeaderBytes: wtype.Fallback(config.MaxHeaderBytes, 1<<20), // 1MB
	}

	return s.serve(ctx, func() error {
		return s.server.ListenAndServe()
	})
}

// RunTLS 啟動伺服器 & 配置 TLS 證書
func (s *Server) RunTLS(ctx context.Context, addr, certFile, keyFile string) error {
	return s.RunTLSWithConfig(ctx, &Config{Addr: addr}, certFile, keyFile)
}

// RunTLSWithConfig 啟動伺服器 & 配置 TLS 證書
func (s *Server) RunTLSWithConfig(ctx context.Context, config *Config, certFile, keyFile string) error {
	if config == nil {
		config = &Config{}
	}

	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return ErrServerRunning
	}
	s.running = true
	s.mu.Unlock()

	s.server = &http.Server{
		Addr:           wtype.Fallback(config.Addr, ":443"),
		Handler:        s,
		ReadTimeout:    wtype.Fallback(config.ReadTimeout, 30*time.Second),
		WriteTimeout:   wtype.Fallback(config.WriteTimeout, 30*time.Second),
		IdleTimeout:    wtype.Fallback(config.IdleTimeout, 60*time.Second),
		MaxHeaderBytes: wtype.Fallback(config.MaxHeaderBytes, 1<<20),
	}

	return s.serve(ctx, func() error {
		return s.server.ListenAndServeTLS(certFile, keyFile)
	})
}

// 通用的服務啟動邏輯
func (s *Server) serve(ctx context.Context, startFn func() error) error {
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	go func() {
		if err := startFn(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			cancel(err)
		} else {
			cancel(http.ErrServerClosed)
		}
	}()

	<-ctx.Done()

	s.mu.Lock()
	s.running = false
	s.mu.Unlock()

	if err := context.Cause(ctx); err != nil {
		if errors.Is(err, http.ErrServerClosed) || errors.Is(err, context.Canceled) {
			return nil
		}
		return err
	}
	return nil
}

// Stop 關閉伺服器
func (s *Server) Stop(ctx context.Context) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.running || s.server == nil {
		return nil
	}
	return s.server.Shutdown(ctx)
}

// IsRunning 狀態查詢方法
func (s *Server) IsRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.running
}

// Addr 取得伺服器 addr
func (s *Server) Addr() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.server != nil {
		return s.server.Addr
	}
	return ""
}

// Shutdown 優雅關機的輔助方法
func (s *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.Stop(ctx)
}

// New 創建伺服器
func New() *Server {
	s := &Server{}
	s.Engine = gin.New()
	return s
}
