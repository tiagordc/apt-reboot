package main

import (
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/tiagordc/go-apt-reboot/src/tplink"
)

func run() {

	address := os.Getenv("TAPO_IP")
	user := os.Getenv("TAPO_USERNAME")
	password := os.Getenv("TAPO_PASSWORD")
	plug := tplink.New(address, user, password)

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
		log.Println("Restarting")
		plug.Switch(false)
		time.Sleep(15 * time.Second)
		plug.Switch(true)
	} else {
		log.Println("Turning on")
		plug.Switch(true)
	}

}

func main() {
	cron := os.Getenv("CRON")
	s := gocron.NewScheduler(time.UTC)
	s.Cron(cron).Do(run)
	s.StartAsync()
	s.StartBlocking()
}
