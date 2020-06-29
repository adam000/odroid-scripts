package thermals

import (
	"io/ioutil"
	"strconv"
)

func SetAutomatic() error {
	return ioutil.WriteFile("/sys/devices/platform/pwm-fan/hwmon/hwmon0/automatic", []byte("1"), 0644)
}

func SetManual(fanSpeed int) error {
	err := ioutil.WriteFile("/sys/devices/platform/pwm-fan/hwmon/hwmon0/pwm1", []byte(strconv.Itoa(fanSpeed)), 0644)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("/sys/devices/platform/pwm-fan/hwmon/hwmon0/automatic", []byte("0"), 0644)
}
