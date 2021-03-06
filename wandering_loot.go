package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	wanderingLoot wanderingData
)

func initWanderingLoot() {
	var err error

	log.Printf("loading wandering loot info")
	err = loadInfo("wandering/wandering_loot.json", &wanderingLoot)
	if err != nil {
		log.Printf("%s", err)
		log.Fatalf("there was an issue reading the wandering file\n")
	}

	// log.Printf("%v", wanderingLoot)

	log.Printf("wandering info loaded")
}

func rollWanderingLoot() (response string, discordEmbed discordgo.MessageEmbed, sendToDM bool) {

	var outcome int

	var lootData wandering

	for _, data := range wanderingLoot.Data {
		if data.Type == "wandering_loot" {
			lootData = data
		}
	}

	if len(lootData.Table) == 0 {
		log.Print("no loot to hand out")
		return
	}

	rolls := roll(1, len(lootData.Table))

	log.Printf("Rolls: '%d'", rolls)

	outcome = total(rolls)

	log.Printf("This is the outcome '%d'", outcome)

	// this should never happen. If it does let me know...
	if outcome == 0 {
		log.Printf("If you ever log this line please open a github issue...")
		return
	}

	for _, value := range lootData.Table {
		if value.Outcome.Exact == outcome {
			renderedResult := fmt.Sprintf("You have attained one '%s', with the following description: %s", value.Item, value.Result)
			response = response + "\n" + renderedResult
		}
	}

	discordEmbed = discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Wandering Loot Bot",
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
