package main

import "flag"

func main() {
	flag.Parse()
	args := flag.Args()
	if args[0] == "client" {
		client()
	} else if args[0] == "server" {
		server()
	}
}
