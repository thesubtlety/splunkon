package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type SourceTypes struct {
	Entry []struct {
		Name    string    `json:"name"`
		Updated time.Time `json:"updated"`
	} `json:"entry"`
}

// GetIndexes returns basic information about SourceTypes
func (s *Client) GetSourceTypes() (*SourceTypes, error) {
	var sourceTypes SourceTypes

	response, err := s.Get("/services/properties/sourcetypes?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&sourceTypes)
	if err != nil {
		return nil, err
	}

	return &sourceTypes, nil
}

func PrintSourceTypes(sourceTypes SourceTypes) {
	fmt.Println("\nSOURCE TYPES:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Updated"})

	for k, s := range sourceTypes.Entry {
		t.AppendRows([]table.Row{
			{k, s.Name, s.Updated},
		})
	}
	t.Render()
}
