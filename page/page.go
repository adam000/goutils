package page

type Page struct {
	title       string
	javascript  []string
	stylesheets []string
}

func (p Page) getTitle() string {
	return p.title
}

func (p Page) setTitle(title string) {
	p.title = title
}

func (p Page) getJSFiles() []string {
	return p.javascript
}

func (p Page) addJSFile(file ...string) {
	p.javascript = append(p.javascript, file...)
}

func (p Page) getCSSFiles() []string {
	return p.stylesheets
}

func (p Page) addCSSFile(file ...string) {
	p.stylesheets = append(p.stylesheets, file...)
}
