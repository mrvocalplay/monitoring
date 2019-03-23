package rpi

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
	return runCmd("vcgencmd measure_temp")
}

// GetCPUSpannung jjj
func GetCPUSpannung() string {
	return CutStr(runCmd("vcgencmd measure_volts"), "volt=", "V")
}

func GetArmClock() string {
	return CutStr(runCmd("vcgencmd measure_clock arm"), "frequency(45)=", "")
}

func runCmd(cmdd string) string {
	cmd := exec.Command("bash", "-c", cmdd)
	output := new(bytes.Buffer)
	cmd.Stdout = output
	cmd.Stderr = output
	if err := cmd.Start(); err != nil {
		// after Start program is continued and script is executing in background
		fmt.Printf("Failed to start " + err.Error())
		os.Exit(1)
	}
	cmd.Wait()
	return output.String()
}

func CutStr(str string, cut1 string, cut2 string) string {
	str = strings.Replace(str, cut1, "", -1)
	str = strings.Replace(str, cut2, "", -1)
	str = strings.TrimRight(str, "\r\n")
	return str
}
