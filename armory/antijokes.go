package armory

import (
	"fmt"
	"log"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

// Conforms to the Armory interface type
type Antijokes struct {
	Tactics map[string][]raidengine.Strike     // Required, allows you to sort which strikes are run for each control
	Log     hclog.Logger                       // Recommended, allows you to set the log level for each log message
	Results map[string]raidengine.StrikeResult // Optional, allows cross referencing between strikes
}

func (a *Antijokes) SetLogger(loggerName string) hclog.Logger {
	a.Log = raidengine.GetLogger(loggerName, false)
	return a.Log
}

func (a *Antijokes) GetTactics() map[string][]raidengine.Strike {
	return a.Tactics
}

// KnockKnock conforms to the Strike function type
func (a *Antijokes) KnockKnock() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "Knock Knock"
	log.Print(strikeName) // Default logs will be set as INFO

	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "This is a failure of a joke for dev purposes",
		Message:     "Strike has not yet started.",
		DocsURL:     "https://maintainer.com/docs/raids/wireframe",
		ControlID:   "CCC-Taxonomy-1",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	// run getJokerName() as a movement
	jokerNameMovement := getJokerName()
	result.Movements["JokerName"] = jokerNameMovement
	if !jokerNameMovement.Passed {
		result.Message = jokerNameMovement.Message
		return
	}

	// run getJokeeName() as a movement
	jokeeNameMovement := getJokeeName()
	result.Movements["JokeeName"] = jokeeNameMovement
	if !jokeeNameMovement.Passed {
		result.Message = jokeeNameMovement.Message
		return
	}

	// Run multiple movements at once by passing the result object as a pointer
	// Previous movement values must be cast to their type from interface{} before being used again
	RunKnockKnock(jokerNameMovement.Value.(string), jokeeNameMovement.Value.(string), &result)
	return
}

// ChickenCrossedRoad conforms to the Strike function type
func (a *Antijokes) ChickenCrossedRoad() (strikeName string, result raidengine.StrikeResult) {
	// If a strike is part of multiple tactics, you can use a map to reference the control ID to the selected tactic
	controlIDs := map[string]string{
		"CCC-Taxonomy":  "CCC-Taxonomy-2",
		"CCC-Hardening": "CCC-Hardening-1",
		"CIS":           "CIS-1",
	}
	strikeName = "Chicken Crossed Road"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "This is a demo strike for dev purposes",
		Message:     "Strike has not yet started.",
		DocsURL:     "https://maintainer.com/docs/raids/wireframe",
		ControlID:   controlIDs[viper.GetString("raids.wireframe.tactic")],
		Movements:   make(map[string]raidengine.MovementResult),
	}

	jokerNameMovement := getJokerName()
	result.Movements["JokerName"] = jokerNameMovement
	if !jokerNameMovement.Passed {
		result.Message = jokerNameMovement.Message
		return
	}

	// Using an hclog logger will allow you to set the log level for each message
	a.Log.Warn(fmt.Sprintf("%s: This joke may offend someone.", jokerNameMovement.Value))
	a.Log.Info(fmt.Sprintf("%s: Why did the chicken cross the road?", jokerNameMovement.Value))
	a.Log.Trace(fmt.Sprintf("%s: (looks to see what the stranger's expression is)", jokerNameMovement.Value))
	a.Log.Info("Stranger: I'm busy, leave me alone.")

	result.Passed = true
	result.Message = "We tried, and that's all we really came here for."
	return
}

// getJokerName is a common movement for the strikes in this raid
func getJokerName() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokerName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	if viper.IsSet("raids.wireframe.JokerName") {
		result.Passed = true
		result.Message = "JokerName is set"
		result.Value = viper.GetString("raids.wireframe.JokerName")
	} else {
		result.Passed = false
		result.Message = "JokerName must be set in the configuration."
	}
	return
}

// getJokerName is a single movement for common use by the strikes in this raid
func getJokeeName() (result raidengine.MovementResult) {
	result = raidengine.MovementResult{
		Description: "JokeeName must be found in the runtime configuration.",
		Function:    utils.CallerPath(0),
	}
	if viper.IsSet("raids.wireframe.JokeeName") {
		result.Passed = true
		result.Message = "JokeeName is set"
		result.Value = viper.GetString("raids.wireframe.JokeeName")
	} else {
		result.Passed = false
		result.Message = "JokeeName must be set in the configuration."
	}
	return
}

// RunKnockKnock is a set of movements for the Knock Knock strike, with each movement being added to the provided result
func RunKnockKnock(jokerName, jokeeName string, result *raidengine.StrikeResult) {
	// say knock knock
	result.Movements["knock knock"] = raidengine.MovementResult{
		Passed:      true,
		Description: "Joke must be started by the joker.",
		Message:     fmt.Sprintf("%s: Knock knock.", jokerName),
		Function:    utils.CallerPath(0), // 0 logs the current function name
	}
	// say who's there
	result.Movements["who's there"] = raidengine.MovementResult{
		Passed:      true,
		Description: "Jokee must respond with 'who's there?'",
		Message:     fmt.Sprintf("%s: Who's there?", jokeeName),
		Function:    utils.CallerPath(0),
	}
	// say punchline
	result.Movements["punchline"] = raidengine.MovementResult{
		Passed:      true,
		Description: "Jokee must respond with the punchline.",
		Message:     fmt.Sprintf("%s: %s", jokerName, jokeeName),
		Function:    utils.CallerPath(0),
	}
}
