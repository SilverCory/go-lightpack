package lightpack

import (
	"bufio"
	"errors"
	"net"
	"strings"
	"sync"
)

// ConnectionNotLightpackError is an error to warn that not all if any of this API will work properly with the program.
// Can be ignored and still work but just a precaution.
var ConnectionNotLightpackError = errors.New("connection isn't lightpack and may not work")

func (a *API) readResponse() (string, error) {
	data, _, err := a.reader.ReadLine()
	return string(data), err
}

func (a *API) sendCommand(cmd string) error {
	_, err := a.conn.Write([]byte(cmd + "\n"))
	return err
}

func (a *API) Connect() error {

	// Create a mutex if there isn't one
	if a.lock == nil {
		a.lock = new(sync.Mutex)
	}

	a.lock.Lock()
	defer a.lock.Unlock()

	conn, err := net.Dial("tcp", a.Address)
	if err != nil {
		return err
	}

	// TODO do we need the conn?
	a.conn = conn
	a.reader = bufio.NewReader(a.conn)

	response, err := a.readResponse()
	if err != nil {
		return err
	}

	// Check the first response which should be a version identifier
	if !strings.HasPrefix(response, "Lightpack API") {
		a.version = "UNKNOWN"
		return ConnectionNotLightpackError
	}

	responseParts := strings.Split(response, " ")
	if len(responseParts) >= 3 {
		a.version = responseParts[2]
	}

	return err
}

// GetVersion returns the version of the server's LightPack API.
// If it's not lightpack it should be "UNKNOWN".
func (a *API) GetVersion() string {
	return a.version
}
