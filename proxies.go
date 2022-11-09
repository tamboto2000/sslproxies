// Package sslproxies is a scraper for finding proxies based on sslproxies.org
package sslproxies

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/tamboto2000/htmltojson"
)

// Proxy contains information about a proxy
type Proxy struct {
	IP        string `json:"ip"`
	Port      string `json:"port"`
	Code      string `json:"code"`
	Country   string `json:"country"`
	Anonymity string `json:"anonymity"`
	Google    bool   `json:"google"`
	HTTPS     bool   `json:"https"`
}

// GetAll get all proxies
func GetAll() ([]Proxy, error) {
	raw, err := request()
	if err != nil {
		return nil, err
	}

	return parse(raw)
}

// Get get proxies with filters.
// Set count to 0 to get all fetched proxies,
// set code to specify from which country this proxy originate,
// set anon to specify the level of anonymity
func Get(count int, code, anon string) ([]Proxy, error) {
	proxs, err := GetAll()
	if err != nil {
		return nil, err
	}

	newProxs := make([]Proxy, 0)
	for _, p := range proxs {
		if code != "" {
			if p.Code != code {
				continue
			}
		}

		if anon != "" {
			if p.Anonymity != anon {
				continue
			}
		}

		newProxs = append(newProxs, p)
		if count > 0 {
			if len(newProxs) == count {
				return newProxs, nil
			}
		}
	}

	return newProxs, nil
}

func request() ([]byte, error) {
	req, err := http.NewRequest("GET", "https://www.sslproxies.org/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")

	cl := new(http.Client)
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 200 {
		return nil, errors.New(string(body))
	}

	return body, nil
}

func parse(raw []byte) ([]Proxy, error) {
	node, err := htmltojson.ParseBytes(raw)
	if err != nil {
		return nil, err
	}

	node = htmltojson.SearchNode(
		htmltojson.Element,
		"table",
		"",
		"class",
		"table table-striped table-bordered",
		node,
	)

	if node == nil {
		return nil, errors.New("table is missing")
	}

	nodes := htmltojson.SearchAllNode(
		htmltojson.Element,
		"tr",
		"",
		"",
		"",
		node,
	)

	if nodes == nil {
		return nil, errors.New("table is missing")
	}

	items := make([]Proxy, 0)
	for i, node := range nodes {
		if i == 0 {
			continue
		}

		nodes := htmltojson.SearchAllNode(
			htmltojson.Element,
			"td",
			"",
			"",
			"",
			&node,
		)

		if nodes == nil {
			return nil, errors.New("table is missing")
		}

		if node.Child[2].Child == nil {
			continue
		}

		var item Proxy

		if node.Child[0].Child != nil {
			item.IP = node.Child[0].Child[0].Data
		}

		if node.Child[1].Child != nil {
			item.Port = node.Child[1].Child[0].Data
		}

		if node.Child[2].Child != nil {
			item.Code = node.Child[2].Child[0].Data
		}

		if node.Child[3].Child != nil {
			item.Country = node.Child[3].Child[0].Data
		}

		if node.Child[4].Child != nil {
			item.Anonymity = node.Child[4].Child[0].Data
		}

		if node.Child[5].Child != nil {
			if len(node.Child[5].Child) >= 1 {
				if node.Child[5].Child[0].Data == "yes" {
					item.Google = true
				}
			}
		} else {
			item.Google = false
		}

		if node.Child[6].Child[0].Data == "yes" {
			item.HTTPS = true
		} else {
			item.HTTPS = false
		}

		items = append(items, item)
	}

	return items, nil
}
