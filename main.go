package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/elastic/go-sysinfo"
	"github.com/shirou/gopsutil/disk"
)

func main() {
	GetOSInfo()
	GetDiskInfo()

}

// Retrieves operating system information.
func GetOSInfo() error {
	// Return error if unable to get host information
	host, err := sysinfo.Host()
	if err != nil {
		log.Fatal(err)
	}

	info := host.Info()
	hostname := info.Hostname
	osName := info.OS.Name
	osFamily := info.OS.Family
	osVersion := info.OS.Version
	osBuild := info.OS.Build
	architecture := info.Architecture

	fmt.Println("OS Information:")
	fmt.Printf("  Hostname: %s\n", hostname)
	fmt.Printf("  OS Family: %s\n", osFamily)
	fmt.Printf("  OS Name: %s\n", osName)
	fmt.Printf("  OS Version: %s\n", osVersion)
	fmt.Printf("  OS Build: %s\n", osBuild)
	fmt.Printf("  Architecture: %s\n", architecture)
	return nil

}

// Retrieves hard disk information
func GetDiskInfo(drive ...string) error {

	targetDrive := "C:" // Default for windows
	if runtime.GOOS != "windows" {
		targetDrive = "/" // Default for other systems (Linux, macOS, etc.)
	}

	// We just want to perform operations on a single drive
	if len(drive) > 0 {
		targetDrive = drive[0]
	}

	usage, err := disk.Usage(targetDrive)
	if err != nil {
		log.Fatal(err)
	}

	totalDiskUsage := usage.Total / 1024 / 1024 / 1024
	freeDiskSpace := usage.Free / 1024 / 1024 / 1024
	usedDiskSpace := usage.Used / 1024 / 1024 / 1024
	usedDiskSpacePercentage := usage.UsedPercent
	filesystemType := usage.Fstype
	path := usage.Path

	fmt.Println("Disk Information:")
	fmt.Printf("  Total: %v GB\n", totalDiskUsage)
	fmt.Printf("  Free: %v GB\n", freeDiskSpace)
	fmt.Printf("  Used: %v GB (%.2f%%)\n", usedDiskSpace, usedDiskSpacePercentage)
	fmt.Printf("  Filesystem: %s\n", filesystemType)
	fmt.Printf("  Mount Point: %s\n", path)
	return nil
}
