package main

import (
	"fmt"
)

const passStatus string = "pass"
const failStatus string = "fail"

type HealthCheck struct {
	serviceID uint8
	status    string
}

func generateCheck() (checks []HealthCheck) {
	var (
		i      uint8
		id     uint8
		status string
	)

	// HealthCheck generation loop
	for i = 0; i < 5; i++ {
		// Generate serviceID
		id = i

		// Generate status
		if i%2 == 0 {
			status = passStatus
		} else {
			status = failStatus
		}

		// Generate HealthCheck and append to slice
		checks = append(checks, HealthCheck{id, status})
	}

	return
}

func main() {
	// Generate HealthChecks
	var checks []HealthCheck = generateCheck()

	// Print HealthChecks ids with passStatus
	for _, check := range checks {
		if check.status == passStatus {
			fmt.Printf("serviceID: %v\tstatus: %v\n", check.serviceID, check.status)
		} else {
			fmt.Println("Check is failed")
		}
	}
}
