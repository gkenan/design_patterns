package abstract_factory

import "fmt"

type FastHtmlDocument struct {
	content string
}

func (f *FastHtmlDocument) ToHtml() string {
	return "Fast Html Document: " + f.content
}

func (f *FastHtmlDocument) Save(fileName string) error {
	fmt.Println("Saving Fast Html Document: " + fileName)
	return nil
}

type FastWordDocument struct {
	content string
}

func (f *FastWordDocument) ToWord() string {
	return "Fast Word Document: " + f.content
}

func (f *FastWordDocument) Save(fileName string) error {
	fmt.Println("Saving Fast Word Document: " + fileName)
	return nil
}
