package main

var langs = map[string]map[string]string{
	"en": {
		"cn": "Chinese",
		"en": "English",
		"jp": "Japanese",
		"tc": "Traditional Chinese",

		"web.database settings":     "Database Settings",
		"web.installation settings": "Installation Settings",
		"web.application settings":  "Application General Settings",
		"web.optional Settings":     "Optional Settings",

		"web.database type":     "Database Type",
		"web.database host":     "Host",
		"web.database user":     "User",
		"web.database password": "Password",
		"web.database name":     "Database Name",
		"web.database file":     "Path",
		"web.database port":     "port",
		"web.database schema":   "Schema",

		"web.theme":             "Theme",
		"web.language":          "Language",
		"web.web framework":     "Web Framework",
		"web.module name":       "Module Name",
		"web.http port":         "HTTP Port",
		"web.url prefix":        "Url Prefix",
		"web.website title":     "Website Title",
		"web.login page logo":   "Login Page Logo",
		"web.sidebar logo":      "SideBar Logo",
		"web.sidebar mini logo": "SideBar Mini Logo",
		"web.use orm":           "Use ORM",
		"web.no use":            "No use",
		"web.input":             "Input",

		"web.simplified chinese":  "Simplified Chinese",
		"web.traditional chinese": "Traditional Chinese",
		"web.english":             "English",
		"web.japanese":            "Japanese",

		"web.where the framework sql data install to":            "Where the framework sql data will be installed to。",
		"web.the file path of sqlite3 database. ":                "The file path of SQLite3 database. ",
		"web.please use absolute path when you start as service": "Please use absolute path when you start as service.",
		"web.module name is the path of go module":               "Module name is the path of go module.",
		"web.port number which application will listen on":       "Port number which application will listen on.",
		"web.url prefix of the running application":              "Url prefix of the running application.",

		"web.official website":                 "Official Website",
		"web.current version":                  "Current Version",
		"web.goadmin web installation program": "Admin Web Installation Program",
		"web.installation page":                "Installation Page",
		"web.install now":                      "Install GoAdmin",

		"web.result":          "Installation Result",
		"web.ok":              "Ok",
		"web.wrong parameter": "Wrong parameter",
		"web.install success": "Install Success~~🍺🍺",
	},
}

var defaultLang = "en"

func setDefaultLangSet(set string) {
	if set != "" && (set == "cn" || set == "en") {
		defaultLang = set
	}
}

func local(lang string) func(string) string {
	if _, ok := langs[defaultLang]; ok {
		return func(msg string) string {
			return langs[lang][msg]
		}
	}
	return nil
}

func getWord(msg string) string {
	if word, ok := langs[defaultLang][msg]; ok {
		return word
	}
	return msg
}
