package main

import (
	"io/ioutil"
	"log"

	"github.com/IOTechSystems/onvif"
	"github.com/IOTechSystems/onvif/event"
)

func main() {
	dev := onvif.NewDevice(onvif.DeviceParams{
		Xaddr: "192.168.12.148", // BOSCH
		//Xaddr:    "192.168.12.149", // Geovision
		Username: "administrator",
		Password: "Password1!",
	})
	err := dev.GetSupportedServices()
	if err != nil {
		log.Fatalln("fail to get supported services:", err)
	}
	// CreateUsers
	res, err := dev.CallMethod(event.GetEventProperties{})
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, _ := ioutil.ReadAll(res.Body)

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
