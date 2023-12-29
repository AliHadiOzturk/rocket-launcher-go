package utils

import (
	"fmt"
	"launcher/models"
	"log"
	"net"
)

type TCP struct {
	host    string
	port    int64
	decoder *Decoder
}

func NewTCPUtil(host string, port int64) *TCP {
	return &TCP{
		host:    host,
		port:    port,
		decoder: NewDecoderUtil(),
	}
}

func (t TCP) Connect(dataReceived func(rocket *models.Rocket)) {
	address := fmt.Sprintf("%s:%d", t.host, t.port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln("Error while connection... :", err)
		return
	}

	log.Printf("Socket connected to %s", address)

	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatalln("Error while reading data:", err)
			break
		}

		rocketData := t.decoder.RocketData(buffer)

		if rocketData == nil {
			continue
		}

		dataReceived(rocketData)
	}

}
