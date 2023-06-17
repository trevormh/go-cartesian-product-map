package cartesian

import (
	"sync"
)

// Iter takes interface-slices and returns a channel, receiving cartesian products
func Iter(params ...map[string][]interface{}) chan map[string]interface{} {
	// create channel
	c := make(chan map[string]interface{})
	// create waitgroup
	var wg sync.WaitGroup
	// call iterator
	wg.Add(1)

	// flatten any params passed in with multiple keys
	var combined []map[string][]interface{}
	for _,m := range params {
		for key,val := range m {
			flattened := map[string][]interface{} {
				key: val,
			}
			combined = append(combined, flattened)
		}
	}
	iterate(&wg, c, map[string]interface{} {}, combined...)

	// call channel-closing go-func
	go func() { wg.Wait(); close(c) }()
	// return channel
	return c
}

// private, recursive Iteration-Function
func iterate(wg *sync.WaitGroup, channel chan map[string]interface{}, result map[string]interface{}, params ...map[string][]interface{}) {
	// dec WaitGroup when finished
	defer wg.Done()
	// no more params left?
	if len(params) == 0 {
		// send result to channel
		channel <- result
		return
	}
	// shift first param
	p, params := params[0], params[1:]

	var pkey string
	for key := range p {
		pkey = key
		break
	}

	// iterate over it
	for i := 0; i < len(p[pkey]); i++ {
		// inc WaitGroup
		wg.Add(1)
		
		// create copy of result
		resultCopy := map[string]interface{}{}
		for k,v := range result {
			resultCopy[k] = v
		}
		resultCopy[pkey] = p[pkey][i]

		// call self with remaining params
		go iterate(wg, channel, resultCopy, params...)
	}
}