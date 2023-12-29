package utils

import (
	"encoding/binary"
	"launcher/models"
	"math"

	"github.com/sigurn/crc16"
)

type Decoder struct{}

func NewDecoderUtil() *Decoder {
	return &Decoder{}
}

func (d *Decoder) RocketData(data []byte) *models.Rocket {
	rocketID := string(data[1:11])

	altitude := math.Float32frombits(binary.BigEndian.Uint32(data[13:17]))
	speed := math.Float32frombits(binary.BigEndian.Uint32(data[17:21]))
	acceleration := math.Float32frombits(binary.BigEndian.Uint32(data[21:25]))
	thrust := math.Float32frombits(binary.BigEndian.Uint32(data[25:29]))
	temperature := math.Float32frombits(binary.BigEndian.Uint32(data[29:33]))

	// TODO: Implement crc check

	crc := binary.BigEndian.Uint16(data[33:35])

	table := crc16.MakeTable(crc16.CRC16_BUYPASS)

	checksum := crc16.Checksum(data[:33], table)

	if crc != checksum {
		return nil
	}

	rocket := models.Rocket{
		ID:           rocketID,
		Altitude:     altitude,
		Speed:        speed,
		Acceleration: acceleration,
		Thrust:       thrust,
		Temperature:  temperature,
	}

	return &rocket
}
