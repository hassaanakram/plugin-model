package plugins

import (
	"fmt"
	iface "github.com/hassaanakram/plugin-model/pkg/plugins/pluginiface"
	"github.com/hassaanakram/plugin-model/pkg/plugins/registry"
)

type EnglishGreeter struct {
	iface.Greeter
}

func (g EnglishGreeter) Greet() {
	fmt.Println("Hello in English")
}

func (g EnglishGreeter) GetLanguage() string {
	return "English"
}

func (g EnglishGreeter) Register() {
	fmt.Println("Registering english greeter...")
	registry.PluginMap[g.GetLanguage()] = EnglishGreeter{}
}

func init() {
	englishGreeter := EnglishGreeter{}
	englishGreeter.Register()
}
