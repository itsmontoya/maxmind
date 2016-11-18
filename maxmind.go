package maxmind

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	// ErrInvalidEntry is returned when an invalid entry is provided
	ErrInvalidEntry = errors.New("invalid entry")
)

// Filter will iterate through an io.Reader and find all the matching entries for the provided filters
func Filter(r io.Reader, city, state, country string) (es []Entry) {
	var (
		e   Entry
		err error

		scn = bufio.NewScanner(r)
	)

	for scn.Scan() {
		if e, err = GetEntry(scn.Bytes()); err != nil {
			continue
		}

		if len(city) > 0 && e.City != city {
			continue
		}

		if len(state) > 0 && e.State != state {
			continue
		}

		if len(country) > 0 && e.Country != country {
			continue
		}

		es = append(es, e)
	}

	return
}

// FilterFile will filter a file
func FilterFile(fileLoc, city, state, country string) (es []Entry, err error) {
	var f *os.File
	if f, err = os.Open(fileLoc); err != nil {
		return
	}

	es = Filter(f, city, state, country)
	return
}

// GetEntry will get an entry
func GetEntry(b []byte) (e Entry, err error) {
	var spl [][]byte
	if spl = bytes.Split(b, []byte{','}); len(spl) != 7 {
		err = ErrInvalidEntry
		return
	}

	if e.Lat, err = strconv.ParseFloat(string(spl[5]), 64); err != nil {
		return
	}

	if e.Lon, err = strconv.ParseFloat(string(spl[6]), 64); err != nil {
		return
	}

	e.Country = strings.ToUpper(string(spl[0]))
	e.City = string(spl[2])
	e.State = string(spl[3])

	return
}

// Entry represents an entry
type Entry struct {
	Country string
	City    string
	State   string
	Lat     float64
	Lon     float64
}

/*
us,wyodak,Wyodak,WY,,44.2913889,-105.3791667

*/
