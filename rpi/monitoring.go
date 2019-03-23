package rpi

import (
	"runtime"

	"github.com/mrvocalplay/monitoring"
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
	return monitoring.RunCmd("vcgencmd measure_temp")
}

// GetCPUSpannung jjj
func GetCPUSpannung() string {
	return monitoring.CutStr(runCmd("vcgencmd measure_volts"), "volt=", "V")
}

func GetArmClock() string {
	return monitoring.CutStr(runCmd("vcgencmd measure_clock arm"), "frequency(45)=", "")
}
