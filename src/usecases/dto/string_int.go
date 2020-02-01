package dto

import (
	"encoding/json"
	"strconv"
)

type stringInt int

func (i stringInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.Itoa(int(i)))
}

func (i *stringInt) UnmarshalJSON(raw []byte) error {
	var str string
	err := json.Unmarshal(raw, &str)

	if err != nil {
		return err
	}

	tmp, err := strconv.Atoi(str)

	if err != nil {
		return err
	}

	*i = stringInt(tmp)
	return nil
}
