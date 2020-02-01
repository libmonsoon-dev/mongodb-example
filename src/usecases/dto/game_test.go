package dto

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestGameDto_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input  []byte
		expect GameDto
	}{
		{
			[]byte(`{}`),
			GameDto{},
		},
		{
			[]byte(`{
				"points_gained":"1933",
				"win_status":"0",
				"game_type":"17",
				"created":"1/17/2019 9:54 AM"
			}`),
			GameDto{
				1933,
				0,
				17,
				DateDto(time.Date(2019, 1, 17, 9, 54, 0, 0, location)),
			},
		},
		{
			[]byte(`{
				"points_gained":"1076",
				"win_status":"1",
				"game_type":"5",
				"created":"5/24/2019 5:36 PM"
			}`),
			GameDto{
				1076,
				1,
				5,
				DateDto(time.Date(2019, 5, 24, 17, 36, 0, 0, location)),
			},
		},
	}
	for _, test := range tests {
		game := GameDto{}
		err := json.Unmarshal(test.input, &game)

		if err != nil {
			t.Fatal(fmt.Errorf("unexpected error: %w", err))
		}

		if game != test.expect {
			t.Fatalf("actual != expect\nactual:\n%#v\nexpect:\n%#v\n", game, test.expect)
		}

	}
}
