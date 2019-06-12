package main

import (
	"log"
	"strconv"
	"strings"
)

var (
	wanderingInfo wanderingDamage
)

type wanderingDamage struct {
	LimbLoss        wandering `json:"limb_loss"`
	WanderingDamage wandering `json:"wandering_damage"`
	RandomDamage    wandering `json:"random_damage"`
}

type wandering struct {
	Roll  wanderingRoll    `json:"roll"`
	Table []wanderingTable `json:"table"`
}

type wanderingRoll struct {
	Dice  int `json:"dice"`
	Value int `json:"value"`
}

type wanderingTable struct {
	Outcome rollOutcome   `json:"outcome,omitempty"`
	Result  string        `json:"result"`
	Roll    wanderingRoll `json:"roll,omitempty"`
	Limb    bool          `json:"limb,omitempty"`
	Wander  bool          `json:"wander,omitempty"`
	Random  bool          `json:"random,omitempty"`
	Damage  bool          `json:"damage,omitempty"`
}

func initWanderingDmg() {
	var err error

	log.Printf("loading wandering damage info")
	err = loadInfo("wandering.json", &wanderingInfo)
	if err != nil {
		log.Fatalf("there was an issue reading the wandering file\n")
	}
	log.Printf("wandering damage info loaded")
}

func rollWanderingDamage() (response string, sendToDM bool, reroll bool) {

	var outcome int
	var damage int

	rolls := roll(wanderingInfo.WanderingDamage.Roll.Dice, wanderingInfo.WanderingDamage.Roll.Value)

	log.Printf("These are the rolls '%d'", rolls)

	outcome = total(rolls)

	log.Printf("This is the outcome '%d'", outcome)

	// this should never happen. If it does let me know...
	if outcome == 0 {
		log.Printf("If you ever log this line please open a github issue...")
		return
	}

	for _, value := range wanderingInfo.WanderingDamage.Table {
		if value.Outcome.Exact == outcome || between(value.Outcome.Range.Min, value.Outcome.Range.Max, outcome) {
			response = value.Result
			if value.Limb {
				log.Printf("rolling for limb loss")
				response = response + rollWanderingDMGLimbLoss()
			} else if value.Random {
				log.Printf("rolling on the random damage table")
				rollResponse, reroll := rollWangeringDmgRandom()
				response = response + rollResponse
				if reroll {
					log.Printf("rerolling for more random damage")
					rollResponse, reroll = rollWangeringDmgRandom()
					response = response + " " + rollResponse
				}
			} else if value.Damage {
				log.Printf("rolling for damage")
				rolls = roll(value.Roll.Dice, value.Roll.Value)
				damage = total(rolls)
				response = strings.Replace(value.Result, "&damage&", strconv.Itoa(damage), -1)
			} else if value.Wander {
				log.Printf("rolling on the wandering table again")
				reroll = value.Wander
			}
		}
	}

	// log.Printf("response: %s", response)

	return
}

func rollWanderingDMGLimbLoss() (result string) {
	log.Printf("rolling for limb loss")
	var outcome int

	rolls := roll(wanderingInfo.LimbLoss.Roll.Dice, wanderingInfo.LimbLoss.Roll.Value)

	outcome = total(rolls)

	for _, value := range wanderingInfo.LimbLoss.Table {
		if value.Outcome.Exact == 0 {
		} else if value.Outcome.Exact == outcome || between(value.Outcome.Range.Min, value.Outcome.Range.Max, outcome) {
			result = value.Result
		}
	}

	return
}

func rollWangeringDmgRandom() (result string, reroll bool) {
	log.Printf("rolling for random damages")
	var outcome int

	rolls := roll(wanderingInfo.RandomDamage.Roll.Dice, wanderingInfo.RandomDamage.Roll.Value)

	outcome = total(rolls)

	for _, value := range wanderingInfo.RandomDamage.Table {
		if value.Outcome.Exact == outcome || between(value.Outcome.Range.Min, value.Outcome.Range.Max, outcome) {
			result = value.Result
			reroll = value.Random
		}
	}

	log.Printf("result %s", result)
	log.Printf("reroll %t", reroll)

	return
}
