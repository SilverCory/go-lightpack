package lightpack

import (
	"fmt"
	"testing"
)

func TestParseStatus(t *testing.T) {

	if status, err := ParseStatus("status:on"); err != nil {
		t.Error(err)
	} else if status != StatusOn {
		t.Error("Parsing status \"status:on\" yeilds invalid response")
	}

	if status, err := ParseStatus("status:off"); err != nil {
		t.Error(err)
	} else if status != StatusOff {
		t.Error("Parsing status \"status:off\" yeilds invalid response")
	}

	if status, err := ParseStatus("status:unknown"); err != nil {
		t.Error(err)
	} else if status != StatusUnknown {
		t.Error("Parsing status \"status:unknown\" yeilds invalid response")
	}

	if status, err := ParseStatus("status:device_error"); err != nil {
		t.Error(err)
	} else if status != StatusDeviceError {
		t.Error("Parsing status \"status:device_error\" yeilds invalid response")
	}

	if status, err := ParseStatus("status:invalid_input"); err == nil && status != StatusUnknown {
		t.Error("Parsing status \"status:invalid_input\" yeilds invalid response")
	}

}

func TestParseStatusAPI(t *testing.T) {

	if status, err := ParseStatusAPI("statusapi:idle"); err != nil {
		t.Error(err)
	} else if status != StatusAPIIdle {
		t.Error("Parsing statusapi \"statusapi:idle\" yeilds invalid response")
	}

	if status, err := ParseStatusAPI("statusapi:busy"); err != nil {
		t.Error(err)
	} else if status != StatusAPIBusy {
		t.Error("Parsing statusapi \"statusapi:busy\" yeilds invalid response")
	}

	if _, err := ParseStatusAPI("statusapi:invalid_input"); err == nil {
		t.Error("Parsing statusapi \"statusapi:invalid_input\" yeilds invalid response")
	}

}

func TestParseColors(t *testing.T) {
	input := `colors:0-97,102,103;1-102,104,105;2-102,104,105;3-101,103,102;4-97,100,103;5-100,100,100;6-102,103,103;7-98,101,103;8-94,96,98;9-92,94,95;10-94,96,98;11-93,95,96;12-93,94,95;13-92,93,92;14-90,93,93;15-90,93,94;16-90,93,93;17-90,93,93;18-89,92,90;19-90,92,92;20-93,94,95;21-93,94,94;22-93,93,93;23-90,93,93;24-90,93,93;25-89,92,92;26-94,96,97;27-96,97,96;28-90,93,94;29-94,96,97;`
	if colors, err := ParseColors(input); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%v\n", colors)
	}
}