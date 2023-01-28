package plugins

import (
	"fmt"
	iface"github.com/hassaanakram/plugin-model/pkg/plugins/pluginiface"
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