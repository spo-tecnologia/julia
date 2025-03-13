package enums

type SampleEnum int8

const (
	SampleEnumType1 SampleEnum = 1
	SampleEnumType2 SampleEnum = 2
)

func (s SampleEnum) String() string {
	switch s {
	case SampleEnumType1:
		return "Type 1"
	case SampleEnumType2:
		return "Type 2"
	}

	return "Desconhecido"
}

func (s SampleEnum) Value() int8 {
	return int8(s)
}
