package plugins

import (
	"fmt"
	iface "github.com/hassaanakram/plugin-model/pkg/plugins/pluginiface"
	"github.com/hassaanakram/plugin-model/pkg/plugins/registry"
)

type SpanishGreeter struct {
	iface.Greeter
}

func (g SpanishGreeter) Greet() {
	fmt.Println("Hello in Spanish pls dont say no pls pls pls")
}

func (g SpanishGreeter) GetLanguage() string {
	return "Spanish"
}

func (g SpanishGreeter) Register() {
	fmt.Println("Registering spanish greeter...")
	registry.PluginMap[g.GetLanguage()] = SpanishGreeter{}
}

func init() {
	spanishGreeter := SpanishGreeter{}
	spanishGreeter.Register()
}
