package rpi

import (
	"runtime"

	mon "github.com/mrvocalplay/monitoring"
)

// GetOs gives you your OS and Arcitecture
func GetOs() string {
	os := runtime.GOOS
	arch := runtime.GOARCH
	info := os + " " + arch
	return info
}

// GetCPUTemp returns the CPU and GPU Temperature from the RPI
func GetCPUTemp() string {
	return mon.RunCmd("vcgencmd measure_temp")
}

// GetCPUSpannung jjj
func GetCPUVolts() string {
	return mon.RunCmd("vcgencmd measure_volts")
}

func GetArmClock() string {
	return mon.RunCmd("vcgencmd measure_clock arm")
}
