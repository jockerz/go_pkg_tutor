package greeting


import "github.com/k0kubun/pp/v3"


func Greet(name string) {
    pp.Printf("Greeting %s\n", name)
}
