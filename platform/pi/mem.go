package pi

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/PierreKieffer/pitop/mem"
	"github.com/PierreKieffer/pitop/pkg/utils"
)

func (t *Temp) ExtractMemStats() *mem.MemStat {

	var memStat mem.MemStat

	memStatBytes, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(memStatBytes), "\n")

	for _, statLine := range dataSlice {
		statSlice := utils.FormatStatSlice(strings.Split(statLine, " "))
		if len(statSlice) > 2 {
			if statSlice[0] != "" {
				switch statSlice[0][:len(statSlice[0])-1] {
				case "MemTotal":
					memStat.MemTotal, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "MemFree":
					memStat.MemFree, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "Buffers":
					memStat.Buffers, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "Cached":
					memStat.Cached, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "SwapTotal":
					memStat.SwapTotal, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "SwapFree":
					memStat.SwapFree, _ = strconv.ParseUint(statSlice[1], 10, 64)
				}
			}
		}
	}

	memStat.MemUsage = (float32(memStat.MemTotal) - float32(memStat.MemFree)) / float32(memStat.MemTotal) * 100
	memStat.SwapUsage = (float32(memStat.SwapTotal) - float32(memStat.SwapFree)) / float32(memStat.SwapTotal) * 100

	return &memStat
}
