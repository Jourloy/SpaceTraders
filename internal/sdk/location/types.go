package location

type Location struct {
	SystemSymbol string    `json:"systemSymbol"`
	Symbol       string    `json:"symbol"`
	Type         string    `json:"type"`
	X            int       `json:"x"`
	Y            string    `json:"y"`
	Orbitals     []Orbital `json:"orbitals"`
	Traits       []Trait   `json:"traits"`
	Chart        Chart     `json:"chart"`
	Faction      Faction   `json:"faction"`
}

type Orbital struct {
	Symbol string `json:"symbol"`
}

type Trait struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Chart struct {
	SubmittedBy string `json:"submittedBy"`
	SubmittedOn string `json:"submittedOn"`
}

type Faction struct {
	Symbol string `json:"symbol"`
}
