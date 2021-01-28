package main

import (
	"fmt"
	"time"

	jitter "github.com/Telefonica/go-jitter"
	"github.com/Telefonica/ping"
)

func main() {
	ping, err := ping.NewPinger("google.com")
	j, err := jitter.NewJitterer("google.com", ping)
	if err != nil {
		fmt.Println(err)
	}

	j.SetBlockSampleSize(10)
	j.SetPingerPrivileged(true)
	j.SetPingerTimeout(time.Second * 10)

	j.Run()

	s := j.Statistics()

	fmt.Println("Squared Deviation: ", s.SquaredDeviation)
	fmt.Println("Uncorrected Deviation: ", s.UncorrectedSD)
	fmt.Println("Corrected Deviation: ", s.CorrectedSD)
	fmt.Println("RTT Range: ", s.RttRange)
	fmt.Println("RTTs: ", s.RTTS)
}
