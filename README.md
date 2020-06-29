This is just a collection of scripts I use on my odroid xu4. Maybe this will
help someone else out :)

### ttop - temperature "top"

Prints the values from the 4 "thermal zones". Defaults to 3 second sleep, but an
argument can be provided as the number of seconds between printing the values.
Previous values are overwritten. Ctrl-c to stop it.

### tempctrl - temperature control daemon

Meant to be run as a service and control the fan speed. It is designed to change its thresholds between "night" and "day" to stay quiet at night but more aggressively cool things off during the day.

You can run it from the command line with `-debug` to see an audit of the decisions it makes.

Values can be set in the json configuration file, which should be placed at `/root/tempctrl/config.json`. None of those values are validated on startup so make sure you have them set correctly. Description of the values:

* TickerIntervalMs - the time in milliseconds between successive readings of temperature
* IdlePeriod - the time in seconds when the fan is below its minimum temperature before the fan speed drops from its minimum speed to 0. This keeps the fan from turning on and off frequently at low load
* LocalZenith - the hour of the day that is the "warmest" in general. This is considered the middle of the "day."
* FanTrigger - The minimum temperature that triggers the fan to run at all, in Celcius
* FanSpeedLimit - the upper and lower bounds of the fan speed. The numbers must be in the range 0-255.
* TempLimit - the upper end of the temperature limit before the fan should be set to its maximum speed. The odroid xu4 has a thermal limit around 85 or 90 Celcius.
