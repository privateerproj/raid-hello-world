package pack

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/privateerproj/privateer-sdk/logging"
	"github.com/privateerproj/privateer-sdk/utils"
	"github.com/spf13/viper"
)

var logger hclog.Logger
var logger2 hclog.Logger // consider how we should do intercept logs

// KnockKnock is a demo test for dev purposes
func KnockKnock() error {
	name, err := getJokeName()
	if err != nil {
		return err
	}
	logger2.Info("Me: Knock Knock")
	logger2.Info(fmt.Sprintf("%s: Who's There?", name))
	// Demo the log timestamp
	for i := 1; i < 5000000; i++ {
		if i%500000 == 0 {
			logger.Trace("Me: (stares at %s)", name)
		}
	}
	logger2.Info(fmt.Sprintf("%s: (lost interest and left)", name))
	return nil
}

// ChickenCrossedRoad is a demo test for dev purposes
func ChickenCrossedRoad() error {
	name, err := getJokeName()
	if err != nil {
		return err
	}
	logger.Warn("Me: This joke may offend someone.")
	logger.Info("Me: Why did the chicken cross the road?")
	logger.Trace(fmt.Sprintf("Me: (looks to see what %s's expression is)", name))
	logger.Info(fmt.Sprintf("%s: I'm busy, leave me alone.", name))
	return nil
}

func getJokeName() (string, error) {
	if viper.IsSet("raids.wireframe.jokename") {
		return viper.GetString("raids.wireframe.jokename"), nil
	}
	return "", utils.ReformatError("JokeName must be set in the config file or env vars (PVTR_WIREFRAME_JOKE_NAME)")
}

func init() {
	logger = logging.GetLogger("cli", viper.GetString("loglevel"), true)
	logger2 = logging.GetLogger("output", viper.GetString("loglevel"), true)
}
