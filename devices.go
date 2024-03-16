package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
)

// GetDeviceMemory returns the total and available memory of the device in bytes.
func GetDeviceMemory() (uint64, uint64, error) {
	mem := &runtime.MemStats{}
	runtime.ReadMemStats(mem)
	totalMem := mem.TotalAlloc
	availableMem := mem.HeapSys - mem.HeapAlloc
	return totalMem, availableMem, nil
}

// GetCPUUsage returns the current CPU usage as a percentage.
func GetCPUUsage() (float64, error) {
	var rusage syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &rusage); err != nil {
		return 0, err
	}
	utime := rusage.Utime.Nano()
	stime := rusage.Stime.Nano()
	totalTime := utime + stime
	uptime := float64(rusage.Utime.Sec) + float64(rusage.Stime.Sec)
	cpuUsage := (float64(totalTime) / 1e9) / uptime * 100
	return cpuUsage, nil
}

// GetStorageInformation returns the total and available storage space of the device in bytes.
func GetStorageInformation() (uint64, uint64, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs("/", &stat); err != nil {
		return 0, 0, err
	}
	totalSpace := stat.Blocks * uint64(stat.Bsize)
	availableSpace := stat.Bavail * uint64(stat.Bsize)
	return totalSpace, availableSpace, nil
}

func main() {
	totalMem, availableMem, err := GetDeviceMemory()
	if err != nil {
		fmt.Println("Failed to get device memory:", err)
	} else {
		fmt.Println("Total Memory:", totalMem, "bytes")
		fmt.Println("Available Memory:", availableMem, "bytes")
	}

	cpuUsage, err := GetCPUUsage()
	if err != nil {
		fmt.Println("Failed to get CPU usage:", err)
	} else {
		fmt.Println("CPU Usage:", cpuUsage, "%")
	}

	totalSpace, availableSpace, err := GetStorageInformation()
	if err != nil {
		fmt.Println("Failed to get storage information:", err)
	} else {
		fmt.Println("Total Storage Space:", totalSpace, "bytes")
		fmt.Println("Available Storage Space:", availableSpace, "bytes")
	}
}