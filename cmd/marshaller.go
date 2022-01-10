package cmd

import (
	"encoding/json"
	"io"

	"github.com/TylerBrock/colorjson"
)

func FormatJSON(body io.ReadCloser) (string, error) {
	var j interface{}

	var err = json.NewDecoder(body).Decode(&j)
	if err != nil {
		panic(err)
	}

	f := colorjson.NewFormatter()
	f.Indent = 2

	formattedJSON, err := f.Marshal(j)
	return string(formattedJSON), err
}
