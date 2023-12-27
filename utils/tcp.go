package utils

import (
	"fmt"
	"launcher/helpers"
	"log"
	"net"
)

type TCP struct {
	host         string
	port         int64
	eventManager *helpers.Event
	decoder      *Decoder
}

func NewTCPUtil(host string, port int64, eventManager *helpers.Event) *TCP {
	return &TCP{
		host:         host,
		port:         port,
		eventManager: eventManager,
		decoder:      NewDecoderUtil(),
	}
}

func (t TCP) Connect() {
	address := fmt.Sprintf("%s:%d", t.host, t.port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln("Error while connection... :", err)
		return
	}

	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatalln("Error while reading data:", err)
			return
		}

		rocketID := t.decoder.RocketData(buffer)

		// Process and use the data (here, we'll just print it)
		t.eventManager.Trigger(helpers.DataReceived, map[string]any{"name": rocketID})
	}

}
