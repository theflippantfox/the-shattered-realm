package character

type Character struct {
	id        string
	name      string
	rating    string //A/B/C/D or E
	skills    []Skill
	equipment []Equipment
}

type Skill struct {
	id          string
	name        string
	description string
	level       int
	xp          int
	nextLevel   int
}

type Equipment struct {
	id           string
	name         string
	description  string
	effects      []string
	bonusEffects []string
}

func CreateCharacter() (Character, error) {
	return Character{}, nil
}
