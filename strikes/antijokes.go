package strikes

import (
	"errors"
	"fmt"
	"log"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"

	"github.com/privateerproj/privateer-sdk/raidengine"
)

type Antijokes struct {
	Log hclog.Logger
}

// KnockKnock is a demo test for dev purposes
func (a *Antijokes) KnockKnock() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "Knock Knock"
	result = raidengine.StrikeResult{
		Passed:  false,
		Message: "",
		DocsURL: "",
	}
	log.Printf("Knock Knock")
	name, err := getJokeName()
	if err != nil {
		return
	}
	log.Printf("Me: Knock Knock")
	log.Printf(fmt.Sprintf("%s: Who's There?", name))
	// Demo the log timestamp
	for i := 1; i < 5000000; i++ {
		if i%500000 == 0 {
			log.Printf("Me: (stares at %s)", name)
		}
	}
	log.Printf(fmt.Sprintf("%s: (lost interest and left)", name))
	return
}

// ChickenCrossedRoad is a demo test for dev purposes
func (a *Antijokes) ChickenCrossedRoad() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "Chicken Crossed Road"
	result = raidengine.StrikeResult{
		Passed:  true,
		Message: "",
		DocsURL: "",
	}

	name, err := getJokeName()
	if err != nil {
		return
	}
	a.Log.Warn("Me: This joke may offend someone.")
	a.Log.Info("Me: Why did the chicken cross the road?")
	a.Log.Trace(fmt.Sprintf("Me: (looks to see what %s's expression is)", name))
	a.Log.Info(fmt.Sprintf("%s: I'm busy, leave me alone.", name))
	return
}

func getJokeName() (string, error) {
	if viper.IsSet("raids.wireframe.jokename") {
		return viper.GetString("raids.wireframe.jokename"), nil
	}
	return "", errors.New("JokeName must be set in the config file or env vars (PVTR_WIREFRAME_JOKE_NAME)")
}
