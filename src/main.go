package main

import (
	"log"
	"os"
	"time"

	"github.com/tiagordc/go-apt-reboot/src/tplink"
)

func main() {

	plugIp := os.Getenv("TAPO_IP")
	user := os.Getenv("TAPO_USERNAME")
	password := os.Getenv("TAPO_PASSWORD")

	plug := tplink.New(plugIp, user, password)

	if err := plug.Handshake(); err != nil {
		log.Panic(err)
	}

	if err := plug.Login(); err != nil {
		log.Panic(err)
	}

	info, err := plug.GetDeviceInfo()

	if err != nil {
		log.Panic(err)
	}

	if info.Result.DeviceON {
		plug.Switch(false)
		time.Sleep(15 * time.Second)
		plug.Switch(true)
	} else {
		plug.Switch(true)
	}

}
