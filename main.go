package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "custom!"
}

type AirQuality struct {
	MapVersion       string `json:"mapVersion"`
	BaseVersion      string `json:"baseVersion"`
	MapVersionString string `json:"mapVersionString"`
	Results          []struct {
		ID                           int     `json:"ID"`
		Label                        string  `json:"Label"`
		DEVICELOCATIONTYPE           string  `json:"DEVICE_LOCATIONTYPE,omitempty"`
		THINGSPEAKPRIMARYID          string  `json:"THINGSPEAK_PRIMARY_ID"`
		THINGSPEAKPRIMARYIDREADKEY   string  `json:"THINGSPEAK_PRIMARY_ID_READ_KEY"`
		THINGSPEAKSECONDARYID        string  `json:"THINGSPEAK_SECONDARY_ID"`
		THINGSPEAKSECONDARYIDREADKEY string  `json:"THINGSPEAK_SECONDARY_ID_READ_KEY"`
		Lat                          float64 `json:"Lat"`
		Lon                          float64 `json:"Lon"`
		PM25Value                    string  `json:"PM2_5Value"`
		LastSeen                     int     `json:"LastSeen"`
		Type                         string  `json:"Type,omitempty"`
		Hidden                       string  `json:"Hidden"`
		DEVICEBRIGHTNESS             string  `json:"DEVICE_BRIGHTNESS,omitempty"`
		DEVICEHARDWAREDISCOVERED     string  `json:"DEVICE_HARDWAREDISCOVERED,omitempty"`
		DEVICEFIRMWAREVERSION        string  `json:"DEVICE_FIRMWAREVERSION,omitempty"`
		Version                      string  `json:"Version,omitempty"`
		LastUpdateCheck              int     `json:"LastUpdateCheck,omitempty"`
		Created                      int     `json:"Created"`
		Uptime                       string  `json:"Uptime,omitempty"`
		RSSI                         string  `json:"RSSI,omitempty"`
		Adc                          string  `json:"Adc,omitempty"`
		P03Um                        string  `json:"p_0_3_um"`
		P05Um                        string  `json:"p_0_5_um"`
		P10Um                        string  `json:"p_1_0_um"`
		P25Um                        string  `json:"p_2_5_um"`
		P50Um                        string  `json:"p_5_0_um"`
		P100Um                       string  `json:"p_10_0_um"`
		Pm10Cf1                      string  `json:"pm1_0_cf_1"`
		Pm25Cf1                      string  `json:"pm2_5_cf_1"`
		Pm100Cf1                     string  `json:"pm10_0_cf_1"`
		Pm10Atm                      string  `json:"pm1_0_atm"`
		Pm25Atm                      string  `json:"pm2_5_atm"`
		Pm100Atm                     string  `json:"pm10_0_atm"`
		IsOwner                      int     `json:"isOwner"`
		Humidity                     string  `json:"humidity,omitempty"`
		TempF                        string  `json:"temp_f,omitempty"`
		Pressure                     string  `json:"pressure,omitempty"`
		AGE                          int     `json:"AGE"`
		Stats                        string  `json:"Stats"`
		ParentID                     int     `json:"ParentID,omitempty"`
	} `json:"results"`
}

func NewAirQuality() (*AirQuality, error) {
	aq := &AirQuality{}
	err := aq.fetchBeachAirQuality()
	return aq, err
}

func (m *AirQuality) fetchBeachAirQuality() error {
	client := http.Client{Timeout: 45 * time.Second}
	resp, err := client.Get("https://www.purpleair.com/json?show=54507")
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, m)
	return err
}

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "airquality",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(basic)
	app.Bind(NewAirQuality)
	app.Run()
}
