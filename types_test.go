package lightpack

import "testing"

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
