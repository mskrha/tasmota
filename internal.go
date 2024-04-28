package tasmota

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const timeFormat = "2006-01-02T15:04:05-07:00"

func (t *Tasmota) getCommand(u string) (ret []byte, err error) {
	res, err := http.Get(fmt.Sprintf("http://%s/cm?cmnd=%s", t.host, url.PathEscape(u)))
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Device responded with HTTP status code %d", res.StatusCode)
		return
	}
	ret, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	return
}

type inputStatus8 struct {
	Status struct {
		Time   string `json:"Time"`
		Energy struct {
			Current        float64 `json:"Current"`
			Factor         float64 `json:"Factor"`
			Voltage        uint64  `json:"Voltage"`
			Today          float64 `json:"Today"`
			Yesterday      float64 `json:"Yesterday"`
			Total          float64 `json:"Total"`
			TotalStartTime string  `json:"TotalStartTime"`
		} `json:"ENERGY"`
	} `json:"StatusSNS"`
}

func (t *Tasmota) getStatus8() (ret inputStatus8, err error) {
	in, err := t.getCommand("Status 8")
	if err != nil {
		return
	}
	err = json.Unmarshal(in, &ret)
	return
}

type inputState struct {
	Time          string `json:"Time"`
	Uptime        string `json:"Uptime"`
	UptimeSeconds uint64 `json:"UptimeSec"`
	PowerOutput   string `json:"POWER"`
	WiFi          struct {
		BSSID   string `json:"BSSId"`
		Channel uint64 `json:"Channel"`
		Mode    string `json:"Mode"`
		RSSI    uint64 `json:"RSSI"`
		Signal  int64  `json:"Signal"`
	} `json:"Wifi"`
}

func (t *Tasmota) getState() (ret inputState, err error) {
	in, err := t.getCommand("State")
	if err != nil {
		return
	}
	err = json.Unmarshal(in, &ret)
	return
}
