package dto

type WinStatusDto stringInt

func (i WinStatusDto) MarshalJSON() ([]byte, error) {
	tmp := stringInt(i)
	return tmp.MarshalJSON()
}

func (i *WinStatusDto) UnmarshalJSON(raw []byte) error {
	tmp := stringInt(*i)
	err := tmp.UnmarshalJSON(raw)
	if err != nil {
		return err
	}
	*i = WinStatusDto(tmp)
	return nil
}
