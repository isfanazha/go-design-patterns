package domain

type HtmlBuilder struct {
	RootName string
	Root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		RootName: rootName,
		Root: HtmlElement{
			Name:     rootName,
			Text:     "",
			Elements: []HtmlElement{},
		},
	}
}

func (builder *HtmlBuilder) String() string {
	return builder.Root.String()
}

func (builder *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{
		Name:     childName,
		Text:     childText,
		Elements: []HtmlElement{},
	}

	builder.Root.Elements = append(builder.Root.Elements, e)
}

func (builder *HtmlBuilder) AddChildFluent(childName string, childText string) *HtmlBuilder {
	e := HtmlElement{
		Name:     childName,
		Text:     childText,
		Elements: []HtmlElement{},
	}

	builder.Root.Elements = append(builder.Root.Elements, e)

	return builder
}

func (builder *HtmlBuilder) Clear() {
	builder.Root = HtmlElement{
		Name:     builder.RootName,
		Text:     "",
		Elements: []HtmlElement{},
	}
}
