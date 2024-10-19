package main

import (
	"fmt"
	"os"
	"os/exec"
)

func execCommand(command string, arg ...string) error {

	cmd := exec.Command(command, arg...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error running command: %s", err)
	}

	return nil
}

func main() {
	fmt.Println("Hello yet another Golang program! - MAC Changer")

	var interfaceName string
	var new_mac_address string

	fmt.Println("Enter the interface name:")
	fmt.Scanln(&interfaceName)

	fmt.Println("Enter the new MAC address:")
	fmt.Scanln(&new_mac_address)

	execCommand("sudo", "ifconfig", interfaceName, "down")
	execCommand("sudo", "ifconfig", interfaceName, "hw", "ether", new_mac_address)
	execCommand("sudo", "ifconfig", interfaceName, "up")
	execCommand("ifconfig", interfaceName)
}
