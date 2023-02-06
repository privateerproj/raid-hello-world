package pack

import (
	"log"

	"github.com/privateerproj/privateer-pack-wireframe/internal/config"
	"github.com/privateerproj/privateer-sdk/utils"
)

// KnockKnock is a demo test for dev purposes
func KnockKnock() error {
	name, err := getJokeName()
	if err != nil {
		return err
	}
	log.Printf("Knock Knock")
	log.Printf("Who's There?")
	log.Printf(name)
	log.Printf("::The listener has lost interest and left::")
	return nil
}

// ChickenCrossedRoad is a demo test for dev purposes
func ChickenCrossedRoad() error {
	name, err := getJokeName()
	if err != nil {
		return err
	}
	log.Printf("Me: Why did the chicken cross the road?")
	log.Printf("%s: I'm busy, leave me alone.", name)
	return nil
}

func getJokeName() (string, error) {
	if config.Vars.Raids.Wireframe.JokeName != "" {
		return config.Vars.Raids.Wireframe.JokeName, nil
	}
	return "", utils.ReformatError("JokeName must be set in the config file or env vars (PVTR_WIREFRAME_JOKE_NAME)")
}
