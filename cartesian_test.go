package cartesian_test

import (
	"fmt"

	"github.com/trevormh/go-cartesian-product-map"
)

func ExampleIter() {
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