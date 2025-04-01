package factory_method

/*
┌─────────────┐      ┌─────────────┐
│   Product   │      │   Factory   │
└─────────────┘      └─────────────┘
       ▲                    ▲
       │                    │
┌─────────────┐      ┌─────────────┐
│ ProductImpl │◀─ ─ ─│ FactoryImpl │
└─────────────┘      └─────────────┘
*/

type NumberFactory interface {
	Parse(s string) (any, error)
}

func Get(s string) NumberFactory {
	switch s {
	case "int":
		return new(iInt)
	case "float":
		return new(iFloat)
	default:
		return nil
	}
}
