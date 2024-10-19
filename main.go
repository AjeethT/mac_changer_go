package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello yet another Golang program! - MAC Changer")

	var interfaceName string
	var new_mac_address string

	fmt.Println("Enter the interface name:")
	fmt.Scanln(&interfaceName)

	fmt.Println("Enter the new MAC address:")
	fmt.Scanln(&new_mac_address)

	cmd := exec.Command("sudo", "ifconfig", interfaceName, "down")
	o, err := cmd.Output()

	if err != nil {
		fmt.Println("Error running ifconfig: ", err)
	}

	fmt.Println(string(o))

	cmd = exec.Command("sudo", "ifconfig", interfaceName, "hw", "ether", new_mac_address)
	o, err = cmd.Output()

	if err != nil {
		fmt.Println("Error running ifconfig: ", err)
	}

	fmt.Println(string(o))

	cmd = exec.Command("sudo", "ifconfig", interfaceName, "up")

	o, err = cmd.Output()
	if err != nil {
		fmt.Println("Error running ifconfig: ", err)
	} else {
		fmt.Println("Changed the MAC address successfully for interface", interfaceName, " as ", new_mac_address)
	}

	fmt.Println(string(o))

	cmd = exec.Command("sudo", "ifconfig", interfaceName)
	o, err = cmd.Output()

	if err != nil {
		fmt.Println("Error running ifconfig: ", err)
	}

	fmt.Println(string(o))
}
