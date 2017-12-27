package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"time"

	"strings"

	"github.com/godbus/dbus"
)

const DeviceName = `ELAN0732:00 04F3:2567 Pen Pen (0)`

func setCoordinateTransformMatrix(deviceName string, matrix ...string) {
	x := []string{`set-prop`, deviceName, `--type=float`, `Coordinate Transformation Matrix`}
	args := append(x, matrix...)
	cmd := exec.Command("/usr/bin/xinput", args...)

	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	// HACK
	time.Sleep(2 * time.Second)

	log.Printf("Set CTM: %v", matrix)

	err := cmd.Run()

	if err != nil {
		log.Printf("Stderr: %s", stderr.String())
		log.Printf("Stdout: %s", stdout.String())
		log.Printf("Error: %q", err)
	}
}

func isScreenOrientationLocked() bool {
	lock := false
	cmd := exec.Command("/usr/bin/gsettings", "get", "org.gnome.settings-daemon.peripherals.touchscreen", "orientation-lock")

	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()

	if err != nil {
		log.Printf("Stderr: %s", stderr.String())
		log.Printf("Stdout: %s", stdout.String())
		log.Printf("Error: %q", err)
		return false
	}

	result := stdout.String()
	if strings.TrimSpace(result) == "true" {
		log.Printf("Screen orientation locked")
		lock = true
	}

	return lock
}

func main() {

	conn, err := dbus.SystemBus()
	if err != nil {
		log.Printf("Dbus: Failed to connect to system bus: %s", err)
		os.Exit(1)
	}

	conn.BusObject().Call("org.freedesktop.DBus.AddMatch",
		0,
		"type='signal',path='/net/hadess/SensorProxy',interface='org.freedesktop.DBus.Properties',sender='net.hadess.SensorProxy'")

	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)

	log.Println("Waiting for dbus events...")

	for v := range c {
		if !isScreenOrientationLocked() {

			args := v.Body[1].(map[string]dbus.Variant)

			val := args["AccelerometerOrientation"].Value()

			switch val {
			case "normal":
				log.Printf("Normal orientation")
				setCoordinateTransformMatrix(DeviceName, "1.000000", "0.000000", "0.000000", "0.000000", "1.000000", "0.000000", "0.000000", "0.000000", "1.000000")
			case "left-up":
				log.Printf("Left-up orientation")
				setCoordinateTransformMatrix(DeviceName, "0.000000", "-1.000000", "1.000000", "1.000000", "0.000000", "0.000000", "0.000000", "0.000000", "1.000000")
			case "bottom-up":
				log.Printf("Bottom-up orientation")
				setCoordinateTransformMatrix(DeviceName, "-1.000000", "0.000000", "1.000000", "0.000000", "-1.000000", "1.000000", "0.000000", "0.000000", "1.000000")
			case "right-up":
				log.Printf("Right-up orientation")
				setCoordinateTransformMatrix(DeviceName, "0.000000", "1.000000", "0.000000", "-1.000000", "0.000000", "1.000000", "0.000000", "0.000000", "1.000000")
			}
		}
	}
}
