package main

import (
	"github.com/bitrise-steplib/bitrise-step-restore-carthage-cache/step"
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/env"
	. "github.com/bitrise-io/go-utils/v2/exitcode"
	"github.com/bitrise-io/go-utils/v2/log"
)

func main() {
	exitCode := run()
	os.Exit(int(exitCode))
}

func run() ExitCode {
	logger := log.NewLogger()
	envRepo := env.NewRepository()
	inputParser := stepconf.NewInputParser(envRepo)
	cmdFactory := command.NewFactory(envRepo)
	cacheStep := step.New(logger, inputParser, envRepo, cmdFactory)

	if err := cacheStep.Run(); err != nil {
		logger.Errorf(err.Error())
		return Failure
	}

	return Success
}
