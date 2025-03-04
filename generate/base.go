package generate

import (
	"github.com/chaosannals/goctl-python/template"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func GenBase(dir string, api *spec.ApiSpec) error {
	data := &template.PyBaseTemplateData{
		ClientName: lo.PascalCase(api.Service.Name),
	}
	return template.GenFile(dir, "base.py", template.ApiBase, data)
}
