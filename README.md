# SSLProxies

[![Go Reference](https://pkg.go.dev/badge/github.com/tamboto2000/sslproxies.svg)](https://pkg.go.dev/github.com/tamboto2000/sslproxies)

SSLProxies is a scraper for finding proxies based on sslproxies.org.

### Installation

SSLProxies require Golang v14 or up
```sh
$ go get github.com/tamboto2000/sslproxies
```

# Examples

### Example 1
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

### Example 2
Load proxies from saved file
```go
package main

import (
	"github.com/tamboto2000/sslproxies"
)

func main() {
	// get proxies from saved files
	proxs, err := sslproxies.Load()
	if err != nil {
		panic(err.Error())
	}

	// use the proxs...
}
```

### Example 3
Load proxies from custom path
```go
package main

import (
	"github.com/tamboto2000/sslproxies"
)

func main() {
	// get proxies from custom path
	proxs, err := sslproxies.LoadFromPath("my_proxies.json")
	if err != nil {
		panic(err.Error())
	}

	// use the proxs...
}
```

License
----

MIT

