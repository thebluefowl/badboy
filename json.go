package badboy

import "fmt"

var (
	ErrJSONMarshal   = fmt.Errorf("bad json")
	ErrJSONUnmarshal = fmt.Errorf("bad json")
)

type BadJSON struct{}

func (b *BadJSON) MarshalJSON() ([]byte, error) {
	return nil, ErrJSONMarshal
}

func (b *BadJSON) UnmarshalJSON(data []byte) error {
	return ErrJSONUnmarshal
}
