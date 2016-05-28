package bigduration

import (
	"testing"
	"time"
)

var cases = map[string]string{
	"1ns": "1ns",
	"1us": "1us",
	"1µs": "1µs",
	"1ms": "1ms",
	"1s":  "1s",
	"1m":  "1m",
	"1h":  "1h",
	"1D":  "24h",
	"1W":  "168h",
	"1M":  "720h",
	"1Y":  "8760h",
	"7Y6M5W4D3h2m1s1ms2us3ns": "66579h2m1s1ms2us3ns",
}

func Test_ParseDuration(t *testing.T) {
	for input, standard := range cases {
		actual, err := ParseDuration(input)
		if err != nil {
			t.Fatal(err)
		}

		expected, err := time.ParseDuration(standard)
		if err != nil {
			t.Fatal(err)
		}

		if actual != expected {
			t.Fatalf("expected %s, but got %s", expected, actual)
		}
	}
}
