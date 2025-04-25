package abstract_factory

type AbstractFactory interface {
	CreateHtml() HtmlDocument
	CreateWord() WordDocument
}

type HtmlDocument interface {
	ToHtml() string
	Save(fileName string) error
}

type WordDocument interface {
	ToWord() string
	Save(fileName string) error
}

func CreateFactory(name string) AbstractFactory {
	switch name {
	case "good":
		return new(goodFactory)
	case "fast":
		return new(fastFactory)
	default:
		return nil
	}
}
