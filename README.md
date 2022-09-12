# ipmi_fan_control
A small program to control Dell PowerEdge server fans via ipmitool using lmsensors temperatures

works on my r710 and r510 (both running proxmox), and it also works on my r210 running pfsense

## Requires
ipmitool and lmsensors

## Build and Install

### Linux

#### build:
```sh
GOOS=linux GOARCH=amd64 go build -o ipmi_fan_control ipmi_fan_control.go
```

#### installation:

##### copy binary:
```sh
cp /path/to/binary /usr/local/bin/ipmi_fan_control
```

##### create a map file
see 'fan_speeds.cfg' for an example

##### do a test run
ipmi_fan_control -s /path/to/fan_speeds.cfg



##### run as a service

check out service/systemd.service 

you may want to rename it to ipmi-fan-control.service on your system.


### PfSense

```sh
GOOS=freebsd GOARCH=amd64 go build -o ipmi_fan_control ipmi_fan_control.go
```

TODO: figure out how to run something as a service on PfSense/BSD



## ipmitool notes
this program assumes that an ipmi device is available at /dev/ipmi0, no arguments regarding ip addresses, usernames, or passwords are passed to ipmitool


