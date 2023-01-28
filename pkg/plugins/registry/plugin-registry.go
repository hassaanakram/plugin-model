package registry

import (
	iface"github.com/hassaanakram/plugin-model/pkg/plugins/pluginiface"
	"github.com/hassaanakram/plugin-model/pkg/plugins"
)

var PluginMap map[string]iface.Greeter = map[string]iface.Greeter {
	"English": plugins.EnglishGreeter{},
	"French": plugins.FrenchGreeter{},
	"Spanish": plugins.SpanishGreeter{},
}