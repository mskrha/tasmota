package tasmota

import (
	"fmt"
	"time"
)

type Tasmota struct {
	host string
}

func New(h string) (*Tasmota, error) {
	if len(h) == 0 {
		return nil, fmt.Errorf("Hostname must not be empty")
	}
	var t Tasmota
	t.host = h
	return &t, nil
}

type DataEnergy struct {
	Voltage   uint64
	Power     float64
	Today     float64
	Yesterday float64
	Total     float64
}

func (t *Tasmota) GetEnergy() (ret DataEnergy, err error) {
	s, err := t.getStatus8()
	if err != nil {
		return
	}
	ret.Voltage = s.Status.Energy.Voltage
	ret.Power = float64(s.Status.Energy.Voltage) * s.Status.Energy.Current * s.Status.Energy.Factor
	ret.Today = s.Status.Energy.Today
	ret.Yesterday = s.Status.Energy.Yesterday
	ret.Total = s.Status.Energy.Total
	return
}

type DataWiFi struct {
	BSSID   string `json:"BSSID"`
	Channel uint64 `json:"Channel"`
	Mode    string `json:"Mode"`
	RSSI    uint64 `json:"RSSI"`
	Signal  int64  `json:"Signal"`
}

type DataStatus struct {
	CurrentTime time.Time `json:"CurrentTime"`
	Uptime      uint64    `json:"Uptime"`
	Output      bool      `json:"Output"`
	WiFi        DataWiFi  `json:"WiFi"`
}

func (t *Tasmota) GetStatus() (ret DataStatus, err error) {
	s, err := t.getState()
	if err != nil {
		return
	}
	ret.CurrentTime, err = time.Parse(timeFormat, s.Time)
	if err != nil {
		return
	}
	ret.Uptime = s.UptimeSeconds
	switch s.PowerOutput {
	case "ON":
		ret.Output = true
	case "OFF":
		ret.Output = false
	default:
		err = fmt.Errorf("Invalid power output state")
		return
	}
	ret.WiFi.BSSID = s.WiFi.BSSID
	ret.WiFi.Channel = s.WiFi.Channel
	ret.WiFi.Mode = s.WiFi.Mode
	ret.WiFi.RSSI = s.WiFi.RSSI
	ret.WiFi.Signal = s.WiFi.Signal
	return
}
