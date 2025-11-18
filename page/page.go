package page

type Page struct {
	Title       string
	Javascript  []Import
	Stylesheets []Import
	Vars        map[string]interface{}
	SiteTitle   string
}

type Import struct {
	Url    string
	SriSha string
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

func (p *Page) GetJsFiles() []Import {
	return p.Javascript
}

func (p *Page) AddJsFiles(file ...string) {
	for _, f := range file {
		imp := Import{
			Url: f,
		}
		p.AddJsFile(imp)
	}
}

func (p *Page) AddJsFile(i Import) {
	p.Javascript = append(p.Javascript, i)
}

func (p *Page) GetCssFiles() []Import {
	return p.Stylesheets
}

func (p *Page) AddCssFiles(file ...string) {
	for _, f := range file {
		imp := Import{
			Url: f,
		}
		p.AddCssFile(imp)
	}
}

func (p *Page) AddCssFile(i Import) {
	p.Stylesheets = append(p.Stylesheets, i)
}

func (p *Page) AddVar(name string, variable interface{}) {
	p.Vars[name] = variable
}
