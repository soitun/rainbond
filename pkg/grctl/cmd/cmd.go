// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	//"os"

	"github.com/Sirupsen/logrus"
	conf "github.com/goodrain/rainbond/cmd/grctl/option"
	"github.com/goodrain/rainbond/pkg/grctl/clients"
	"github.com/urfave/cli"
)

//GetCmds GetCmds
func GetCmds() []cli.Command {
	cmds := []cli.Command{}
	cmds = append(cmds, NewCmdBatchStop())
	cmds = append(cmds, NewCmdStartService())
	cmds = append(cmds, NewCmdStopService())
	cmds = append(cmds, NewCmdTenant())
	cmds = append(cmds, NewCmdTenantRes())
	cmds = append(cmds, NewCmdNode())
	cmds = append(cmds, NewCmdNodeRes())
	cmds = append(cmds, NewCmdExec())
	cmds = append(cmds, NewCmdLog())
	cmds = append(cmds, NewCmdEvent())
	cmds = append(cmds, NewCmdGet())


	cmds = append(cmds, NewCmdInit())

	cmds = append(cmds, NewCmdAddNode())

	cmds = append(cmds, NewCmdCheckComputeServices())
	cmds = append(cmds, NewCmdComputeGroup())

	cmds = append(cmds, NewCmdBaseManageGroup())
	cmds = append(cmds, NewCmdManageGroup())
	cmds = append(cmds, NewCmdCheckManageServices())
	cmds = append(cmds, NewCmdCheckManageBaseServices())

	//cmds = append(cmds, NewCmdPlugin())
	//todo
	return cmds
}
func Common(c *cli.Context) {
	config, err := conf.LoadConfig(c)
	if err != nil {
		logrus.Warnf("Load config file error.", err.Error())
		//os.Exit(1)
		return
	}

	if err := clients.InitClient(*config.Kubernets); err != nil {
		logrus.Warnf("error config k8s")
	}
	//clients.SetInfo(config.RegionAPI.URL, config.RegionAPI.Token)
	if err := clients.InitRegionClient(*config.RegionAPI); err != nil {
		logrus.Warnf("error config region")
	}
	if err := clients.InitNodeClient("http://127.0.0.1:6100"); err != nil {
		logrus.Warnf("error config region")
	}



}