package pluginiface

type Greeter interface {
	Greet()
	GetLanguage() string
	Register()
}
