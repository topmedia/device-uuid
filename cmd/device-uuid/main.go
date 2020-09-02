package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

type Settings struct {
	PrivacyMode                         bool   `json:"privacyMode"`
	EnableIncomingSupport               bool   `json:"enableIncomingSupport"`
	DeviceUID                           string `json:"deviceUID"`
	URI                                 string `json:"uri"`
	AllowConnectionWhenNobodyIsLoggedIn bool   `json:"allowConnectionWhenNobodyIsLoggedIn"`
	AccountUID                          string `json:"accountUID"`
	SiteUID                             string `json:"siteUID"`
	NotificationURI                     string `json:"notificationUri"`
}

func main() {
	var (
		settingsFile string
		settings     Settings
	)

	if runtime.GOOS == "darwin" {
		settingsFile = "/usr/local/share/CentraStage/AEMAgent/Settings.json"
	} else {
		settingsFile = "C:/ProgramData/CentraStage/AEMAgent/Settings.json"
	}

	jsonFile, err := os.Open(settingsFile)
	if err != nil {
		log.Fatalf("Error opening JSON: %s", err)
	}

	defer jsonFile.Close()

	jsonValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error reading JSON: %s", err)
	}

	json.Unmarshal(jsonValue, &settings)

	fmt.Println(settings.DeviceUID)
}
