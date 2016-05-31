package bigduration

import (
	"fmt"
)

func Example() {
	examples := []string{
		"1D",
		"1W",
		"1M",
		"1Y",
		"4Y3M2D1h30m15s",
	}

	for _, example := range examples {
		duration, err := ParseDuration(example)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s -> %s\n", example, duration)
	}

	// Output:
	// 1D -> 24h0m0s
	// 1W -> 168h0m0s
	// 1M -> 720h0m0s
	// 1Y -> 8760h0m0s
	// 4Y3M2D1h30m15s -> 37249h30m15s
}
