package sslproxies

import (
	"encoding/json"
	"os"
)

// Save saves proxies to ./proxies.json
func Save(proxs []Proxy) error {
	return save(proxs, "./proxies.json")
}

// SaveToPath saves proxies to path
func SaveToPath(proxs []Proxy, path string) error {
	return save(proxs, path)
}

// save to a file
func save(proxs []Proxy, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(proxs)
}
