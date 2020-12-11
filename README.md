# SSLProxies

[![Go Reference](https://pkg.go.dev/badge/github.com/tamboto2000/sslproxies.svg)](https://pkg.go.dev/github.com/tamboto2000/sslproxies)

SSLProxies is a scraper for finding proxies based on sslproxies.org.

### Installation

SSLProxies require Golang v14 or up
```sh
$ go get github.com/tamboto2000/sslproxies
```

### Example
```go

package main

import (
	"github.com/tamboto2000/sslproxies"
)

func main() {
	// get 20 proxies from United States of America with anonymity level of Elite Proxy
	proxies, err := sslproxies.Get(20, sslproxies.UnitedStates, sslproxies.Elite)
	if err != nil {
		panic(err.Error())
	}

	// save proxies
	if err := sslproxies.Save(proxies); err != nil {
		panic(err.Error())
	}

	// save proxies to custom path
	if err := sslproxies.SaveToPath(proxies, "my_proxies.json"); err != nil {
		panic(err.Error())
	}
}

```

License
----

MIT

