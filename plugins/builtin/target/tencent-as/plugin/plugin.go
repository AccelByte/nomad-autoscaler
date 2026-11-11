// Copyright AccelByte Inc. 2026
// SPDX-License-Identifier: MPL-2.0

// Package plugin wraps github.com/huxunpei/nomad-autoscaler-tencent as a
// builtin target plugin for the AccelByte nomad-autoscaler fork.
// It is registered alongside aws-asg, azure-vmss, gce-mig, and
// oci-instance-pool so all providers ship in a single binary.
package plugin

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad-autoscaler/plugins"
	tencent "github.com/huxunpei/nomad-autoscaler-tencent/plugin"
)

// PluginConfig is used by the autoscaler agent to embed this plugin at compile
// time (internal plugin mode). External deployments can still override it by
// placing a tencent-as binary in plugin_dir.
var PluginConfig = &plugins.InternalPluginConfig{
	Factory: func(l hclog.Logger) interface{} { return tencent.Factory(l) },
}
