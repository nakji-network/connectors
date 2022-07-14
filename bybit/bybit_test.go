package bybit

import (
	"testing"
	// "time"
)

var mockBybit BybitConnector

func TestSetDefaults(t *testing.T) {

	// Test for resetting default values
	mockBybit.Name = "WrongName"
	mockBybit.Enabled = false
	mockBybit.Verbose = false

	err := mockBybit.SetDefaults()

	if err != nil {
		t.Error("failed to set defaults for ByBit. error: ", err)
	}

	if mockBybit.Name != "Bybit" ||
		mockBybit.Enabled != true ||
		mockBybit.Verbose != true {
		t.Error("failed to set default values for ByBit")
	}
}

// Commented this test out as it is bad practice to make external HTTP calls

// func TestSendHTTPRequest(t *testing.T) {
// 	var result = struct {
// 		ResultWrapper
// 		Result []Instrument `json:"result"`
// 	}{}

// 	tests := []struct {
// 		name        string
// 		APIEndpoint string
// 		expectedErr error
// 	}{
// 		{"valid endpoint",
// 			"/v2/public/tickers",
// 			nil,
// 		},
// 		{"invalid endpoint",
// 			"aweiawjfasf",
// 			fmt.Errorf("Get \"https://api.bybit.comaweiawjfasf\": dial tcp: lookup api.bybit.comaweiawjfasf: no such host"),
// 		},
// 	}

// 	// 1. Test runs and fetches requests as expected
// 	// 2. Test fails as expected
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := mockBybit.SendHTTPRequest(tt.APIEndpoint, &GenericRequestParams{}, &result)
// 			if err != nil && err.Error() != tt.expectedErr.Error() {
// 				t.Errorf("expected error got = %v, want = %v", err, tt.expectedErr)
// 			}
// 		})
// 	}
// }

func TestGetInstruments(t *testing.T) {
	ts, instruments, err := mockBybit.GetInstruments(&GenericRequestParams{})

	if err != nil {
		t.Error("failed to call GetInstruments. error: ", err)
	}

	if ts.IsZero() {
		t.Error("failed to get timestamp GetInstruments.", err)
	}

	if len(instruments) == 0 {
		t.Error("failed to get instruments GetInstruments.", err)
	}
}
