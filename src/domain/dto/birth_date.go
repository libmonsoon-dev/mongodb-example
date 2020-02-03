package dto

import (
	"encoding/json"
	"time"
)

type BirthDate time.Time

const BirthDateFormat = "Monday, January 2, 2006 3:04 PM"

func (i BirthDate) MarshalJSON() ([]byte, error) {
	tmp := time.Time(i)
	return []byte(tmp.Format(BirthDateFormat)), nil
}

func (i *BirthDate) UnmarshalJSON(raw []byte) error {
	var str string

	err := json.Unmarshal(raw, &str)

	if err != nil {
		return err
	}

	tmp, err := time.ParseInLocation(BirthDateFormat, str, location)

	if err != nil {
		return err
	}

	*i = BirthDate(tmp)

	return nil
}
