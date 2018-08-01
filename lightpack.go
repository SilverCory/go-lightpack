package lightpack

import (
	"bufio"
	"errors"
	"net"
	"strconv"
	"strings"
	"sync"
)

type API struct {
	// Config parts
	Address string `json:"address"`
	LEDMap  []int  `json:"led_map"`
	APIKey  string `json:"api_key"`

	// IO parts
	conn   net.Conn
	reader *bufio.Reader
	lock   *sync.Mutex

	// Information
	version string
}

func (a *API) GetStatus() (Status, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if err := a.sendCommand("getstatus"); err != nil {
		return StatusUnknown, err
	}

	response, err := a.readResponse()
	if err != nil {
		return StatusUnknown, err
	}

	return ParseStatus(response)
}

func (a *API) GetStatusAPI() (StatusAPI, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if err := a.sendCommand("getstatusapi"); err != nil {
		return StatusAPIBusy, err
	}

	response, err := a.readResponse()
	if err != nil {
		return StatusAPIBusy, err
	}

	return ParseStatusAPI(response)
}

func (a *API) GetCountLEDs() (int, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if err := a.sendCommand("getcountleds"); err != nil {
		return -1, err
	}

	response, err := a.readResponse()
	if err != nil {
		return -1, err
	} else if !strings.HasPrefix(response, "countleds:") {
		return -1, errors.New("invalid response for getcountleds")
	}

	return strconv.Atoi(response[10:])

}
