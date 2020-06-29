package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adam000/odroid-scripts/tempctrl/thermals"
)

func main() {
	if os.Geteuid() != 0 {
		log.Println("This program can only be run as root.")
		return
	}

	configFilename := "/root/tempctrl/config.json"
	cfg, err := LoadConfig(configFilename)
	if err != nil {
		log.Println("Could not load config file:", configFilename)
		return
	}

	debug := flag.Bool("debug", false, "debug output")
	flag.Parse()

	defer func() {
		if err := thermals.SetAutomatic(); err != nil {
			panic(err)
		}
	}()

	thermalZones := []string{
		"/sys/devices/virtual/thermal/thermal_zone0/temp",
		"/sys/devices/virtual/thermal/thermal_zone1/temp",
		"/sys/devices/virtual/thermal/thermal_zone2/temp",
		"/sys/devices/virtual/thermal/thermal_zone3/temp",
	}

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	tickerInterval := time.Duration(cfg.TickerIntervalMs) * time.Millisecond
	ticker := time.NewTicker(tickerInterval)

	currentFanSpeed := 0

	inIdlePeriod := false
	var idlePeriodStart time.Time
	idlePeriod := time.Duration(cfg.IdlePeriod) * time.Second

	// In a loop, read values from each of the thermal zones
	for {
		select {
		case <-sigterm:
			ticker.Stop()
			if debug != nil && *debug == true {
				log.Println("SIGTERM or SIGINT received. Shutting down")
			}
			return
		case <-ticker.C:
			thermals.SetManual(currentFanSpeed)

			temps := thermals.Read(thermalZones)
			maxTemp := 0
			for _, temp := range temps {
				if maxTemp < temp {
					maxTemp = temp
				}
			}
			nextFanSpeed := getNextFanSpeed(cfg, maxTemp/1000, currentFanSpeed)
			if debug != nil && *debug == true {
				log.Printf("MaxTmp: %d\tCurSp: %d\tNextSp: %d", maxTemp, currentFanSpeed, nextFanSpeed)
			}

			if inIdlePeriod {
				if nextFanSpeed >= cfg.FanSpeedLimit.Minimum {
					// Exit idle period
					if debug != nil && *debug == true {
						log.Println("Exiting idle period because fan ramped up")
					}
					inIdlePeriod = false
				} else if time.Now().Sub(idlePeriodStart) < idlePeriod {
					// Stay in idle period
					if debug != nil && *debug == true {
						log.Println("Staying in idle period...")
					}
					continue
				} else {
					// Expire idle period
					if debug != nil && *debug == true {
						log.Println("Idle period complete. Fans should die")
					}
					inIdlePeriod = false
				}
			} else if currentFanSpeed != 0 && nextFanSpeed == 0 && currentFanSpeed >= cfg.FanSpeedLimit.Minimum {
				// Start idle period
				if debug != nil && *debug == true {
					log.Println("Starting idle period")
				}
				inIdlePeriod = true
				idlePeriodStart = time.Now()
				currentFanSpeed = cfg.FanSpeedLimit.Minimum
				continue
			}

			currentFanSpeed = nextFanSpeed
		}
	}
}
