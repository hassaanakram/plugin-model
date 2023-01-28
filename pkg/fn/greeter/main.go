package main

import (
	"github.com/hassaanakram/plugin-model/pkg/plugins/registry"
)

func main() {
	inputLang := "Spanish"
	registry.PluginMap[inputLang].Greet()
}