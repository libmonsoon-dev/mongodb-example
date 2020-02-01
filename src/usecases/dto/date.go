package dto

import (
	"encoding/json"
	"time"
)

type DateDto time.Time

const dateFormat = "1/2/2006 3:04 PM"

func (i DateDto) MarshalJSON() ([]byte, error) {
	tmp := time.Time(i)
	return []byte(tmp.Format(dateFormat)), nil
}

func (i *DateDto) UnmarshalJSON(raw []byte) error {
	var str string

	err := json.Unmarshal(raw, &str)

	if err != nil {
		return err
	}

	tmp, err := time.ParseInLocation(dateFormat, str, location)

	if err != nil {
		return err
	}

	*i = DateDto(tmp)

	return nil
}
