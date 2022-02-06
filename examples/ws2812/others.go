//go:build !digispark && !arduino && !qtpy
// +build !digispark,!arduino,!qtpy

package main

import "machine"

// Replace neo and led in the code below to match the pin
// that you are using if different.
var neo machine.Pin = machine.WS2812
var led = machine.LED
