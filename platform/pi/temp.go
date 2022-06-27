package pi

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Temp struct{}

func (t *Temp) ExtractTemp() *float64 {
	tempBytes, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tempInfo, _ := strconv.ParseUint(strings.Split(string(tempBytes), "\n")[0], 10, 64)

	rt := float64(tempInfo) / 1000

	return &rt
}
