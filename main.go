package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Title    string    `json:"title"`
	Addr     Address   `json:"address"`
	Position Positions `json:"position"`
}

type Address struct {
	Label string `json:"label"`
}

type Positions struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

const (
	apiKey = "Skg1vxohQl5I32LLn_7tWUg6FOxpmzx56Vnqnw5TOnA"
	URL    = "https://discover.search.hereapi.com/v1/discover"
)

func WriteFile(data []byte) error {
	err := os.WriteFile("./map.geojson", data, 64)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	resp, err := http.Get("https://discover.search.hereapi.com/v1/discover?in=circle:21.01035,105.80826;r=100000&q=aha+coffe&limit=10&apiKey=Skg1vxohQl5I32LLn_7tWUg6FOxpmzx56Vnqnw5TOnA")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	features := []interface{}{}

	data := Items{}
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		log.Fatalln(err1)
	}
	for _, val := range data.Items {
		coordinates := []interface{}{}
		coordinates = append(coordinates, val.Position.Lat)
		coordinates = append(coordinates, val.Position.Lng)

		geojsonFeature := map[string]interface{}{
			"type": "Feature",
			"properties": map[string]interface{}{
				"name":     val.Title,
				"location": val.Addr.Label,
			},
			"geometry": map[string]interface{}{
				"type":        "Point",
				"coordinates": coordinates,
			},
		}
		features = append(features, geojsonFeature)
	}
	result := map[string]interface{}{
		"type":     "FeatureCollection",
		"features": features,
	}

	endData, _ := json.Marshal(result)
	WriteFile(endData)
	fmt.Println("Success")

}
