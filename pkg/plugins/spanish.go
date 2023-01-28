package plugins

import (
	"fmt"
	iface"github.com/hassaanakram/plugin-model/pkg/plugins/pluginiface"
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