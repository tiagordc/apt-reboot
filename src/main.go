package main

import (
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-ping/ping"
	"github.com/tiagordc/go-apt-reboot/src/tplink"
)

func restart() {

	address := os.Getenv("TAPO_IP")
	user := os.Getenv("TAPO_USERNAME")
	password := os.Getenv("TAPO_PASSWORD")
	plug := tplink.New(address, user, password)

	if err := plug.Handshake(); err != nil {
		log.Println("Handshake error:", err)
		return
	}

	if err := plug.Login(); err != nil {
		log.Println("Login error:", err)
		return
	}

	info, err := plug.GetDeviceInfo()

	if err != nil {
		log.Println("Could not get plug data:", err)
		return
	}

	if info.Result.DeviceON {
		log.Println("Restarting")
		if err := plug.Switch(false); err != nil {
			log.Println("Could not switch off:", err)
			return
		}
		time.Sleep(5 * time.Second)
	} else {
		log.Println("Starting")
	}

	if err := plug.Switch(true); err != nil {
		log.Println("Could not switch on:", err)
	}

}

func isOnline(address string) bool {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		log.Println("Ping error:", err)
		return false
	}
	pinger.Count = 1
	pinger.Timeout = 1 * time.Second
	pinger.Run()
	stats := pinger.Statistics()
	return stats.PacketsSent == stats.PacketsRecv
}

func main() {
	log.Println("Starting application")
	s := gocron.NewScheduler(time.UTC)
	atv := os.Getenv("ATV_ADDRESS")
	cron, present := os.LookupEnv("CRON")
	if present {
		s.Cron(cron).Do(restart)
	}
	s.Every(10).Seconds().Do(func() {
		if !isOnline(atv) {
			restart()
		}
	})
	s.StartAsync()
	s.StartBlocking()
}
