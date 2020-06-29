This is just a collection of scripts I use on my odroid xu4. Maybe this will
help someone else out :)

### ttop - temperature "top"

Prints the values from the 4 "thermal zones". Defaults to 3 second sleep, but an
argument can be provided as the number of seconds between printing the values.
Previous values are overwritten. Ctrl-c to stop it.

### tempctrl - temperature control daemon

Meant to be run as a service and control the fan speed.

Values can be set in a json configuration file.
