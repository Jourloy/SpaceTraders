package agent

type Agent struct {
	AccountId      string `json:"accountId"`
	Symbol         string `json:"symbol"`
	Headquarters   string `json:"headquarters"`
	Credits        int    `json:"credits"`
	StaringFaction string `json:"staringFaction"`
}

type AgentBody struct {
	Data Agent
}
