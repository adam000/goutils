package page

var config = struct {
	defaultStylesheets []string
	defaultJavascript  []string
	siteTitle          string
}{}

func SetDefaultStylesheets(stylesheets []string) {
	config.defaultStylesheets = stylesheets
}

func GetDefaultStylesheets() []string {
	return config.defaultStylesheets
}

func SetDefaultJavascript(javascript []string) {
	config.defaultJavascript = javascript
}

func GetDefaultJavascript() []string {
	return config.defaultJavascript
}

func SetSiteTitle(title string) {
	config.siteTitle = title
}
