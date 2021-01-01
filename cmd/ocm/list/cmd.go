/*
Copyright (c) 2020 Red Hat, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package list

import (
	"github.com/openshift-online/ocm-cli/cmd/ocm/list/addon"
	"github.com/openshift-online/ocm-cli/cmd/ocm/list/cluster"
	"github.com/openshift-online/ocm-cli/cmd/ocm/list/idp"
	"github.com/openshift-online/ocm-cli/cmd/ocm/list/ingress"
	"github.com/openshift-online/ocm-cli/cmd/ocm/list/machinepool"
	"github.com/openshift-online/ocm-cli/cmd/ocm/list/org"
	"github.com/openshift-online/ocm-cli/cmd/ocm/list/region"
	"github.com/openshift-online/ocm-cli/cmd/ocm/list/upgradepolicy"
	"github.com/openshift-online/ocm-cli/cmd/ocm/list/user"
	"github.com/openshift-online/ocm-cli/cmd/ocm/list/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list RESOURCE",
	Short: "List all resources of a specific type",
	Long:  "List all resources of a specific type",
}

func init() {
	Cmd.AddCommand(addon.Cmd)
	Cmd.AddCommand(cluster.Cmd)
	Cmd.AddCommand(idp.Cmd)
	Cmd.AddCommand(ingress.Cmd)
	Cmd.AddCommand(org.Cmd)
	Cmd.AddCommand(machinepool.Cmd)
	Cmd.AddCommand(region.Cmd)
	Cmd.AddCommand(upgradepolicy.Cmd)
	Cmd.AddCommand(user.Cmd)
	Cmd.AddCommand(version.Cmd)
}
