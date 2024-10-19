package main

import (
	"flag"
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

	var interfaceName, new_mac_address string
	flag.StringVar(&interfaceName, "interface", "", "Interface name to change MAC address  (e.g., eth0)  (required)")
	flag.StringVar(&interfaceName, "i", "", "Short form of Interface name to change MAC address")

	flag.StringVar(&new_mac_address, "mac", "", "New MAC address (e.g., 00:11:22:33:44:55)  (required)")
	flag.StringVar(&new_mac_address, "m", "", "Short form of New MAC address")

	flag.Parse()

	if interfaceName == "" || new_mac_address == "" {
		fmt.Println("Error: Both interface and MAC address are required.")
		os.Exit(1)
	}

	execCommand("sudo", "ifconfig", interfaceName, "down")
	execCommand("sudo", "ifconfig", interfaceName, "hw", "ether", new_mac_address)
	execCommand("sudo", "ifconfig", interfaceName, "up")
	execCommand("ifconfig", interfaceName)
}
