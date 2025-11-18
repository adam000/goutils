package page

var config = struct {
	defaultStylesheets []Import
	defaultJavascript  []Import
}{}

func SetDefaultStylesheets(stylesheets []Import) {
	config.defaultStylesheets = stylesheets
}

func GetDefaultStylesheets() []Import {
	return config.defaultStylesheets
}

func SetDefaultJavascript(javascript []Import) {
	config.defaultJavascript = javascript
}

func GetDefaultJavascript() []Import {
	return config.defaultJavascript
}
