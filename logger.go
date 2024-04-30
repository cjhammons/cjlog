package cjlog

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Logger struct {
	Context map[string]interface{}
}

type LogMessage struct {
	Level   string                 `json:"level"`
	Time    string                 `json:"time"`
	Context map[string]interface{} `json:"context"`
	Message string                 `json:"message"`
}

func (l *Logger) SetContext(key string, value interface{}) {
	l.Context[key] = value
}

func NewLogger() *Logger {
	return &Logger{
		Context: make(map[string]interface{}),
	}
}

func (l *Logger) Debug(msg string) {
	l.log("DEBUG", msg)
}

func (l *Logger) Info(msg string) {
	l.log("INFO", msg)
}

func (l *Logger) Warn(msg string) {
	l.log("WARN", msg)
}

func (l *Logger) Error(msg string) {
	l.log("ERROR", msg)
}

func (l *Logger) log(level string, msg string) {

	logMessage := LogMessage{
		Level:   level,
		Time:    time.Now().Format(time.RFC3339),
		Context: l.Context,
		Message: msg,
	}

	fmt.Println(logMessage.JSON())
}

func (m *LogMessage) JSON() string {
	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("error marshaling to JSON: %v", err) // Use log.Fatalf to log the error and stop the program
	}
	return string(jsonData)
}
