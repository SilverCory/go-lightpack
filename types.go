package lightpack

import (
	"errors"
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
