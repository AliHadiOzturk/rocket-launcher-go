package models

import (
	"encoding/json"
)

func UnmarshalRocket(data []byte) (Rocket, error) {
	var r Rocket
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Rocket) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Rocket struct {
	ID           string     `json:"id"`
	Model        string     `json:"model"`
	Mass         float32    `json:"mass"`
	Payload      Payload    `json:"payload"`
	Telemetry    Telemetry  `json:"telemetry"`
	Status       string     `json:"status"`
	Timestamps   Timestamps `json:"timestamps"`
	Altitude     float32    `json:"altitude"`
	Speed        float32    `json:"speed"`
	Acceleration float32    `json:"acceleration"`
	Thrust       float32    `json:"thrust"`
	Temperature  float32    `json:"temperature"`
}

// func (r *Rocket) Connect(eventHelper *helpers.Event) {

// }

type Payload struct {
	Description string  `json:"description"`
	Weight      float32 `json:"weight"`
}

type Telemetry struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

type Timestamps struct {
	Launched  interface{} `json:"launched"`
	Deployed  interface{} `json:"deployed"`
	Failed    interface{} `json:"failed"`
	Cancelled interface{} `json:"cancelled"`
}
