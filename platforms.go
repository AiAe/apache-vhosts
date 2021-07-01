package main

import "runtime"

type Platform struct {
	Windows bool
	Linux   bool
	Darwin  bool
}

func checkPlatform() *Platform {
	p := Platform{}
	os := runtime.GOOS
	switch os {
	case "windows":
		p.Windows = true
	case "darwin":
		p.Darwin = true
	case "linux":
		p.Linux = true
	}

	return &p
}
