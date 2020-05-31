package config

import (
	eventhub "github.com/leandro-lugaresi/hub"
)

var (
	hub *eventhub.Hub
)

func InitEventHub() *eventhub.Hub {
	hub = eventhub.New()
	return hub
}

func GetEventHub() *eventhub.Hub {
	return hub
}
