package main

type wandering struct {
	Type  string               `json:"type"`
	Roll  wanderingRoll        `json:"roll"`
	Table []wanderingTableItem `json:"table"`
}

type wanderingSettings struct {
	Enabled bool `json:"enabled,omitempty"`
}

type wanderingData struct {
	Data []wandering `json:"data"`
}

type wanderingRoll struct {
	Dice  int `json:"dice"`
	Value int `json:"value"`
}

type wanderingTableItem struct {
	Item string `json:"item,omitempty"`
	// Outcome is the value for the items in the list.
	Outcome rollOutcome `json:"outcome,omitempty"`
	// A string response for describing what happened or what was received
	Result string `json:"result"`
	// what to roll when calculateing wandering items.
	Roll wanderingRoll `json:"roll,omitempty"`
	// these are effects that are applied in the roll process
	// Limb is if the limb loss subtable is to be rolled against.
	Limb bool `json:"limb,omitempty"`
	// Wander is if the wandering damage table is to be rolled.
	Wander bool `json:"wander,omitempty"`
	// random is if the random damage table is to be rolled.
	Random bool `json:"random,omitempty"`
	// damage is if there is a roll for damage as part of the outcome.
	Damage bool `json:"damage,omitempty"`
}
