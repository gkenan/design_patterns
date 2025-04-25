package abstract_factory

import "fmt"

type goodFactory struct{}

func (g goodFactory) CreateHtml() HtmlDocument {
	return new(goodHtmlDocument)
}

func (g goodFactory) CreateWord() WordDocument {
	return new(goodWordDocument)
}

type goodHtmlDocument struct {
	content string
}

func (g *goodHtmlDocument) ToHtml() string {
	return "Good Html Document: " + g.content
}

func (g *goodHtmlDocument) Save(fileName string) error {
	fmt.Println("Saving Good Html Document", fileName)
	g.content = fileName
	return nil
}

type goodWordDocument struct {
	content string
}

func (g *goodWordDocument) ToWord() string {
	return "Good Word Document: " + g.content
}

func (g *goodWordDocument) Save(fileName string) error {
	fmt.Println("Saving Good Word Document", fileName)
	g.content = fileName
	return nil
}
