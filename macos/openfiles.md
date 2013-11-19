sudo launchctl limit maxfiles 1000000 unlimited
To make this permanent (i.e not reset when you reboot), create /etc/launchd.conf containing:

limit maxfiles 1000000 unlimited
Then you can use ulimit (but without the sudo) to adjust your process limit.
