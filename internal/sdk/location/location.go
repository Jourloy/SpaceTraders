package location

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetLocation(headquarters string) Location {
	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found")
	}

	uri := os.Getenv("SPACE_TRADERS_API")
	if uri == "" {
		log.Fatal("You must set your 'SPACE_TRADERS_API' environmental variable.")
	}

	splited := strings.Split(headquarters, "-")
	// sector := splited[0]
	system := fmt.Sprintf("%s-%s", splited[0], splited[1])
	waypoint := fmt.Sprintf("%s-%s-%s", splited[0], splited[1], splited[2])

	httpRequest, err := http.NewRequest(http.MethodGet, uri+fmt.Sprintf("/systems/%s/waypoints/%s", system, waypoint), nil)
	if err != nil {
		log.Fatal("Cannot create request")
	}

	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SPACE_TRADERS_TOKEN")))

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Fatal("Cannot get location")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(`Body close error`)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Cannot read location")
	}

	location := &LocationBody{}
	err = json.Unmarshal(body, location)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot unmarshal body | Err: %s", err))
	}

	return location.Data
}
