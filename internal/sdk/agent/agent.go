package agent

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

func GetAgent() Agent {
	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found")
	}

	uri := os.Getenv("SPACE_TRADERS_API")
	if uri == "" {
		log.Fatal("You must set your 'SPACE_TRADERS_API' environmental variable.")
	}

	httpRequest, err := http.NewRequest(http.MethodGet, uri + "/my/agent", nil)
	if err != nil {
		log.Fatal("Cannot create request")
	}

	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SPACE_TRADERS_TOKEN")))

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Fatal("Cannot get agent")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(`Body close error`)
		}
	} (resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Cannot read agent")
	}

	agent := &AgentBody{}
	err = json.Unmarshal(body, agent)
	if err != nil {
		log.Fatal("Cannot unmarshal body")
	}

	return agent.Data
}