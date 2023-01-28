package plugins

import (
	"fmt"
	"github.com/hassaanakram/plugin-model/pkg/greeteriface"
)

func (g main.Greeter) Greet() {
	fmt.Println("Hello in English")
}

func (g main.Greeter) GetLanguage() string {
	return "English"
}