package main

import (
	"MQTT/internal/config"
	"MQTT/internal/mqtt"
	"MQTT/pkg/logging"
	"fmt"
	"os"

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
		logging.LogToFile(err, "ERROR MQTT")
		os.Exit(1)
	}

	logging.LogToFile(clientSensor, "OK MQTT")

	// Отправляем данные на сервер
	status, err := mqtt.SendJsonPOST(clientSensor)
	if err != nil {
		logging.LogToFile(fmt.Sprintf("%v - %v", status, err), "ERROR SERVER")
		os.Exit(1)
	}
	
	logging.LogToFile(status, "OK SERVER")
}
