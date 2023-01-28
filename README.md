### Golang Compiled-in Extensions
Plugin model programming allows you to write code that can be easily extended by third-party developers without having to change the original code. 
Writing plugins in other languages is relatively easier however, in Golang it can be a bit tricky. There are a few methods that different projects use and one of
those is ```Compiled-in Extensions```. This is an example of this method.

In this sample implementation, you have a ```Greeter``` that says hello in different languages. Each language is implemented as a plugin. The ```Greeter``` is a 
Golang interface that satisfies the following methods:
- ```Greet()```: Prints Greetings
- ```GetLanguage() string```: Returns the language of the Greeter
- ```Register()```: Registers the new Greeter language plugin with the main program

In order to implement a new ```Greeter``` in a different language, you can follow the existing code in ```pkg/plugins -> english.go, spanish.go, french.go```
The new ```Greeter``` must have the above three methods to satisfy the requriements of the base interface.
