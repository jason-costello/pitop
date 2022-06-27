package pi

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PierreKieffer/pitop/interfaces"
	"github.com/PierreKieffer/pitop/pkg/utils"
)

func (t *Temp) GetNetStats() *interfaces.NetStat {

	var netStats []interfaces.NetStat

	netStatBytes, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(netStatBytes), "\n")

	for _, statLine := range dataSlice[2:] {
		statSlice := utils.FormatStatSlice(strings.Split(statLine, " "))
		extractNetStats(&netStats, statSlice)
	}

	var totalNetStat interfaces.NetStat

	for _, netStat := range netStats {
		totalNetStat.TotalBytesRecv += netStat.TotalBytesRecv
		totalNetStat.TotalBytesSent += netStat.TotalBytesSent

	}

	return &totalNetStat
}

func extractNetStats(netStats *[]interfaces.NetStat, statSlice []string) {
	if len(statSlice) > 1 && statSlice[0] != "" {
		var netStat interfaces.NetStat
		netStat.TotalBytesRecv, _ = strconv.ParseUint(statSlice[1], 10, 64)
		netStat.TotalBytesSent, _ = strconv.ParseUint(statSlice[9], 10, 64)

		*netStats = append(*netStats, netStat)
	}
}

func (t *Temp) ComputeNetStats() *interfaces.NetStat {
	prevNetStats := t.GetNetStats()
	time.Sleep(time.Second)
	netStats := t.GetNetStats()

	if prevNetStats.TotalBytesRecv == 0 && prevNetStats.TotalBytesSent == 0 {
		netStats.BytesRecv = 0
		netStats.BytesSent = 0

		return netStats
	}

	netStats.BytesRecv = netStats.TotalBytesRecv - prevNetStats.TotalBytesRecv
	netStats.BytesSent = netStats.TotalBytesSent - prevNetStats.TotalBytesSent

	return netStats
}
