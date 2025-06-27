package timestamp_test

import (
	"fmt"
	"log"

	"github.com/northbright/timestamp"
)

func ExampleTimestamp() {
	arr := []string{
		"00:00:00",
		"10:20:30.500",
		"20:30:40,900",
	}

	for _, str := range arr {
		ts, err := timestamp.New(str)
		if err != nil {
			log.Printf("New() error: %v", err)
			return
		}
		fmt.Printf("%s -> String(): %s, StringForSRT(): %s, SecondStr(): %s\n", str, ts.String(), ts.StringForSRT(), ts.SecondStr())
	}

	fArr := []float32{
		0.0,
		3.14,
		3882.46000,
	}

	for _, f := range fArr {
		ts, err := timestamp.NewFromSecond(f)
		if err != nil {
			log.Printf("NewFromSecondStr() error: %v", err)
			return
		}
		fmt.Printf("%f -> String(): %s, StringForSRT(): %s, SecondStr(): %s\n", f, ts.String(), ts.StringForSRT(), ts.SecondStr())
	}

	// Output:
	// 00:00:00 -> String(): 00:00:00.000, StringForSRT(): 00:00:00,000, SecondStr(): 0.000
	// 10:20:30.500 -> String(): 10:20:30.500, StringForSRT(): 10:20:30,500, SecondStr(): 37230.500
	// 20:30:40,900 -> String(): 20:30:40.900, StringForSRT(): 20:30:40,900, SecondStr(): 73840.900
	// 0.000000 -> String(): 00:00:00.000, StringForSRT(): 00:00:00,000, SecondStr(): 0.000
	// 3.140000 -> String(): 00:00:03.140, StringForSRT(): 00:00:03,140, SecondStr(): 3.140
	// 3882.459961 -> String(): 01:01:42.460, StringForSRT(): 01:01:42,460, SecondStr(): 3702.460
}
