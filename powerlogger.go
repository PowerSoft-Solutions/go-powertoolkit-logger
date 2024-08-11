package powerlogger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type PowerLogger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

var withFile bool

func Init() *PowerLogger {
	return &PowerLogger{
		debug: log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
		info:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warn:  log.New(os.Stdout, "WARN: ", log.LstdFlags),
		error: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
	}
}

func (l *PowerLogger) Debug(v ...any) {
	if withFile {
		l.debug.Println(v...)
	}
	log.Printf("DEBUG: %s", v)
}

func (l *PowerLogger) Info(v ...any) {
	if withFile {
		l.info.Println(v...)
	}
	log.Printf("INFO: %s", v)
}

func (l *PowerLogger) Warn(v ...any) {
	if withFile {
		l.warn.Println(v...)
	}
	log.Printf("WARN: %s", v)
}

func (l *PowerLogger) Error(v ...any) {
	if withFile {
		l.error.Println(v...)
	}
	log.Printf("ERROR: %s", v)
}

func (l *PowerLogger) Panic(v ...any) {
	if withFile {
		l.error.Println(v...)
	}
	log.Printf("PANIC: %s", v)
	panic(fmt.Sprint(v...))
}

func (l *PowerLogger) SetLoggerFile(logName, logDir string) error {
	if logName == "" {
		return fmt.Errorf("log name is required")
	}

	logFileName := fmt.Sprintf("%s_%s.log", time.Now().Local().Format("2006-01-02"), logName)
	logFile, err := os.OpenFile(filepath.Join(logDir, logFileName), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error creating log file: %s", err.Error())
	}
	log.SetOutput(logFile)
	withFile = true
	return nil
}
