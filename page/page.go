package page

type Page struct {
	Title       string
	Javascript  []string
	Stylesheets []string
	Vars        map[string]interface{}
	SiteTitle   string
}

func NewPage() Page {
	ret := Page{
		Javascript:  config.defaultJavascript,
		Stylesheets: config.defaultStylesheets,
		Vars:        make(map[string]interface{}),
	}

	return ret
}

func (p *Page) GetTitle() string {
	return p.Title
}

func (p *Page) SetTitle(title string) {
	p.Title = title
}

func (p *Page) GetSiteTitle() string {
	return p.SiteTitle
}

func (p *Page) SetSiteTitle(siteTitle string) {
	p.SiteTitle = siteTitle
}

func (p *Page) GetJsFiles() []string {
	return p.Javascript
}

func (p *Page) AddJsFiles(file ...string) {
	p.Javascript = append(p.Javascript, file...)
}

func (p *Page) GetCssFiles() []string {
	return p.Stylesheets
}

func (p *Page) AddCssFiles(file ...string) {
	p.Stylesheets = append(p.Stylesheets, file...)
}

func (p *Page) AddVar(name string, variable interface{}) {
	p.Vars[name] = variable
}
