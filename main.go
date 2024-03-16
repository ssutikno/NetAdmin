package main

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/user/NetAdmin/device"
)

func main() {
	app := buffalo.New(buffalo.Options{
		Env:         envy.Get("GO_ENV", "development"),
		SessionName: "_yourappname_session",
	})

	app.GET("/", HomeHandler)

	admin := app.Group("/admin")
	admin.GET("/dashboard", DashboardHandler)

	app.Serve()
}

func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

func DashboardHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("dashboard.html"))
}
func DeviceAdminHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("device_admin.html"))
}

func DeviceStatusHandler(c buffalo.Context) error {
	// Retrieve device statuses (memory, storage spaces, CPU usage, etc.)
	memoryStatus := getMemoryStatus()
	storageStatus := getStorageStatus()
	cpuUsage := getCPUUsage()

	// Render the device_status.html template with the device statuses
	return c.Render(200, r.HTML("device_status.html", buffalo.Data{
		"MemoryStatus":  memoryStatus,
		"StorageStatus": storageStatus,
		"CPUUsage":      cpuUsage,
	}))
}

func DeviceAdminHandler(c buffalo.Context) error {
	// Retrieve all devices
	devices := device.GetAllDevices()

	// Render the device_admin.html template with the devices
	return c.Render(200, r.HTML("device_admin.html", buffalo.Data{
		"Devices": devices,
	}))
}
