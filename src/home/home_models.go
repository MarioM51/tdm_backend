package home

import (
	"encoding/json"
	"io/ioutil"
	"users_api/src/errorss"
)

type LocalBusinessLDJSON struct {
	Context                   string                      `json:"@context"`
	Type                      string                      `json:"@type"`
	Image                     []string                    `json:"image"`
	Name                      string                      `json:"name"`
	Founder                   string                      `json:"founder"`
	Keywords                  string                      `json:"keywords"`
	Description               string                      `json:"description"`
	URL                       string                      `json:"url"`
	Telephone                 string                      `json:"telephone"`
	Email                     string                      `json:"email"`
	Address                   Address                     `json:"address"`
	Geo                       Geo                         `json:"geo"`
	OpeningHoursSpecification []OpeningHoursSpecification `json:"openingHoursSpecification"`
}

type Address struct {
	Type            string `json:"@type"`
	StreetAddress   string `json:"streetAddress"`
	AddressLocality string `json:"addressLocality"`
	AddressRegion   string `json:"addressRegion"`
	PostalCode      string `json:"postalCode"`
	AddressCountry  string `json:"addressCountry"`
}

type Geo struct {
	Type      string  `json:"@type"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OpeningHoursSpecification struct {
	Type      string   `json:"@type"`
	DayOfWeek []string `json:"dayOfWeek"`
	Opens     string   `json:"opens"`
	Closes    string   `json:"closes"`
}

func (LocalBusinessLDJSON) readFromLocalFile(localInfo *LocalBusinessLDJSON) {
	content, err := ioutil.ReadFile("./data/LocalBusinessLDJSON.json")
	if err != nil {
		panic(errorss.ErrorResponseModel{
			HttpStatus: 500,
			Cause:      "Error reading the local business info: " + err.Error()},
		)
	}

	err = json.Unmarshal(content, &localInfo)
	localInfo.URL = constants.Domain
	if err != nil {
		panic(errorss.ErrorResponseModel{
			HttpStatus: 500,
			Cause:      "Error unmarshalling the local business info: " + err.Error()},
		)
	}

}
