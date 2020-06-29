package main

import (
	"encoding/json"
	"io/ioutil"
)

type Range struct {
	Minimum int
	Maximum int
}

type NightDayVariance struct {
	Night int
	Day   int
}

type Config struct {
	TickerIntervalMs int
	IdlePeriod       int
	LocalZenith      int
	FanTrigger       NightDayVariance
	FanSpeedLimit    Range
	TempLimit        NightDayVariance
}

func LoadConfig(path string) (Config, error) {
	// I think this is specific to Isolate's OS package
	//path = os.GetRelativeDir(path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config

	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
