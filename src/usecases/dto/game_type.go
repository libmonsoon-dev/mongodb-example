package dto

type GameTypeDto stringInt

func (i GameTypeDto) MarshalJSON() ([]byte, error) {
	tmp := stringInt(i)
	return tmp.MarshalJSON()
}

func (i *GameTypeDto) UnmarshalJSON(raw []byte) error {
	tmp := stringInt(*i)
	err := tmp.UnmarshalJSON(raw)
	if err != nil {
		return err
	}
	*i = GameTypeDto(tmp)
	return nil
}
