package registry

import (
	iface "github.com/hassaanakram/plugin-model/pkg/plugins/pluginiface"
)

var PluginMap map[string]iface.Greeter = map[string]iface.Greeter{}
