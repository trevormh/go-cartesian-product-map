# go-cartesian-product-map

a package for building [cartesian products](https://en.wikipedia.org/wiki/Cartesian_product) for maps in golang. This is a modified version of [schwarmco/go-cartesian-product](https://github.com/schwarmco/go-cartesian-product) to enable creating cartesian products while retaining map keys.

keep in mind, that because [how golang handles maps](https://blog.golang.org/go-maps-in-action#TOC_7.) your results will not be "in order"

## Installation

In order to start, `go get` this repository:

```
go get trevormh/go-cartesian-product-map
```

## Usage

```go
import (
    "fmt"
    "github.com/trevormh/go-cartesian-product-map
)

func main() {
    
	a := map[string][]interface{} {
		"some_key": {1, 2, "c"},
		"another_key": {"ten","nine","eight"},
	}
	
	b := map[string][]interface{} {
		"b_key": {10,11,12},
	}

	c := map[string][]interface{} {
		"key-c": {"test"},
	}

	d := cartesian.Iter(a, b, c)

	// receive products through channel
	for product := range d {
		fmt.Println(product)
	}

	// Unordered Output:
	// map[another_key:eight b_key:12 key-c:test some_key:c]
	// map[another_key:eight b_key:12 key-c:test some_key:2]
	// map[another_key:eight b_key:12 key-c:test some_key:1]
	// map[another_key:ten b_key:12 key-c:test some_key:2]
	// map[another_key:nine b_key:12 key-c:test some_key:1]
	// map[another_key:eight b_key:10 key-c:test some_key:1]
	// map[another_key:eight b_key:11 key-c:test some_key:1]
	// map[another_key:ten b_key:10 key-c:test some_key:2]
	// map[another_key:ten b_key:11 key-c:test some_key:2]
	// map[another_key:nine b_key:10 key-c:test some_key:1]
	// map[another_key:nine b_key:11 key-c:test some_key:1]
	// map[another_key:nine b_key:12 key-c:test some_key:2]
	// map[another_key:nine b_key:10 key-c:test some_key:2]
	// map[another_key:nine b_key:11 key-c:test some_key:2]
	// map[another_key:ten b_key:12 key-c:test some_key:1]
	// map[another_key:eight b_key:10 key-c:test some_key:2]
	// map[another_key:ten b_key:12 key-c:test some_key:c]
	// map[another_key:ten b_key:10 key-c:test some_key:c]
	// map[another_key:eight b_key:11 key-c:test some_key:2]
	// map[another_key:ten b_key:10 key-c:test some_key:1]
	// map[another_key:ten b_key:11 key-c:test some_key:1]
	// map[another_key:ten b_key:11 key-c:test some_key:c]
	// map[another_key:nine b_key:12 key-c:test some_key:c]
	// map[another_key:nine b_key:10 key-c:test some_key:c]
	// map[another_key:nine b_key:11 key-c:test some_key:c]
	// map[another_key:eight b_key:10 key-c:test some_key:c]
	// map[another_key:eight b_key:11 key-c:test some_key:c]
}
```
