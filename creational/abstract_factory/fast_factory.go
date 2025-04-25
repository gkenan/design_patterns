package abstract_factory

import "fmt"

type fastFactory struct{}

func (f fastFactory) CreateHtml() HtmlDocument {
	return new(fastHtmlDocument)
}

func (f fastFactory) CreateWord() WordDocument {
	return new(fastWordDocument)
}

type fastHtmlDocument struct {
	content string
}

func (f *fastHtmlDocument) ToHtml() string {
	return "Fast Html Document: " + f.content
}

func (f *fastHtmlDocument) Save(fileName string) error {
	fmt.Println("Saving Fast Html Document: " + fileName)
	f.content = fileName
	return nil
}

type fastWordDocument struct {
	content string
}

func (f *fastWordDocument) ToWord() string {
	return "Fast Word Document: " + f.content
}

func (f *fastWordDocument) Save(fileName string) error {
	fmt.Println("Saving Fast Word Document: " + fileName)
	f.content = fileName
	return nil
}
