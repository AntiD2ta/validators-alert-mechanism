package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetData(validator ValidatorData, interval int, data chan<- *Validator) {
	url := fmt.Sprintf("%s/api?module=account&action=getminedblocks&address=%s", validator.URL, validator.Address)
	rawResponse, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s\n", err)
		data <- nil
		return
	}

	var response Response
	jsonFile, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
		data <- nil
		return
	}
	err = json.Unmarshal(jsonFile, &response)
	if err != nil {
		fmt.Printf("%s\n", err)
		data <- nil
		return
	}

	if response.Message != "OK" {
		fmt.Print("Bad response")
		data <- nil
		return
	}

	layout := "2006-01-02 15:04:05.000000Z"
	date := time.Now().UTC()

	blockCount := 0
	for _, block := range response.Result {
		blockDate, err := time.Parse(layout, block.TimeStamp)
		if err != nil {
			fmt.Println(err)
			data <- nil
			return
		}

		if date.Sub(blockDate) < time.Duration(interval)*time.Minute {
			blockCount++
		} else {
			break
		}
	}

	data <- &Validator{
		Name:        validator.Name,
		MinedBlocks: blockCount,
	}
}
