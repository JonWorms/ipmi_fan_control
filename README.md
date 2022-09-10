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

##### copy service file
```sh
cp /path/to/service/systemd.service /usr/systemd/system/ipmi-fan-control.service
systemctl daemon-reload
service ipmi-fan-control start
```
### PfSense

```sh
GOOS=freebsd GOARCH=amd64 go build -o ipmi_fan_control ipmi_fan_control.go
```





## ipmitool notes
this program assumes that an ipmi device is available at /dev/ipmi0, no arguments regarding ip addresses, usernames, or passwords are passed to ipmitool


