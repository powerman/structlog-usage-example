package main

import (
	"flag"

	"github.com/powerman/structlog"
	"github.com/powerman/structlog-usage-example/pkg"
)

// Log field names.
const (
	LogHost   = "host"
	LogPort   = "port"
	LogRemote = "remote" // aligned IPv4:Port "   192.168.0.42:1234 "
	LogFunc   = "func"   // RPC method name, REST resource path
	LogUser   = "userID"
)

func main() {
	logLevel := flag.String("log.level", "debug", "log `level` (debug|info|warn|err)")
	flag.Parse()

	// Wrong log.level is not fatal, it will be reported and set to "debug".
	structlog.DefaultLogger.SetLogLevel(structlog.ParseLevel(*logLevel))

	// - Переменная log вместо logger.
	// - SetOutput не нужен.
	// - Не надо в каждой функции вызывать structlog.New().
	var log = structlog.New()

	// log.Info("started", "version", "0.1.0")

	structlog.DefaultLogger.
		AppendPrefixKeys(
			LogRemote,
			LogFunc,
		).
		SetSuffixKeys(
			LogUser,
			structlog.KeyStack,
		).
		SetKeysFormat(map[string]string{
			structlog.KeyUnit: " %6[2]s:", // set to max KeyUnit/package length
			LogHost:           " %[2]s",
			LogPort:           ":%[2]v",
			LogRemote:         " %-21[2]s",
			LogFunc:           " %[2]s:",
			LogUser:           " %[2]v",
			"version":         " %s %v",
			"err":             " %s: %v",
			"json":            " %s=%#q",
			"ptr":             " %[2]p", // for debugging references
		})

	// - Info и key/value вместо Printf.
	log.Info("started", "version", "0.1.0")

	log.SetDefaultKeyvals(
		structlog.KeyUnit, "main",
	)

	log.Info("started", "version", "0.1.0")

	request(log)

	err := pkg.Something(42)
	if err != nil {
		log.Warn("temporary error", "err", err)
	}
	err = pkg.Something2(42)
	if err != nil {
		log.Warn("temporary error", "err", err)
	}

	log.SetLogFormat(structlog.JSON)
	log.Info("started", "version", "0.1.0")
}

// Log is a synonym for convenience.
type Log = *structlog.Logger

func request(log Log) {
	log = log.New(LogRemote, "127.0.0.1:1234")
	auth(log)

	log.SetDefaultKeyvals(LogRemote, "123.123.123.123:12345")
	handleOtherthing(log)
}

func auth(log Log) {
	log = log.New(LogUser, "someuser")
	handleSomething(log)
}

func handleSomething(log Log) {
	log = log.New(LogFunc, "/some/thing")
	log.Info("incoming request")
}

func handleOtherthing(log Log) {
	log = log.New(LogFunc, "/other/thing")
	log.Info("incoming request", "someparam", "value")
}
