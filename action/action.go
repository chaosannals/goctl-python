package action

import (
	"github.com/chaosannals/goctl-python/generate"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

func Python(ctx *cli.Context) error {
	plugin, err := plugin.NewPlugin()
	if err != nil {
		return err
	}

	return generate.PythonCommand(plugin)
}
