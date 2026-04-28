package main

import (
	"MQTT/internal/config"
	"MQTT/internal/mqtt"
	"MQTT/pkg/logging"

	"log"
)

var params *config.Config

func init() {
	pr, err := config.LoadEnvFile("./config.env")
	if err != nil {
		log.Fatalln(err)
	}

	if err := pr.ValidateConfig(); err != nil {
		log.Fatalln(err)
	}

	params = pr
}

func main() {
	// Запрашиваем готовые топики с покозаниями
	clientSensor, err := mqtt.RunApp(params)
	if err != nil {
		logging.LogToFile(err, `"ERROR" MQTT: `)
		log.Fatalln(err)
	}

	logging.LogToFile(clientSensor, `"OK" MQTT: `)

	// Отправляем данные на сервер
	status, err := mqtt.SendJsonPOST(clientSensor)
	if err != nil {
		logging.LogToFile(status+" "+err.Error(), `"ERROR" SERVER: `)
		log.Fatalln(err)
	}

	logging.LogToFile(status, `"OK" SERVER: `)
}
