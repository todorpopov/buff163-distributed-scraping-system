package logger

import "time"

type LogMessage struct {
	Service  string
	Datetime time.Time
	Type     string
	Message  string
}

type LoggerConfig struct {
	Url     string
	Port    string
	Service string
}

type LoggerClient struct {
	Config LoggerConfig
}
