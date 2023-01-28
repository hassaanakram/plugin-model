package plugins

import (
	"fmt"
	iface "github.com/hassaanakram/plugin-model/pkg/plugins/pluginiface"
	"github.com/hassaanakram/plugin-model/pkg/plugins/registry"
)

type FrenchGreeter struct {
	iface.Greeter
}

func (g FrenchGreeter) Greet() {
	fmt.Println("Bonjour this is French pls dont say otherwise")
}

func (g FrenchGreeter) GetLanguage() string {
	return "French"
}

func (g FrenchGreeter) Register() {
	fmt.Println("Registering french greeter...")
	registry.PluginMap[g.GetLanguage()] = FrenchGreeter{}
}

func init() {
	frenchGreeter := FrenchGreeter{}
	frenchGreeter.Register()
}
