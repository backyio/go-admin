<p align="center">
  <a href="https://github.com/backyio/go-admin">
    <img width="48%" alt="go-admin" src="http://quick.go-admin.cn/official/assets/imgs/github_logo.png">
  </a>
</p>

<p align="center">
    the missing golang data admin panel builder tool.
</p>

<p align="center">
    Inspired by <a href="https://github.com/z-song/laravel-admin" target="_blank">laravel-admin</a>
</p>

## Preface

GoAdmin is a toolkit to help you build a data visualization admin panel for your golang app.

![interface](http://file.go-admin.cn/introduction/interface_en_3.png)

## Features

- 🚀 **Fast**: build a production admin panel app in **ten** minutes.
- 🎨 **Theming**: beautiful ui themes supported(default adminlte, more themes are coming.)
- 🔢 **Plugins**: many plugins to use(more useful and powerful plugins are coming.)
- ✅ **Rbac**: out of box rbac auth system.
- ⚙️ **Frameworks**: support most of the go web frameworks.


Following three steps to run it.

Note: now you can quickly start by doing like this.

```shell
$ go install github.com/backyio/go-admin/adm@latest
$ mkdir new_project && cd new_project
$ adm init
```

Or (use adm whose version higher or equal than v1.2.16)

```shell
$ mkdir new_project && cd new_project
$ go install github.com/backyio/go-admin/adm@latest
$ adm init web
```

### Step 2: create main.go

<details><summary>main.go</summary>
<p>

```go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/backyio/go-admin/adapter/gin"
	_ "github.com/backyio/go-admin/modules/db/drivers/mysql"
	"github.com/backyio/go-admin/engine"
	"github.com/backyio/go-admin/plugins/admin"
	"github.com/backyio/go-admin/modules/config"
	"github.com/backyio/go-admin/themes/adminlte"
	"github.com/backyio/go-admin/template"
	"github.com/backyio/go-admin/template/chartjs"
	"github.com/backyio/go-admin/template/types"
	"github.com/backyio/go-admin/examples/datamodel"
	"github.com/backyio/go-admin/modules/language"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	// global config
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:         "127.0.0.1",
				Port:         "3306",
				User:         "root",
				Pwd:          "root",
				Name:         "goadmin",
				MaxIdleConns: 50,
				MaxOpenConns: 150,
				ConnMaxLifetime: time.Hour,
				Driver:       "mysql",
			},
        	},
		UrlPrefix: "admin",
		// STORE is important. And the directory should has permission to write.
		Store: config.Store{
		    Path:   "./uploads", 
		    Prefix: "uploads",
		},
		Language: language.EN,
		// debug mode
		Debug: true,
		// log file absolute path
		InfoLogPath: "/var/logs/info.log",
		AccessLogPath: "/var/logs/access.log",
		ErrorLogPath: "/var/logs/error.log",
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	// add component chartjs
	template.AddComp(chartjs.NewChart())

	_ = eng.AddConfig(&cfg).
		AddGenerators(datamodel.Generators).
	        // add generator, first parameter is the url prefix of table when visit.
    	        // example:
    	        //
    	        // "user" => http://localhost:9033/admin/info/user
    	        //		
		AddGenerator("user", datamodel.GetUserTable).
		Use(r)
	
	// customize your pages
	eng.HTML("GET", "/admin", datamodel.GetContent)

	_ = r.Run(":9033")
}
```

</p>
</details>

More framework examples: [https://github.com/backyio/go-admin/tree/master/examples](https://github.com/backyio/go-admin/tree/master/examples)

### Step 3: run

```shell
GO111MODULE=on go run main.go
```