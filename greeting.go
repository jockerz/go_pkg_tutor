package github.com/jockerz/go_pkg_tutor


import "github.com/k0kubun/pp/v3"


func Greet(name string) {
    pp.Fprintf("Greeting %s\n", name)
}
