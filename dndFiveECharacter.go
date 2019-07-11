package main

type dndFiveECharacterInfo struct {
	ID      int                        `json:"id"`
	Player  pcPlayerInfo               `json:"player_info"`
	Stats   dndFiveECharacterStats     `json:"stats"`
	Purse   dndFiveEPurse              `json:"purse,omitempty"`
	Equip   dndFiveECharacterEquipment `json:"equipment"`
	Details dndFiveECharacterDetails   `json:"details"`
}

type dndFiveECharacterStats struct {
	EXP       int                        `json:"experience_points,omitempty"`
	HP        dndFiveECharacterHitPoints `json:"hit_points"`
	Ability   dndFiveEStatAbilities      `json:"abilities"`
	Armor     int                        `json:"armor_class"`
	Init      int                        `json:"initiative"`
	Speed     dndFiveESpeed              `json:"speed"`
	InspPoint int                        `json:"inspiration_points,omitmempty"`
	ProfBonus int                        `json:"proficiency_bonus,omitmempty"`
	PassWis   int                        `json:"passive_wisdom,omitmempty"`
	Skills    dndFiveECharacterSkills    `json:"skills"`
}

type dndFiveECharacterHitPoints struct {
	Max     int `json:"max"`
	Current int `json:"current"`
	Temp    int `json:"temp,omitempty"`
}

type dndFiveEStatAbilities struct {
	Str dndFiveEStatAbility `json:"strength,omitempty"`
	Dex dndFiveEStatAbility `json:"dexterity,omitempty"`
	Con dndFiveEStatAbility `json:"constitution,omitempty"`
	Int dndFiveEStatAbility `json:"intelligence,omitempty"`
	Wis dndFiveEStatAbility `json:"wisdom,omitempty"`
	Cha dndFiveEStatAbility `json:"charisma,omitempty"`
}

type dndFiveEStatAbility struct {
	AblScore   int  `json:"score"`
	Proficient bool `json:"proficient,omitempty"`
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

	Other []dndFiveECharacterOtherSkills `json:"other,omitempty"`
}

type dndFiveECharacterOtherSkills struct {
	ID         int    `json:"id"`
	SkillName  string `json:"skill_name,omitempty"`
	Descrption string `json:"description,omitempty"`
	Base       string `json:"base_stat"`
	Trained    bool   `json:"trained,omitempty"`
}

// TODO: add json structs to these
type dndFiveECharacterEquipment struct {
	Name        string
	Quantity    int
	Description string
	Heavy       bool
	Light       bool
	Loading     bool
	Ranged      bool
	Reach       bool
	Special     bool
	Thrown      bool
	Versatile   bool
	Improvised  bool
	Silvered    bool
	Enchanted   bool
}

type dndFiveECharacterDetails struct {
	Name      dndFiveECharacterName `json:"name"`
	Alignment dndFiveEAlignment     `json:"alignment"`
	RaceID    int                   `json:"race_id"`
	Age       int                   `json:"age,omitmempty"`
	Height    dndFiveEHeight        `json:"height"`
	Weight    dndFiveEWeight        `json:"weignt"`
	HairColor string                `json:"hair_color"`
	EyeColor  string                `json:"eye_color"`
	SkinColor string                `json:"skin_color"`
	LangIDs   []int                 `json:"language_ids"`
	Backstory []string              `json:"backstory,omitmempty"`
	Ideals    []string              `json:"ideals,omitmempty"`
	Bonds     []string              `json:"bonds,omitmempty"`
	Flaws     []string              `json:"flaws,omitmempty"`
}

type dndFiveECharacterName struct {
	Honorific  string `json:"honorific,omitempty"`
	Title      string `json:"title,omitempty"`
	FirstName  string `json:"first,omitempty"`
	MiddleName string `json:"middle,omitempty"`
	LastName   string `json:"last,omitempty"`
}

type dndFiveECharacterClasses struct {
	Classes []dndFiveECharacterClass `json:"classes"`
}

type dndFiveECharacterClass struct {
	ClassID   int    `json:"id"`
	ClaccName string `json:"class_name"`
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

type dndFiveEPurse struct {
	Plat   int `json:"platinum,omitmempty"`
	Gold   int `json:"gold,omitmempty"`
	Silver int `json:"silver,omitmempty"`
	Copper int `json:"copper,omitmempty"`
}
