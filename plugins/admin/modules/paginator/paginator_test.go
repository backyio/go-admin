package paginator

import (
	"testing"

	"github.com/backyio/go-admin/modules/config"
	"github.com/backyio/go-admin/plugins/admin/modules/parameter"
	_ "github.com/backyio/go-admin/themes/sword"
)

func TestGet(t *testing.T) {
	p := parameter.BaseParam()
	config.Initialize(&config.Config{Theme: "sword"})
	Get(Config{
		Size:         105,
		Param:        p.SetPage("7"),
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}
