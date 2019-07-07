package main

type dndFiveECharacterInfo struct {
	ID      int                        `json:"id"`
	Player  pcPlayerInfo               `json:"player_info"`
	Stats   dndFiveECharacterStats     `json:"stats"`
	EXP     int                        `json:"experience_points,omitempty"`
	HP      dndFiveECharacterHitPoints `json:"hit_points"`
	Armor   int                        `json:"armor_class"`
	Init    int                        `json:"initiative"`
	Speed   dndFiveESpeed              `json:"speed"`
	Skills  dndFiveECharacterSkills    `json:"skills"`
	PassWis int                        `json:"passive_wisdom"`
	Details dndFiveECharacterDetails   `json:"details"`
}

type dndFiveECharacterDetails struct {
	Name      dndFiveECharacterName `json:"name"`
	RaceID    int                   `json:"race_id"`
	Alignment dndFiveEAlignment     `json:"alignment"`
	Age       int                   `json:"age,omitmempty"`
	Height    dndFiveEHeight        `json:"height"`
	Weight    dndFiveEWeight        `json:"weignt"`
	HairColor string                `json:"hair_color"`
	EyeColor  string                `json:"eye_color"`
	SkinColor string                `json:"skin_color"`
	LangIDs   []int                 `json:"language_ids"`
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
	Acrobatics    bool `json:"acrobatics,omitempty"`
	AnimalHandle  bool `json:"animal_handling,omitempty"`
	Arcane        bool `json:"arcane,omitempty"`
	Athletics     bool `json:"athletics,omitempty"`
	Deception     bool `json:"deception,omitempty"`
	History       bool `json:"history,omitempty"`
	Insight       bool `json:"insight,omitempty"`
	Intimidation  bool `json:"intimidation,omitempty"`
	Investigation bool `json:"investigation,omitempty"`
	Medicine      bool `json:"medicine,omitempty"`
	Nature        bool `json:"nature,omitempty"`
	Perception    bool `json:"percetption,omitempty"`
	Performance   bool `json:"performance,omitempty"`
	Persuasion    bool `json:"persuasion,omitempty"`
	Religion      bool `json:"religion,omitempty"`
	Sleight       bool `json:"sleight_of_hand,omitempty"`
	Stealth       bool `json:"stealth,omitempty"`
	Survival      bool `json:"survival,omitempty"`
}

type dndFiveECharacterClasses struct {
	Classes []dndFiveECharacterClass `json:"classes"`
}

type dndFiveECharacterClass struct {
	ClassID   int    `json:"id"`
	ClaccName string `json:"class_name"`
}

type dndFiveECharacterEquipment struct {
}

type dndFiveERaces struct {
	Races []dndFiveERace `json:"races"`
}

type dndFiveERace struct {
	RaceID    int                 `json:"id"`
	RaceName  string              `json:"race_name"`
	AdultAge  int                 `json:"adult_age"`
	Alignment dndFiveEAlignment   `json:"alignment"`
	Size      int                 `json:"size_id"`
	Speed     dndFiveESpeed       `json:"speed"`
	Height    dndFiveEHeightRange `json:"height,omitempty"`
	Length    dndFiveEHeightRange `json:"length,omitempty"`
	Weight    dndFiveEWeightRange `json:"weight"`
	LangIDs   []int               `json:"language_ids"`
}

type dndFiveEAlignment struct {
	// Good vs Evil
	GVE int `json:"gve"`
	// Lawful vs Chaos
	LVC int `json:"lvc"`
}

type dndFiveESizes struct {
	Sizes []dndFiveESize `json:"sizes"`
}

type dndFiveESize struct {
	SizeName int                 `json:"size_name"`
	ACMod    int                 `json:"ac_modifier"`
	SPCMod   int                 `json:"special_modifier"`
	HideMod  int                 `json:"hide_modifier"`
	Height   dndFiveEHeightRange `json:"height,omitempty"`
	Length   dndFiveEHeightRange `json:"length,omitempty"`
	Weight   dndFiveEWeightRange `json:"weight"`
	Space    int
}

type dndFiveEHeightRange struct {
	Max dndFiveEHeight `json:"max,omitempty"`
	Min dndFiveEHeight `json:"min,omitempty"`
}

type dndFiveEWeightRange struct {
	Max dndFiveEWeight `json:"max,omitempty"`
	Min dndFiveEWeight `json:"min,omitempty"`
}

type dndFiveEHeight struct {
	Feet   int `json:"feet,omitempty"`
	Inches int `json:"inches,omitempty"`
	CM     int `json:"centimeters,omitempty"`
}

type dndFiveEWeight struct {
	Pounds int `json:"pounds,omitempty"`
	Ounces int `json:"ounces,omitempty"`
	Grams  int `json:"grams,omitempty"`
}

type dndFiveELanguages struct {
	Languages []dndFiveELanguage `json:"languages"`
}

type dndFiveELanguage struct {
	LangID   int    `json:"id"`
	LangName string `json:"lang_name"`
}

type dndFiveESpeed struct {
	Feet int `json:"feet,omitempty"`
	CM   int `json:"centimeters,omitempty"`
}
