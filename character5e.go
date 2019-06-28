package main

type dndFiveECharacterInfo struct {
	ID    int                        `json:"id"`
	Name  dndFiveECharacterName      `json:"name"`
	EXP   int                        `json:"experience_points,omitempty"`
	HP    dndFiveECharacterHitPoints `json:"hit_points"`
	Armor int                        `json:"armor_class"`
	Init  int                        `json:"initiative"`
	Speed int                        `json:"speed"`
	Stats dndFiveECharacterStats     `json:"stats"`
}

type dndFiveECharacterName struct {
	Honorific  string `json:"honorific,omitempty"`
	Title      string `json:"title,omitempty"`
	FirstName  string `json:"first,omitempty"`
	MiddleName string `json:"middle,omitempty"`
	LastName   string `json:"last,omitempty"`
}

type dndFiveECharacterHitPoints struct {
	Max     int `json:"max"`
	Current int `json:"current"`
	Temp    int `json:"temp,omitempty"`
}

type dndFiveECharacterStats struct {
	Str int `json:"strength,omitempty"`
	Dex int `json:"dexterity,omitempty"`
	Con int `json:"constitution,omitempty"`
	Int int `json:"intelligence,omitempty"`
	Wis int `json:"wisdom,omitempty"`
	Cha int `json:"charisma,omitempty"`
}

type dndFiveECharacterSkills struct {
}

type dndFiveECharacterClass struct {
}

type dndFiveECharacterRace struct {
}

type dndFiveECharacterDetails struct {
}
