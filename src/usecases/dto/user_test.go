package dto

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestUserDto_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input  []byte
		expect UserDto
	}{
		{
			[]byte(`{}`),
			UserDto{},
		},
		{
			[]byte(`{
				"email":"Alessandra_Carter7577@infotech44.tech",
				"last_name":"Carter",
				"country":"Peru",
				"city":"Berna",
				"gender":"Female",
				"birth_date":"Sunday, May 16, 0488 5:52 AM"
			}`),
			UserDto{
				"Alessandra_Carter7577@infotech44.tech",
				"Carter",
				"Peru",
				"Berna",
				"Female",
				BirthDateDto(time.Date(488, 5, 16, 5, 52, 0, 0, location)),
			},
		},
		{
			[]byte(`{
				"email":"Celia_Graves1431@joiniaa.com",
				"last_name":"Graves",
				"country":"Mongolia",
				"city":"Rome",
				"gender":"Female",
				"birth_date":"Thursday, June 1, 8862 5:03 AM"
			}`),
			UserDto{
				"Celia_Graves1431@joiniaa.com",
				"Graves",
				"Mongolia",
				"Rome",
				"Female",
				BirthDateDto(time.Date(8862, 6, 1, 5, 3, 0, 0, location)),
			},
		},
	}
	for _, test := range tests {
		user := UserDto{}
		err := json.Unmarshal(test.input, &user)

		if err != nil {
			t.Fatal(fmt.Errorf("unexpected error: %w", err))
		}

		if user != test.expect {
			t.Fatalf("actual != expect\nactual:\n%#v\nexpect:\n%#v\n", user, test.expect)
		}

	}
}
