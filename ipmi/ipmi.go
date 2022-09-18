package ipmi

import "os/exec"

var ipmitool = "ipmitool"

func SetExecutablePath(path string) {
	ipmitool = path
}

func Command(args ...string) (string, error) {
	out, err := exec.Command(ipmitool, args...).Output()
	if err != nil {
		return "", err
	}
	return string(out[:]), nil
}
