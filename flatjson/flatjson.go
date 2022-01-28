// Package flatjson has a method that converts an input JSON object into a flattened version.
//
// For example, th object `{"abc": {"def": 1}}` will become `{"abc.def": 1}`.
package flatjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// Params define source of the input JSON and the destination where to write the result.
type Params struct {
	Reader io.Reader
	Writer io.Writer
}

// FlatJSON reads JSON object from the params.Reader, flattens it, and finally writes
// the result into params.Writer.
func FlatJSON(params *Params) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(params.Reader)
	if err != nil {
		return fmt.Errorf("failed to read from the reader: %w", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		return fmt.Errorf("failed to unmarshal json object: %w", err)
	}

	out := make(map[string]interface{})
	for n, v := range data {
		if err := traverse(v, n, out); err != nil {
			return err
		}
	}

	enc := json.NewEncoder(params.Writer)
	return enc.Encode(out)
}

func traverse(obj interface{}, path string, out map[string]interface{}) error {
	switch v := obj.(type) {
	case float64,
		bool,
		string:
		out[path] = v
	case map[string]interface{}:
		var pathPrefix string
		if path != "" {
			pathPrefix = fmt.Sprintf("%s.", path)
		}
		for name, val := range v {
			subPath := fmt.Sprintf("%s%s", pathPrefix, name)
			if err := traverse(val, subPath, out); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unexpected type %T for value %v", obj, v)
	}
	return nil
}
