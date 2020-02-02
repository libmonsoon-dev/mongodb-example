package dto

type WinStatus stringInt

func (i WinStatus) MarshalJSON() ([]byte, error) {
	tmp := stringInt(i)
	return tmp.MarshalJSON()
}

func (i *WinStatus) UnmarshalJSON(raw []byte) error {
	tmp := stringInt(*i)
	err := tmp.UnmarshalJSON(raw)
	if err != nil {
		return err
	}
	*i = WinStatus(tmp)
	return nil
}
