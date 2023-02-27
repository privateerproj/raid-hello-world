package pack

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/privateerproj/privateer-sdk/logging"
	"github.com/privateerproj/privateer-sdk/utils"
	"github.com/spf13/viper"
)

var logger hclog.Logger // enables formatted logging (logger.Trace, etc)

// KnockKnock is a demo test for dev purposes
func KnockKnock() error {
	name, err := getJokeName()
	if err != nil {
		return err
	}
	logger.Info("Me: Knock Knock")
	logger.Info(fmt.Sprintf("%s: Who's There?", name))
	// Demo the log timestamp
	for i := 1; i < 5000000; i++ {
		if i%500000 == 0 {
			logger.Trace(fmt.Sprintf("Me: (stares at %s)", name))
		}
	}
	logger.Error(fmt.Sprintf("%s: (lost interest and left)", name))
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
	logger.Debug("Me: Although, wait, there is something else you should know")
	logger.Trace(fmt.Sprintf(
		"Me: (looks to see what %s's expression is)", name))
	logger.Error(fmt.Sprintf(
		"%s: I'm busy, leave me alone.", name))
	return nil
}

func getJokeName() (string, error) {
	if viper.GetString("PVTR_WIREFRAME_JOKE_NAME") != "" {
		return viper.GetString("PVTR_WIREFRAME_JOKE_NAME"), nil
	}
	return "", utils.ReformatError("JokeName must be set in the config file or env vars (PVTR_WIREFRAME_JOKE_NAME)")
}

func init() {
	logger = logging.GetLogger("this", viper.GetString("loglevel"), false) // loglevel not yet set
}
