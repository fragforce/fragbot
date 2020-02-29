package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	wanderingLute wanderingData
)

func initWanderingLute() {
	var err error

	log.Printf("loading wandering lute info")
	err = loadInfo("wandering/wandering_lute.json", &wanderingLute)
	if err != nil {
		log.Printf("%s", err)
		log.Fatalf("there was an issue reading the wandering file\n")
	}

	// log.Printf("%v", wanderingLute)

	log.Printf("wandering lute info loaded")
}

func rollWanderingLute() (response string, discordEmbed discordgo.MessageEmbed, sendToDM bool) {

	var outcome int

	var luteData wandering

	for _, data := range wanderingLute.Data {
		if data.Type == "wandering_lute" {
			luteData = data
		}
	}

	if len(luteData.Table) == 0 {
		log.Print("no loot to hand out")
		return
	}

	rolls := roll(1, len(luteData.Table))

	log.Printf("Rolls: '%d'", rolls)

	outcome = total(rolls)

	log.Printf("This is the outcome '%d'", outcome)

	// this should never happen. If it does let me know...
	if outcome == 0 {
		log.Printf("If you ever log this line please open a github issue...")
		return
	}

	for _, value := range luteData.Table {
		if value.Outcome.Exact == outcome {
			renderedResult := fmt.Sprintf("%s\n%s", value.Item, value.Result)
			response = response + "\n" + renderedResult
		}
	}

	discordEmbed = discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Wandering Lute Bot",
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
