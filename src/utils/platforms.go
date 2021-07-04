package utils

import "runtime"

var Plt Platform

type Platform struct {
	Windows bool
	Linux   bool
	Darwin  bool
}

func CheckPlatform(p *Platform) {
	os := runtime.GOOS
	switch os {
	case "windows":
		p.Windows = true
	case "darwin":
		p.Darwin = true
	case "linux":
		p.Linux = true
	}
}
