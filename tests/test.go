package tests

import (
	"net/http"
	"strings"
	"testing"

	"github.com/backyio/go-admin/modules/config"
	"github.com/backyio/go-admin/modules/db"
	"github.com/backyio/go-admin/modules/db/dialect"
	"github.com/backyio/go-admin/plugins/admin/modules/table"
	"github.com/backyio/go-admin/tests/common"
	"github.com/gavv/httpexpect"
	gobuffalo "github.com/gobuffalo/buffalo"
)

func Cleaner(config config.DatabaseList) {

	checkStatement := ""

	if config.GetDefault().Driver != "sqlite" {
		if config.GetDefault().Dsn == "" {
			checkStatement = config.GetDefault().Name
		} else {
			checkStatement = config.GetDefault().Dsn
		}
	} else {
		if config.GetDefault().Dsn == "" {
			checkStatement = config.GetDefault().File
		} else {
			checkStatement = config.GetDefault().Dsn
		}
	}

	if !strings.Contains(checkStatement, "test") {
		panic("wrong database")
	}

	var allTables = [...]string{
		"admin_users",
		"admin_user_permissions",
		"admin_session",
		"admin_roles",
		"admin_role_users",
		"admin_role_permissions",
		"admin_role_menu",
		"admin_permissions",
		"admin_operation_log",
		"admin_menu",
	}
	var autoIncrementTable = [...]string{
		"admin_menu",
		"admin_permissions",
		"admin_roles",
		"admin_users",
	}
	var insertData = map[string][]dialect.H{
		"admin_users": {
			{"username": "admin", "name": "admin", "password": "$2a$10$TEDU/aUxLkr2wCxGxI62/.yOtzrzfv426DLLdyha9H2GpWRggB0di", "remember_token": "tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh"},
			{"username": "operator", "name": "operator", "password": "$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.", "remember_token": "tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh"},
		},
		"admin_roles": {
			{"name": "Administrator", "slug": "administrator"},
			{"name": "Operator", "slug": "operator"},
		},
		"admin_permissions": {
			{"name": "All permission", "slug": "*", "http_method": "", "http_path": "*"},
			{"name": "Dashboard", "slug": "dashboard", "http_method": "GET,PUT,POST,DELETE", "http_path": "/"},
		},
		"admin_menu": {
			{"parent_id": 0, "type": 1, "order": 2, "title": "Admin", "icon": "fa-tasks", "uri": ""},
			{"parent_id": 1, "type": 1, "order": 2, "title": "Users", "icon": "fa-users", "uri": "/info/manager"},
			{"parent_id": 0, "type": 1, "order": 3, "title": "test2 menu", "icon": "fa-angellist", "uri": "/example/test"},
			{"parent_id": 1, "type": 1, "order": 4, "title": "Permission", "icon": "fa-ban", "uri": "/info/permission"},
			{"parent_id": 1, "type": 1, "order": 5, "title": "Menu", "icon": "fa-bars", "uri": "/menu"},
			{"parent_id": 1, "type": 1, "order": 6, "title": "Operation log", "icon": "fa-history", "uri": "/info/op"},
			{"parent_id": 0, "type": 1, "order": 1, "title": "Dashboard", "icon": "fa-bar-chart", "uri": "/"},
			{"parent_id": 0, "type": 1, "order": 7, "title": "User", "icon": "fa-users", "uri": "/info/user"},
		},
		"admin_role_users": {
			{"user_id": 1, "role_id": 1},
			{"user_id": 2, "role_id": 2},
		},
		"admin_user_permissions": {
			{"user_id": 1, "permission_id": 1},
			{"user_id": 2, "permission_id": 2},
		},
		"admin_role_permissions": {
			{"role_id": 1, "permission_id": 1},
			{"role_id": 1, "permission_id": 2},
			{"role_id": 2, "permission_id": 2},
		},
		"admin_role_menu": {
			{"role_id": 1, "menu_id": 1},
			{"role_id": 1, "menu_id": 7},
			{"role_id": 2, "menu_id": 7},
			{"role_id": 1, "menu_id": 8},
			{"role_id": 2, "menu_id": 8},
			{"role_id": 1, "menu_id": 3},
		},
	}
	conn := db.GetConnectionByDriver(config.GetDefault().Driver).InitDB(config)
	// clean data
	for _, t := range allTables {
		_ = db.WithDriver(conn).Table(t).Delete()
	}
	// reset auto increment
	switch config.GetDefault().Driver {
	case db.DriverPostgresql:
		for _, t := range autoIncrementTable {
			checkErr(conn.Exec(`ALTER SEQUENCE ` + t + `_myid_seq RESTART WITH  1`))
		}
	}
	// insert data
	for t, data := range insertData {
		for _, d := range data {
			checkErr(db.WithDriver(conn).Table(t).Insert(d))
		}
	}
}

func BlackBoxTestSuitOfBuiltInTables(t *testing.T, fn HandlerGenFn, config config.DatabaseList) {
	BlackBoxTestSuit(t, fn, config, nil, Cleaner, common.Test)
}

func checkErr(_ interface{}, err error) {
	if err != nil {
		panic(err)
	}
}

func BlackBoxTestSuit(t *testing.T, fn HandlerGenFn,
	config config.DatabaseList,
	gens table.GeneratorList,
	cleaner DataCleaner,
	tester Tester) {
	// Clean Data
	cleaner(config)
	// Test
	tester(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(fn(config, gens)),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}

type Tester func(e *httpexpect.Expect)
type DataCleaner func(config config.DatabaseList)
type HandlerGenFn func(config config.DatabaseList, gens table.GeneratorList) http.Handler
type GoBuffaloHttpHandlerGenFn func(config config.DatabaseList, gens table.GeneratorList) gobuffalo.Handler
