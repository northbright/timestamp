package timestamp

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Timestamp represents the timestamp for video and SRT file.
// The timestamp starts from 00:00:00.000.
type Timestamp struct {
	hh  int
	mm  int
	ss  int
	mmm int
}

// New parses the given string and returns a timestamp.
// Accepted timestamp string formats:
// 1. "HH:MM:SS"
// 2. "HH:MM:SS.mmm"
// 3. "HH:MM:SS,mmm"
func New(str string) (*Timestamp, error) {
	re := regexp.MustCompile(`^(\d{2}):([0-5][0-9]):([0-5][0-9])([\.|,](\d{3}))?$`)
	arr := re.FindStringSubmatch(str)
	l := len(arr)

	var hh, mm, ss, mmm int

	if l != 6 {
		return nil, fmt.Errorf("incorrect timestamp string format")
	} else {
		hh, _ = strconv.Atoi(arr[1])
		mm, _ = strconv.Atoi(arr[2])
		ss, _ = strconv.Atoi(arr[3])

		if arr[5] != "" {
			mmm, _ = strconv.Atoi(arr[5])
		}

		return &Timestamp{hh: hh, mm: mm, ss: ss, mmm: mmm}, nil
	}
}

// NewFromSecond converts the seconds in float to a timestamp.
func NewFromSecond(second float32) *Timestamp {
	integer, frac := math.Modf(float64(second))
	sec := int(integer)

	str := fmt.Sprintf("%.3f", frac)
	str = strings.TrimPrefix(str, "0.")
	mmm, _ := strconv.Atoi(str)

	hh := sec / 3600
	mm := sec / 3600 % 60
	ss := sec % 3600 % 60

	return &Timestamp{hh: hh, mm: mm, ss: ss, mmm: mmm}
}

// Str returns the timestamp string.
// If forSRT is true, it returns string in the format: "HH:MM:SS,mmm".
// Otherwise, the format is "HH:MM:SS.mmm".
// mmm is the millisecond.
func (ts *Timestamp) Str(forSRT bool) string {
	sep := "."
	if forSRT {
		sep = ","
	}

	return fmt.Sprintf("%02d:%02d:%02d%s%03d", ts.hh, ts.mm, ts.ss, sep, ts.mmm)
}

// String returns the timestamp string in the format: "HH:MM:SS.mmm".
func (ts *Timestamp) String() string {
	return ts.Str(false)
}

// StringForSRT returns the timestamp string for SRT file which is in the format: "HH:MM:SS,mmm".
func (ts *Timestamp) StringForSRT() string {
	return ts.Str(true)
}

// Second returns the number of seconds in float.
func (ts *Timestamp) Second() float32 {
	return float32(ts.hh*3600) + float32(ts.mm*60) + float32(ts.ss) + float32(ts.mmm)/1000
}

// SecondStr returns the second string in the format: "s.mmm" format.
// mmm is the millisecond.
// It may used for the "start" / "end" option of "trim" filter.
func (ts *Timestamp) SecondStr() string {
	return fmt.Sprintf("%d.%03d", ts.hh*3600+ts.mm*60+ts.ss, ts.mmm)
}

// Sub returns a new timestamp which = ts1 - ts2.
// It returns an error if ts1 < ts2.
func (ts1 *Timestamp) Sub(ts2 *Timestamp) (*Timestamp, error) {
	sec1 := ts1.Second()
	sec2 := ts2.Second()

	if sec1 < sec2 {
		return nil, fmt.Errorf("ts1 < ts2")
	}

	return NewFromSecond(sec1 - sec2), nil
}
