package powerlog

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func Test_Init_ShouldSetLogFlags(t *testing.T) {
	log.SetFlags(0) // Reset flags to default
	Init()
	if log.Flags() != (log.LstdFlags | log.Lshortfile) {
		t.Errorf("Init() log flags = %v, want %v", log.Flags(), log.LstdFlags|log.Lshortfile)
	}
}

func Test_Init_ShouldInitializePowerLogger(t *testing.T) {
	logger := newPowerLogger()
	if logger == nil {
		t.Errorf("newPowerLogger() = %v, want non-nil", logger)
	}
}

func Test_NewPowerLog_ShouldReturnNonNil(t *testing.T) {
	if got := newPowerLogger(); got == nil {
		t.Errorf("newPowerLogger() = %v, want non-nil", got)
	}
}

func Test_NewPowerLog_ShouldInitializeLoggers(t *testing.T) {
	got := newPowerLogger()
	if got.debug == nil || got.info == nil || got.warn == nil || got.error == nil {
		t.Errorf("newPowerLogger() loggers not initialized properly")
	}
}

func Test_NewPowerLog_ShouldSetCorrectLoggerPrefixes(t *testing.T) {
	got := newPowerLogger()
	if got.debug.Prefix() != "DEBUG: " {
		t.Errorf("newPowerLogger() debug logger prefix = %v, want %v", got.debug.Prefix(), "DEBUG: ")
	}
	if got.info.Prefix() != "INFO: " {
		t.Errorf("newPowerLogger() info logger prefix = %v, want %v", got.info.Prefix(), "INFO: ")
	}
	if got.warn.Prefix() != "WARN: " {
		t.Errorf("newPowerLogger() warn logger prefix = %v, want %v", got.warn.Prefix(), "WARN: ")
	}
	if got.error.Prefix() != "ERROR: " {
		t.Errorf("newPowerLogger() error logger prefix = %v, want %v", got.error.Prefix(), "ERROR: ")
	}
}

func Test_powerLog_Debug(t *testing.T) {
	type fields struct {
		debug *log.Logger
		info  *log.Logger
		warn  *log.Logger
		error *log.Logger
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Should be Success Debug Log", fields{debug: log.New(os.Stdout, "DEBUG: ", log.LstdFlags)}, args{v: []any{"test"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &powerLogger{
				debug: tt.fields.debug,
				info:  tt.fields.info,
				warn:  tt.fields.warn,
				error: tt.fields.error,
			}
			l.Debug(tt.args.v...)
		})
	}
}

func Test_powerLog_Error(t *testing.T) {
	type fields struct {
		debug *log.Logger
		info  *log.Logger
		warn  *log.Logger
		error *log.Logger
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &powerLogger{
				debug: tt.fields.debug,
				info:  tt.fields.info,
				warn:  tt.fields.warn,
				error: tt.fields.error,
			}
			l.Error(tt.args.v...)
		})
	}
}

func Test_powerLog_Info(t *testing.T) {
	type fields struct {
		debug *log.Logger
		info  *log.Logger
		warn  *log.Logger
		error *log.Logger
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &powerLogger{
				debug: tt.fields.debug,
				info:  tt.fields.info,
				warn:  tt.fields.warn,
				error: tt.fields.error,
			}
			l.Info(tt.args.v...)
		})
	}
}

func Test_powerLog_Panic(t *testing.T) {
	type fields struct {
		debug *log.Logger
		info  *log.Logger
		warn  *log.Logger
		error *log.Logger
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &powerLogger{
				debug: tt.fields.debug,
				info:  tt.fields.info,
				warn:  tt.fields.warn,
				error: tt.fields.error,
			}
			l.Panic(tt.args.v...)
		})
	}
}

func Test_powerLog_Warn(t *testing.T) {
	type fields struct {
		debug *log.Logger
		info  *log.Logger
		warn  *log.Logger
		error *log.Logger
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &powerLogger{
				debug: tt.fields.debug,
				info:  tt.fields.info,
				warn:  tt.fields.warn,
				error: tt.fields.error,
			}
			l.Warn(tt.args.v...)
		})
	}
}

func TestSetLoggerFile(t *testing.T) {
	type args struct {
		logName string
		logDir  string
	}

	var logName = "logName"
	logFileName := fmt.Sprintf("%s_%s.log", time.Now().Local().Format("2006-01-02"), logName)
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		wantErrorMessage string
	}{
		{"Should be Success", args{logName, "/tmp"}, false, ""},
		{"Should be Invalid Path", args{logName, "/invalid/tmp"}, true, fmt.Sprintf("error creating log file: open %s: no such file or directory", "/invalid/tmp/"+logFileName)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetLoggerFile(tt.args.logName, tt.args.logDir); err != nil && tt.wantErrorMessage != err.Error() {
				t.Errorf("SetLoggerFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
