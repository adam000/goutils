package page

type Page struct {
	Title       string
	Javascript  []string
	Stylesheets []string
	Vars        map[string]interface{}
	SiteTitle   string
}

func NewPage() Page {
	siteTitle := config.siteTitle
	if config.siteTitle == "" {
		siteTitle = "adam zero dot net"
	}
	ret := Page{
		SiteTitle: siteTitle,
		Vars:      make(map[string]interface{}),
	}

	return ret
}

func (p *Page) GetTitle() string {
	return p.Title
}

func (p *Page) SetTitle(title string) {
	p.Title = title
}

func (p *Page) GetJSFiles() []string {
	return p.Javascript
}

func (p *Page) AddJSFiles(file ...string) {
	p.Javascript = append(p.Javascript, file...)
}

func (p *Page) GetCSSFiles() []string {
	return p.Stylesheets
}

func (p *Page) AddCSSFiles(file ...string) {
	p.Stylesheets = append(p.Stylesheets, file...)
}

func (p *Page) AddVar(name string, variable interface{}) {
	p.Vars[name] = variable
}
