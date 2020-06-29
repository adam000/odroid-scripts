package main

import (
	"time"
)

func getNextFanSpeed(cfg Config, temp int, currentFanSpeed int) int {
	desiredFanSpeed := getDesiredFanSpeed(cfg, temp)
	if desiredFanSpeed >= currentFanSpeed {
		return desiredFanSpeed
	}
	if desiredFanSpeed == 0 {
		return 0
	}
	return (desiredFanSpeed + currentFanSpeed) / 2
}

func getDesiredFanSpeed(cfg Config, temp int) int {
	zenithDist := getDistanceFromZenith(cfg.LocalZenith)
	// Test if temperature is below or above interpolation range
	lowTemp := getLowFanTriggerTemp(cfg, zenithDist)
	if temp < lowTemp {
		return 0
	}
	highTemp := getHighTemp(cfg, zenithDist)
	if temp >= highTemp {
		return cfg.FanSpeedLimit.Maximum
	}

	// Otherwise, interpolate in the middle (quadratically for now)
	fanSpeedRange := float64(cfg.FanSpeedLimit.Maximum - cfg.FanSpeedLimit.Minimum)
	tempNormalized := float64(temp-lowTemp) / float64(highTemp-lowTemp)

	fanSpeed := int(float64(cfg.FanSpeedLimit.Minimum) + tempNormalized*tempNormalized*fanSpeedRange)

	return fanSpeed
}

func getLowFanTriggerTemp(cfg Config, zenithDist int) int {
	// Zenith is when it's at its strongest (triangular wave)
	// So subtract the difference between day and night for how
	// far it is from the zenith
	diff := cfg.FanTrigger.Day - cfg.FanTrigger.Night

	return cfg.FanTrigger.Day - (diff * zenithDist / 12)
}

func getHighTemp(cfg Config, zenithDist int) int {
	// Zenith is when it's at its strongest (triangular wave)
	// So subtract the difference between day and night for how
	// far it is from the zenith
	diff := cfg.TempLimit.Day - cfg.TempLimit.Night

	return cfg.TempLimit.Day - (diff * zenithDist / 12)
}

func getDistanceFromZenith(zenith int) int {
	return getHourDistance(zenith, time.Now().Hour())
}

func getHourDistance(zenith int, hour int) int {
	var dist int
	if zenith > hour {
		dist = zenith - hour
	} else {
		dist = hour - zenith
	}
	if dist > 12 {
		dist = 24 - dist
	}
	return dist
}
