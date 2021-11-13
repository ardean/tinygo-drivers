// This is a sensor station that uses a ESP32 running nina-fw over SPI.
// It creates a UDP connection you can use to get info to/from your computer via the microcontroller.
//
// In other words:
// Your computer <--> UART0 <--> MCU <--> SPI <--> ESP32
//
package main

import (
	"machine"
	"strconv"
	"time"

	"tinygo.org/x/drivers/net"
	"tinygo.org/x/drivers/wifinina"
)

// access point info
const ssid = ""
const pass = ""

// IP address of the server aka "hub". Replace with your own info.
const hubIP = ""

var (
	// this is the ESP chip that has the WIFININA firmware flashed on it
	adaptor *wifinina.Device
)

func main() {

	// Init esp8266/esp32
	// Configure SPI for 8Mhz, Mode 0, MSB First
	machine.NINA_SPI.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})

	// these are the default pins for the Arduino Nano33 IoT.
	// change these to connect to a different UART or pins for the ESP8266/ESP32
	adaptor = wifinina.New(machine.NINA_SPI,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN)
	adaptor.Configure()

	// connect to access point
	connectToAP()

	// now make UDP connection
	ip := net.ParseIP(hubIP)
	raddr := &net.UDPAddr{IP: ip, Port: 2222}
	laddr := &net.UDPAddr{Port: 2222}

	println("Dialing UDP connection...")
	conn, _ := net.DialUDP("udp", laddr, raddr)

	for {
		// send data
		println("Sending data...")
		for i := 0; i < 25; i++ {
			conn.Write([]byte("hello " + strconv.Itoa(i) + "\r\n"))
		}
		time.Sleep(1000 * time.Millisecond)
	}

	// Right now this code is never reached. Need a way to trigger it...
	println("Disconnecting UDP...")
	conn.Close()
	println("Done.")
}

// connect to access point
func connectToAP() {
	time.Sleep(2 * time.Second)
	message("Connecting to " + ssid)
	adaptor.SetPassphrase(ssid, pass)
	for st, _ := adaptor.GetConnectionStatus(); st != wifinina.StatusConnected; {
		message("Connection status: " + st.String())
		time.Sleep(1 * time.Second)
		st, _ = adaptor.GetConnectionStatus()
	}
	message("Connected.")
	time.Sleep(2 * time.Second)
	ip, _, _, err := adaptor.GetIP()
	for ; err != nil; ip, _, _, err = adaptor.GetIP() {
		message(err.Error())
		time.Sleep(1 * time.Second)
	}
	message(ip.String())
}

func message(msg string) {
	println(msg, "\r")
}
