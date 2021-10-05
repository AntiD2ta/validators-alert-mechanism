package init

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/AntiD2ta/validators-alert-mechanism/internal"
)

func Init() (config internal.Config, err error) {
	pwd, _ := os.Getwd()
	jsonFile, err := ioutil.ReadFile(pwd + "/validators.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		return
	}
	return
}
