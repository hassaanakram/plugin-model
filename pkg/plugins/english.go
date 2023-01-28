package plugins

import (
	"fmt"
	iface"github.com/hassaanakram/plugin-model/pkg/plugins/pluginiface"
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