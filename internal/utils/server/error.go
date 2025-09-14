package server

import "errors"

var (
	ErrServerRunning = errors.New("server is already running")
)
