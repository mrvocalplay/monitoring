package monitoring

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func RunCmd(cmdd string) string {
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

func ToInt(s string) float64 {
	str := s

	// String to Int conversation => strconv.Atoi function return integer with error if any.
	// Atoi means "A" String to "i" integer
	intI, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	}
	return intI
}
