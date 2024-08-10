package powerlog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type powerLogger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

var withFile bool

func Init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	newPowerLogger()
}

func newPowerLogger() *powerLogger {
	return &powerLogger{
		debug: log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
		info:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warn:  log.New(os.Stdout, "WARN: ", log.LstdFlags),
		error: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
	}
}

func (l *powerLogger) Debug(v ...any) {
	if withFile {
		l.debug.Println(v...)
	}
	log.Printf("DEBUG: %s", v)
}

func (l *powerLogger) Info(v ...any) {
	if withFile {
		l.info.Println(v...)
	}
	log.Printf("INFO: %s", v)
}

func (l *powerLogger) Warn(v ...any) {
	if withFile {
		l.warn.Println(v...)
	}
	log.Printf("WARN: %s", v)
}

func (l *powerLogger) Error(v ...any) {
	if withFile {
		l.error.Println(v...)
	}
	log.Printf("ERROR: %s", v)
}

func (l *powerLogger) Panic(v ...any) {
	if withFile {
		l.error.Println(v...)
	}
	log.Printf("PANIC: %s", v)
	panic(fmt.Sprint(v...))
}

func SetLoggerFile(logName, logDir string) error {
	logFileName := fmt.Sprintf("%s_%s.log", time.Now().Local().Format("2006-01-02"), logName)
	logFile, err := os.OpenFile(filepath.Join(logDir, logFileName), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error creating log file: %s", err.Error())
	}
	log.SetOutput(logFile)
	withFile = true
	return nil
}
