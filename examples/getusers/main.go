package main

import (
	"io"
	"log"

	"github.com/mattiasberlin/onvif"
	"github.com/mattiasberlin/onvif/device"
)

func main() {
	dev := onvif.NewDevice(onvif.DeviceParams{
		Xaddr:    "192.168.12.149",
		Username: "administrator",
		Password: "Password1!",
	})
	// CreateUsers
	res, err := dev.CallMethod(device.GetUsers{})
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, _ := io.ReadAll(res.Body)

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
