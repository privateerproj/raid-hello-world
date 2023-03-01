package pack

import (
	"log"

	"github.com/hashicorp/go-hclog"
	"github.com/privateerproj/privateer-sdk/logging"
	"github.com/privateerproj/privateer-sdk/utils"
	"github.com/spf13/viper"
)

var logger hclog.Logger // enables formatted logging (log.Printf("[TRACE] , etc)

// KnockKnock is a demo test for dev purposes
func KnockKnock() error {
	name, err := getJokeName()
	if err != nil {
		return err
	}
	log.Printf("[INFO] Me: Knock Knock")
	log.Printf("[INFO] %s: Who's There?", name)
	// Demo the log timestamp
	for i := 1; i < 5000000; i++ {
		if i%500000 == 0 {
			log.Printf("[TRACE] Me: (stares at %s)", name)
		}
	}
	log.Printf("%s: (lost interest and left)", name)
	return nil
}

// ChickenCrossedRoad is a demo test for dev purposes
func ChickenCrossedRoad() error {
	name, err := getJokeName()
	if err != nil {
		return err
	}
	logger.Warn("[Warn!] Me: This joke may offend someone.")
	logger.Info("[Info!] Me: Why did the chicken cross the road?")
	log.Print("[TRACE] Me: Although, wait, there is something else you should know")
	log.Printf("[TRACE] Me: (looks to see what %s's expression is)", name)
	log.Printf("%s: I'm busy, leave me alone.", name)
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
