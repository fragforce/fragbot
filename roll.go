package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//                ____      __  __                   ___
//    _________  / / /     / /_/ /_  ___        ____/ (_)_______
//   / ___/ __ \/ / /_____/ __/ __ \/ _ \______/ __  / / ___/ _ \
//  / /  / /_/ / / /_____/ /_/ / / /  __/_____/ /_/ / / /__/  __/
// /_/   \____/_/_/      \__/_/ /_/\___/      \__,_/_/\___/\___/

type rtdInfo struct {
	ChannelID string        `json:"channel_id,omitempty"`
	Sides     []int         `json:"sides,omitempty"`
	Wandering rollWandering `json:"wandering,omitempty"`
}

type rollOutcome struct {
	Exact int              `json:"exact,omitempty"`
	Range rollOutcomeRange `json:"range,omitempty"`
}

type rollOutcomeRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type rollWandering struct {
	Damage wanderingSettings `json:"damage,omitempty"`
	Loot   wanderingSettings `json:"loot,omitempty"`
	Lute   wanderingSettings `json:"lute,omitempty"`
}

func rollTheDiceInit() {
}

func getSeed() (int64, error) {
	c := 8
	b := make([]byte, c)
	_, err := crand.Read(b)
	if err != nil {
		return 0, err
	}
	return (int64)(binary.BigEndian.Uint64(b)), nil
}

func rollHandler(messageContent string) (response string, discordEmbed discordgo.MessageEmbed, sendToDM bool) {
	if strings.TrimPrefix(messageContent, chn.Prefix+"roll") == "" || !strings.HasPrefix(messageContent, chn.Prefix+"roll ") {
		response = "How to use Roll the Dice\n`!roll (dice)d(sides)[+/-][proficiency]x[multiple times]` I.E. `!roll 1d20+3x5`\n `!roll stats` to roll a base set of stats."
		return
	}

	switch strings.TrimPrefix(messageContent, chn.Prefix+"roll ") {
	case "wandering dmg":
		log.Printf("rolling wandering damage")
		response, discordEmbed, sendToDM = rollWanderingDamage()
	case "wandering loot":
		log.Printf("rolling wandering loot")
		response, discordEmbed, sendToDM = rollWanderingLoot()
	case "wandering lute":
		log.Printf("rolling wandering lute")
		response, discordEmbed, sendToDM = rollWanderingLute()
	case "stats":
		log.Printf("stats")
		response, sendToDM = rollStats("")
	default:
		log.Printf("rolling the dice")
		response, sendToDM = rollTheDice(strings.TrimPrefix(messageContent, chn.Prefix+"roll "))
	}

	return
}

func rollTheDice(message string) (response string, sendToDM bool) {
	var err error

	// a users proficiency
	var proficiency int

	// if a roll is to be run multiple times
	multiRoll := 1

	log.Printf("roll the dice")
	// Example !roll 1d6+2
	validID, err := regexp.Compile(`(\d+)\s?d\s?(\d+)\s?(?:(\+|\-)\s?(\d*))?(?:\s?(?:x\s?)(\d*)|)`)
	if err != nil {
		log.Printf("There was an error compiling the regex for the roll command")
		return
	}

	dieInfo := validID.FindStringSubmatch(message)

	if len(dieInfo) == 0 {
		return
	}

	rollCount, err := strconv.Atoi(dieInfo[1])
	if err != nil {
		log.Printf("There was an error converting the number of rolls")
	}

	dieValue, err := strconv.Atoi(dieInfo[2])
	if err != nil {
		log.Printf("There was an error converting the number of sides")
	}

	if dieInfo[4] != "" {
		proficiency, err = strconv.Atoi(dieInfo[4])
		if err != nil {
			log.Printf("There was an error converting proficiency")
		}
	}

	if !hasElem(chn.RTD.Sides, dieValue) {
		log.Printf("The amount of sides %d, is not supported", dieValue)
		response = fmt.Sprintf("Only dice with %s sides are supported.", arrayToString(chn.RTD.Sides))
		return
	}

	if rollCount > 10 {
		response = fmt.Sprintf("rolls are limited to 10 at a time")
		return
	}

	if dieInfo[5] != "" {
		// log.Printf("rolling %s sets", dieInfo[5])
		multiRoll, err = strconv.Atoi(dieInfo[5])
		if err != nil {
			log.Printf("There was an error converting the number of rolls")
		}
		response = fmt.Sprintf("I have rolled %d sets of rolls for you coming out with \n", multiRoll)
	}

	if multiRoll > 5 {
		response = fmt.Sprintf("Sorry I only support up to 5 sets of rolls.")
		return
	}

	for i := 1; i <= multiRoll; i++ {
		response = response + rollDie(dieInfo[3], dieValue, rollCount, proficiency)
	}

	return
}

func roll(rollCount int, dieValue int) (rolls []int) {
	seed, err := getSeed()
	if err != nil {
		log.Print("Error generating seed")
	}
	rand.Seed(seed)

	for i := 0; i < rollCount; i++ {
		rolls = append(rolls, rand.Intn(dieValue)+1)
	}

	// log.Printf("%d", rolls)
	return
}

func flipCoin() (response string, sendToDM bool) {
	side := total(roll(1, 2))

	// log.Printf("side = %d", side)

	if side == 1 {
		response = "I have flipped the coin getting a heads"
	} else if side == 2 {
		response = "I have flipped the coin getting a tails"
	}
	return
}

func rollDie(addSub string, dieValue, rollCount, proficiency int) (response string) {
	// strings that are sent back
	var prettyRolls string
	var profString string

	// log.Printf("rolling a %d sided die %d times", dieValue, rollCount)
	allRolls := roll(rollCount, dieValue)
	prettyRolls = arrayToString(allRolls)

	rollTotal := total(allRolls)

	// log.Printf("roll total = %d", rollTotal)

	if addSub == "" {
		// log.Printf("No profeciency was added to the roll")
	} else {
		if addSub == "+" {
			log.Printf("Adding %d to the roll", proficiency)
			rollTotal = rollTotal + proficiency
			profString = fmt.Sprintf("adding %d ", proficiency)
		} else if addSub == "-" {
			log.Printf("subtracting %d to the roll", proficiency)
			rollTotal = rollTotal - proficiency
			profString = fmt.Sprintf("subtracting %d ", proficiency)
		} else {

		}
	}

	if rollCount == 1 {
		response = fmt.Sprintf("You have rolled a %s \n", prettyRolls)
	} else {
		response = fmt.Sprintf("You have rolled %s %sfor a total of %d \n", prettyRolls, profString, rollTotal)
	}

	return
}

func rollStats(message string) (response string, sendToDM bool) {
	log.Printf("rolling stats for the user")

	var allRolls []int
	var rerollOne bool

	if message == "" {
		rerollOne = true
	}

	for len(allRolls) < 6 {
		// log.Printf("roll set %d", len(allRolls)+1)
		allRolled := []int{}
		for len(allRolled) < 4 {
			rolled := total(roll(1, 6))
			for rolled <= 1 && rerollOne {
				// log.Printf("rerolling a 1 that was rolled")
				rolled = total(roll(1, 6))
			}
			allRolled = append(allRolled, rolled)
		}

		// log.Printf("%d", allRolled)

		sort.Ints(allRolled)

		// log.Printf("%d", allRolled)

		allRolled = allRolled[1:]

		// log.Printf("%d", allRolled)

		allRolls = append(allRolls, total(allRolled))
	}

	log.Printf("all roll set totals '%d'", allRolls)

	response = fmt.Sprintf("I have rolled the dice and return the following stat rolls for you. %d", allRolls)

	return
}

func total(dice []int) (total int) {
	for _, die := range dice {
		total = total + die
	}

	// log.Printf("total = %d", total)
	return
}

func between(min, max, num int) (isBetween bool) {
	if max >= num && num >= min {
		isBetween = true
	}

	return
}

func arrayToString(intArray []int) (pretty string) {
	for rtdi, val := range intArray {
		pretty = pretty + strconv.Itoa(val)
		if rtdi == len(intArray)-2 {
			pretty = pretty + ", and "
		} else if rtdi != len(intArray)-1 {
			pretty = pretty + ", "
		}
	}

	return
}

// if array has an element
func hasElem(array interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(array)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}
