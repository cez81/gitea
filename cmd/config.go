// Copyright 2016 The Gogs Authors. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"os"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/setting"

	"github.com/urfave/cli"
)

var (
	// CmdAdmin represents the available config sub-command.
	CmdConfig = cli.Command{
		Name:  "config",
		Usage: "Dump the complete config file",
		Description: `Allow using internal logic of Gitea without hacking into the source code
to make automatic initialization process more smoothly`,
		Subcommands: []cli.Command{
			subcmdViewConfig,
		},
	}

	subcmdViewConfig = cli.Command{
		Name:   "view",
		Usage:  "View the complete config file including defaults",
		Action: viewConfig,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config, c",
				Value: "custom/conf/app.ini",
				Usage: "Custom configuration file path",
			},
		},
	}

)

func viewConfig(c *cli.Context) error {
	setting.NewContext()
	models.LoadConfigs()

	setting.Cfg.WriteToIndent(os.Stdout, "\t")
	return nil
}
