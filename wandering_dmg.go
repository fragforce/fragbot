package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	wanderingInfo wanderingDamage
)

type wanderingDamage struct {
	LimbLoss        wandering `json:"limb_loss"`
	WanderingDamage wandering `json:"wandering_damage"`
	RandomDamage    wandering `json:"random_damage"`
}

func initWanderingDmg() {
	var err error

	log.Printf("loading wandering damage info")
	err = loadInfo("wandering/wandering_dmg.json", &wanderingInfo)
	if err != nil {
		log.Fatalf("there was an issue reading the wandering file\n")
	}
	log.Printf("wandering damage info loaded")
}

func rollWanderingDamage() (response string, discordEmbed discordgo.MessageEmbed, sendToDM bool) {

	var outcome int
	var damage int
	var wander = true

	for wander {
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
				response = response + "\n" + value.Result
				if value.Limb {
					log.Printf("rolling for limb loss")
					response = response + rollWanderingDMGLimbLoss()
				}

				if value.Random {
					log.Printf("rolling on the random damage table")
					response = response + rollWangeringDmgRandom()
				}

				if value.Wander {
					log.Printf("rolling on the wandering table again")
				}

				if !value.Wander {
					wander = value.Wander
				}

				if value.Damage {
					log.Printf("rolling for damage")
					rolls = roll(value.Roll.Dice, value.Roll.Value)
					damage = total(rolls)
					response = strings.Replace(value.Result, "&damage&", strconv.Itoa(damage), -1)
				}
			}
		}
	}

	discordEmbed = discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Wandering Monster Bot",
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "",
				Value: "",
			},
		},
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

func rollWangeringDmgRandom() (result string) {
	log.Printf("rolling for random damages")
	var outcome int
	var reroll = true

	for reroll {
		rolls := roll(wanderingInfo.RandomDamage.Roll.Dice, wanderingInfo.RandomDamage.Roll.Value)

		outcome = total(rolls)
		log.Printf("outcome is %d", outcome)

		log.Printf("rolling for wandering damage random subtable")
		for _, value := range wanderingInfo.RandomDamage.Table {
			if value.Outcome.Exact == outcome || between(value.Outcome.Range.Min, value.Outcome.Range.Max, outcome) {
				result = result + "\n" + value.Result
				log.Printf("%t", value.Random)
				if value.Limb {
					log.Printf("rolling for limb loss")
					result = result + rollWanderingDMGLimbLoss()
				}

				if value.Random == false {
					reroll = false
				} else {
					reroll = value.Random
				}
			}
		}

		log.Printf("result %s", result)
		log.Printf("reroll %t", reroll)
	}

	return
}
