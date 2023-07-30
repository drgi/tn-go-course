package main

import "fmt"

func main() {
	var iface interface{}
	var f func()
	var m map[string]interface{}
	var p *int = nil
	var iface2 interface{} = p
	cycle(iface, f, m, p, iface2)
}

func cycle(ifaces ...interface{}) {
	for i, iface := range ifaces {
		if iface == nil {
			fmt.Println(i)
		}
	}
}
