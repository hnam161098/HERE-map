package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"map/model"
	"net/http"
	"os"
)

func WriteFile(data []byte) error {
	err := os.WriteFile("./map.geojson", data, 64)
	if err != nil {
		return err
	}
	return nil
}

func GetData(url, locate, radius, keyword, limit, key string) (model.Items, error) {
	var data model.Items
	params := fmt.Sprintf("%v?in=circle:%v;r=%v&q=%v&limit=%v&apikey=%v", url, locate, radius, keyword, limit, key)
	resp, err := http.Get(params)
	if err != nil {
		return model.Items{}, err
	}
	body, errRead := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Items{}, errRead
	}
	errData := json.Unmarshal(body, &data)
	if err != nil {
		log.Println(errData)
	}
	return data, nil
}

func GenerateGeojson() {
	url := model.URL
	locate := "21.01035,105.80826"
	radius := "1000000"
	keyword := "coffee"
	limit := "100"
	key := model.Key

	data, err := GetData(url, locate, radius, keyword, limit, key)
	if err != nil {
		log.Fatal(err)
	}
	var features []interface{}
	for _, v := range data.Items {
		coordinates := []interface{}{}
		coordinates = append(coordinates, v.Position.Lng, v.Position.Lat)

		geojsonFeature := map[string]interface{}{
			"type": "Feature",
			"properties": map[string]interface{}{
				"name":     v.Title,
				"location": v.Addr.Label,
				"city":     v.Addr.Country,
				"district": v.Addr.City,
				"street":   v.Addr.Street,
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

	dataWrite, _ := json.Marshal(result)
	WriteFile(dataWrite)
	log.Println("Done.")

}

func main() {

}
