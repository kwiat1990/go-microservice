package data

import (
	"encoding/json"
	"io"
)

// While working with io.Reader/Writer it's better to use json.NewEncoder/NewDecoder
// instead of json.Marshal/Unmarshal as it's slightly more performant.
//
// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
