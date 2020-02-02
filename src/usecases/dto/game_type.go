package dto

type GameType stringInt

func (i GameType) MarshalJSON() ([]byte, error) {
	tmp := stringInt(i)
	return tmp.MarshalJSON()
}

func (i *GameType) UnmarshalJSON(raw []byte) error {
	tmp := stringInt(*i)
	err := tmp.UnmarshalJSON(raw)
	if err != nil {
		return err
	}
	*i = GameType(tmp)
	return nil
}
