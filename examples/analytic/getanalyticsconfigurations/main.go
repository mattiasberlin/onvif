package main

import (
	"io"
	"log"

	"github.com/mattiasberlin/onvif"
	"github.com/mattiasberlin/onvif/media2"
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
		log.Fatalln("fail to get supported services:", err)
	}

	res, err := dev.CallMethod(media2.GetAnalyticsConfigurations{})
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, _ := io.ReadAll(res.Body)

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
