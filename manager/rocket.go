package manager

import (
	"encoding/json"
	"launcher/models"
	"launcher/utils"
	"log"
)

type RocketManager struct {
	Rockets []models.Rocket
}

func NewRocketManager() *RocketManager {
	return &RocketManager{}
}

func (r *RocketManager) Init() *RocketManager {

	response := utils.NewHttpUtil().Get("http://localhost:5000/rockets", map[string]string{"X-API-KEY": "API_KEY_1"})

	if response == nil {
		log.Println("Something happened while getting rockets")
		panic("Something happened while getting rockets")
	}

	var rockets []models.Rocket
	if err := json.Unmarshal([]byte(response), &rockets); err != nil {
		panic(err)
	}

	log.Println("Rocket count", len(rockets))

	r.Rockets = rockets

	return r

}

// func (r *RocketManager) ConnectRockets(eventHelper *helpers.Event) {
// 	for _, rocket := range r.rockets {
// 		tcpUtil := utils.NewTCPUtil(rocket.Telemetry.Host, rocket.Telemetry.Port, eventHelper)
// 		go tcpUtil.Connect()
// 	}
// }
