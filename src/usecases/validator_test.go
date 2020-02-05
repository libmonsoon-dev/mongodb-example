package usecases

import (
	"reflect"
)

type FieldLevelStub string

func (f FieldLevelStub) Top() reflect.Value {
	panic("implement me")
}

func (f FieldLevelStub) Parent() reflect.Value {
	panic("implement me")
}

func (f FieldLevelStub) Field() reflect.Value {
	return reflect.ValueOf(f)
}

func (f FieldLevelStub) FieldName() string {
	panic("implement me")
}

func (f FieldLevelStub) StructFieldName() string {
	panic("implement me")
}

func (f FieldLevelStub) Param() string {
	panic("implement me")
}

func (f FieldLevelStub) GetTag() string {
	panic("implement me")
}

func (f FieldLevelStub) ExtractType(field reflect.Value) (value reflect.Value, kind reflect.Kind, nullable bool) {
	panic("implement me")
}

func (f FieldLevelStub) GetStructFieldOK() (reflect.Value, reflect.Kind, bool) {
	panic("implement me")
}

func (f FieldLevelStub) GetStructFieldOKAdvanced(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool) {
	panic("implement me")
}

func (f FieldLevelStub) GetStructFieldOK2() (reflect.Value, reflect.Kind, bool, bool) {
	panic("implement me")
}

func (f FieldLevelStub) GetStructFieldOKAdvanced2(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool, bool) {
	panic("implement me")
}

//func TestEmailValidator(t *testing.T) {
//	tests := []struct {
//		input FieldLevelStub
//		want  bool
//	}{
//		{
//			"test@test.org",
//			true,
//		},
//	}
//	for _, test := range tests {
//		if got := emailValidator(test.input); got != test.want {
//			t.Errorf("emailValidator() = %v, want %v", got, test.want)
//		}
//	}
//}
