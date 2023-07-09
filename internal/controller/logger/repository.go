package logger

import (
	"log"
	"os"
)

type LoggerRepositoryImpl struct {
	lvl string
}

func NewLogger(lvl string) *LoggerRepositoryImpl {
	file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(file)
	return &LoggerRepositoryImpl{lvl: lvl}
}

func (l *LoggerRepositoryImpl) Debug(msg string) {
	if l.lvl == "debug" {
		log.Println("DEBUG | " + msg)
	}
}

func (l *LoggerRepositoryImpl) Info(msg string) {
	if l.lvl == "debug" || l.lvl == "info" {
		log.Println("Info | " + msg)
	}
}

func (l *LoggerRepositoryImpl) Warn(msg string) {
	log.Println("WARN | " + msg)
}

func (l *LoggerRepositoryImpl) Error(msg string) {
	log.Println("ERROR | " + msg)
	os.Exit(1)
}
