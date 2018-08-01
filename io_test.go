package lightpack

import (
	"fmt"
	"testing"
)

var a *API

func TestAPI_Connect(t *testing.T) {

	a = &API{
		Address: "127.0.0.1:3636",
	}

	if err := a.Connect(); err != nil {
		t.Error(err)
	}

	fmt.Println("Version: ", a.GetVersion())
	fmt.Println(a.GetCountLEDs())
	fmt.Println(a.GetStatus())
	fmt.Println(a.GetStatusAPI())

}
