package main

import (
	"github.com/codegoalie/golibnotify"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sevlyar/go-daemon"
	"os"
	"time"
)

func main() {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stdout, FormatTimestamp: func(i interface{}) string { return "" }})
	if len(os.Args) < 2 {
		logger.Fatal().Msg("Please provide duration")
	}
	d, err := time.ParseDuration(os.Args[1])
	if err != nil {
		logger.Fatal().Err(err).Msg("Couldn't parse the duration")
	}
	summary := "Time's up"
	text := ""
	if len(os.Args) >= 3 {
		summary = os.Args[2]
	}
	if len(os.Args) >= 4 {
		text = os.Args[3]
	}
	cntxt := &daemon.Context{
		Umask: 027,
		Args:  os.Args,
	}

	child, err := cntxt.Reborn()
	if err != nil {
		logger.Fatal().Err(err).Msg("Couldn't start daemon")
	}
	if child != nil {
		logger.Info().Int("pid", child.Pid).Msg("Success. Notification will be sent.")
		os.Exit(0)
	}
	defer cntxt.Release()
	logger.Info().Msg("Child")
	n := golibnotify.NewSimpleNotifier("gonotify")
	select {
	case <-time.NewTimer(d).C:
	}
	err = n.Update(summary, text, "")
	if err != nil {
		logger.Panic().Msg("Couldn't send notification")
	}
}
