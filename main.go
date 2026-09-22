package main

import (
	"fmt"

	"github.com/caseymrm/go-smc"
)

func main() {
	smc.OpenSMC()
defer smc.CloseSMC()

	temps := smc.ReadTemperatures()

	for key, temp := range temps {
		fmt.Printf("%s: %.2f°C\n", key, temp)
	}

	fmt.Printf("Hottest: %.2f°C\n", smc.ReadTemperature())
}
