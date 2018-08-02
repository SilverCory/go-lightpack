package lightpack

import (
	"errors"
		"strconv"
	"strings"
)

type Status string

const (
	StatusOn          Status = "on"
	StatusOff         Status = "off"
	StatusUnknown     Status = "unknown"
	StatusDeviceError Status = "device_error"
)

func ParseStatus(response string) (Status, error) {
	if strings.HasPrefix(response, "status:") {
		response = response[7:]
	}

	ret := Status(response)
	if ret != StatusOn && ret != StatusOff && ret != StatusDeviceError && ret != StatusUnknown {
		return StatusUnknown, errors.New("invalid status")
	}

	return ret, nil
}

type StatusAPI string

const (
	StatusAPIBusy StatusAPI = "busy"
	StatusAPIIdle StatusAPI = "idle"
)

func ParseStatusAPI(response string) (StatusAPI, error) {
	if strings.HasPrefix(response, "statusapi:") {
		response = response[10:]
	}

	ret := StatusAPI(response)
	if ret != StatusAPIIdle && ret != StatusAPIBusy {
		return StatusAPIBusy, errors.New("invalid status")
	}

	return ret, nil
}

// I really hate spelling it wrong.
type Color struct {
	R, G, B int16
}

func ParseColors(response string) ([]Color, error) {
	if strings.HasPrefix(response, "colors:") {
		response = response[7:]
	}

	ledColors := strings.Split(response, ";")
	ret := make([]Color, len(ledColors))

	for _, v := range ledColors {
		// Fixes trailing ; from making this fail.
		if v == "" {
			continue
		}

		led, color, err := ParseColor(v)
		if err != nil {
			return ret, err
		}

		ret[led] = color
	}

	return ret, nil
}

func ParseColor(color string) (int, Color, error) {
	colorParts := strings.SplitN(color, "-", 2)
	ledNumber, err := strconv.Atoi(colorParts[0])
	if err != nil {
		return ledNumber, Color{}, err
	}

	RGBVals := strings.Split(colorParts[1], ",")
	if len(RGBVals) != 3 {
		return ledNumber, Color{}, errors.New("invalid RGB values provided")
	}

	retColor := Color{}

	if r, err := strconv.ParseInt(RGBVals[0], 10, 16); err != nil {
		return ledNumber, Color{}, err
	} else {
		retColor.R = int16(r)
	}

	if g, err := strconv.ParseInt(RGBVals[1], 10, 16); err != nil {
		return ledNumber, Color{}, err
	} else {
		retColor.G = int16(g)
	}

	if b, err := strconv.ParseInt(RGBVals[2], 10, 16); err != nil {
		return ledNumber, Color{}, err
	} else {
		retColor.B = int16(b)
	}

	return ledNumber, retColor, nil
}