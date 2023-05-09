package main

import (
	"flag"
	"fmt"
)

func main() {
	metricsFlag := flag.Bool("metrics", false, "Execute metrics function")
	alertsFlag := flag.Bool("alerts", false, "Execute alerts function")

	flag.Parse()

	if *metricsFlag && *alertsFlag {
		fmt.Println("Error: Only one of --metrics or --alerts should be specified.")
		return
	}

	if *metricsFlag {
		printMetrics()
	} else if *alertsFlag {
		printAlerts()
	} else {
		fmt.Println("Error: Please provide either --metrics or --alerts flag.")
	}
}
