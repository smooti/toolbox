package main

import (
	"fmt"
	"log"

	"github.com/elastic/go-sysinfo"
)

func main() {
	osDetails, err := GetOSInfo()
	if err != nil {
		log.Fatal(err)
	}

	hostname := osDetails.Hostname
	osName := osDetails.Name
	osFamily := osDetails.Family
	osVersion := osDetails.Version
	osBuild := osDetails.Build
	architecture := osDetails.Architecture

	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("OS Family: %s\n", osFamily)
	fmt.Printf("OS Name: %s\n", osName)
	fmt.Printf("OS Version: %s\n", osVersion)
	fmt.Printf("OS Build: %s\n", osBuild)
	fmt.Printf("Architecture: %s\n", architecture)

}

type OSInfo struct {
	Hostname     string
	Name         string
	Family       string
	Version      string
	Build        string
	Architecture string
}

func GetOSInfo() (OSInfo, error) {
	// Return error if unable to get host information
	host, err := sysinfo.Host()
	if err != nil {
		log.Fatal(err)
	}

	info := host.Info()
	osInfo := OSInfo{
		Hostname:     info.Hostname,
		Name:         info.OS.Name,
		Family:       info.OS.Family,
		Version:      info.OS.Version,
		Build:        info.OS.Build,
		Architecture: info.Architecture,
	}
	return osInfo, nil

}
