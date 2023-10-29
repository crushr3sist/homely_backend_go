package convert

import (
	"encoding/json"
	"fmt"
	"os"
)

type ShowEpisodes struct {
	Episodes []string `json:"converted"`
}

// Define a struct for the entire JSON structure
type ConvertedData struct {
	Converted map[string]ShowEpisodes `json:"converted"`
}

func filterPathList(paths []string) []string {
	return []string{}
}

func readConvertedJson() ConvertedData {

	file, err := os.Open("../videos/converted.json")
	if err != nil {
		os.Create("../videos/converted.json")
		fmt.Println("error opening the file", err)
	} else {

		defer file.Close()
	}

	decoder := json.NewDecoder(file)

	var readData ConvertedData

	if err := decoder.Decode(&readData); err != nil {
		fmt.Println("error decoding json", err)
	}

	return readData
}
