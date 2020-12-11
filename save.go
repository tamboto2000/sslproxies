package sslproxies

import (
	"encoding/json"
	"io/ioutil"
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

// Load load proxies from ./proxies.json
func Load() ([]Proxy, error) {
	return load("./proxies.json")
}

// LoadFromPath load proxies from path
func LoadFromPath(path string) ([]Proxy, error) {
	return load(path)
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

func load(path string) ([]Proxy, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	raw, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	proxs := new([]Proxy)
	if err = json.Unmarshal(raw, proxs); err != nil {
		return nil, err
	}

	return *proxs, nil
}
