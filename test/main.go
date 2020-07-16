package main

import (
	"encoding/json"
	"os"
	"sslproxies"
)

func main() {
	proxies, err := sslproxies.GetAll()
	if err != nil {
		panic(err.Error())
	}

	f, err := os.Create("proxies.json")
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()
	json.NewEncoder(f).Encode(proxies)
}
