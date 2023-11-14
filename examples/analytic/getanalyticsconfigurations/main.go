package main

import (
	"io"
	"log"

	"github.com/IOTechSystems/onvif"
	"github.com/IOTechSystems/onvif/media2"
)

func main() {
	dev := onvif.NewDevice(onvif.DeviceParams{
		Xaddr:    "192.168.12.148", // BOSCH
		Username: "administrator",
		Password: "Password1!",
		AuthMode: onvif.UsernameTokenAuth,
	})
	err := dev.GetSupportedServices()
	if err != nil {
		log.Fatalln("failed to get supported services:", err)
	}

	res, err := dev.CallMethod(media2.GetAnalyticsConfigurations{})
	if err != nil {
		log.Fatalln("failed to call method:", err)
	}
	bs, _ := io.ReadAll(res.Body)

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
