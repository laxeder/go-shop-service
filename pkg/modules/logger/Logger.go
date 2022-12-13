package logger

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	PathName     string = "logs/pool.log"
	MaxSizeFile  int    = 512  // 512 MB
	MaxDaysFile  int    = 28   // dias que dura um arquivo
	MaxBackups   int    = 3    // máximo de backups armazenados
	CompressFile bool   = true // comprimir o arquivo
)

var LOGGER *zerolog.Logger

var StaticCaller string
var StaticError error

func New() *zerolog.Logger {
	LOGGER := LoggerProvider{}
	return LOGGER.New()
}

type LoggerProviderHook struct {
	loggerStream zerolog.Logger
	loggerEvent  *zerolog.Event
}

func (lph *LoggerProviderHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	lph.LoggerStream(level.String(), StaticError, StaticCaller, msg)
	StaticCaller = ""
	StaticError = nil
}

func (lph *LoggerProviderHook) LoggerStream(level string, errorStack error, caller string, msg string) *zerolog.Logger {

	configStreamOut := &lumberjack.Logger{
		Filename:   PathName,
		MaxSize:    MaxSizeFile,
		MaxBackups: MaxBackups,
		MaxAge:     MaxDaysFile,
		Compress:   CompressFile,
	}

	lph.loggerStream = zerolog.New(configStreamOut).With().Logger()

	lph.HandleLevel(level)
	lph.AddDateTime()
	lph.AddCaller(caller)
	lph.AddErrorStack(errorStack)
	lph.loggerEvent.Msg(msg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for {
			<-c
			configStreamOut.Rotate()
		}
	}()

	return &lph.loggerStream
}

func (lph *LoggerProviderHook) AddDateTime() *zerolog.Event {
	return lph.loggerEvent.Str("time", time.Now().Format(time.RFC3339Nano))
}

func (lph *LoggerProviderHook) AddCaller(caller string) *zerolog.Event {
	return lph.loggerEvent.Str("caller", caller)
}

func (lph *LoggerProviderHook) AddErrorStack(errorStack error) *zerolog.Event {
	if errorStack == nil {
		return lph.loggerEvent.Str("stack", "")
	}
	return lph.loggerEvent.Str("stack", fmt.Sprintf("%v", errorStack))
}

func (lph *LoggerProviderHook) HandleLevel(level string) *zerolog.Event {

	switch strings.ToLower(level) {
	case "trace":
		lph.loggerEvent = lph.loggerStream.Panic()
	case "debug":
		lph.loggerEvent = lph.loggerStream.Debug()
	case "info":
		lph.loggerEvent = lph.loggerStream.Info()
	case "warn":
		lph.loggerEvent = lph.loggerStream.Warn()
	case "error":
		lph.loggerEvent = lph.loggerStream.Error()
	case "fatal":
		lph.loggerEvent = lph.loggerStream.Fatal()
	case "panic ":
		lph.loggerEvent = lph.loggerStream.Panic()
	default:
		lph.loggerEvent = lph.loggerStream.Log().Str("level", "log")
	}
	return lph.loggerEvent
}

type LoggerProvider struct {
	instance  zerolog.Logger
	stdOut    zerolog.ConsoleWriter
	streamOut zerolog.Logger
}

func (l *LoggerProvider) New() *zerolog.Logger {

	var loggerProviderHook LoggerProviderHook

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackFieldName = "stack"

	stdOut := l.StdOut()

	l.instance = zerolog.New(stdOut).With().Timestamp().Caller().Stack().Logger().Hook(&loggerProviderHook)

	return &l.instance
}

func (l *LoggerProvider) StdOut() *zerolog.ConsoleWriter {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackFieldName = "stack"

	zerolog.CallerMarshalFunc = func(file string, line int) string {
		caller := file + ":" + strconv.Itoa(line)
		StaticCaller = caller
		return caller
	}

	zerolog.ErrorMarshalFunc = func(err error) interface{} {
		StaticError = err
		return err
	}

	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		return pkgerrors.MarshalStack(err)
	}

	l.stdOut = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	l.stdOut.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("[%v]", i))
	}

	l.stdOut.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf(" %s ", i)
	}

	l.stdOut.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}

	return &l.stdOut
}
