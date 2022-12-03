package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Indexes struct {
	Entry []struct {
		Name    string    `json:"name"`
		Updated time.Time `json:"updated"`
	} `json:"entry"`
}

// GetIndexes returns basic information about users
func (s *Client) GetIndexes() (*Indexes, error) {
	var indexes Indexes

	response, err := s.Get("/services/properties/indexes?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&indexes)
	if err != nil {
		return nil, err
	}

	return &indexes, nil
}

func PrintIndexes(indexes Indexes) {
	fmt.Println("\nINDEXES:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Updated"})

	for k, i := range indexes.Entry {
		t.AppendRows([]table.Row{
			{k, i.Name, i.Updated},
		})
	}
	t.Render()
}
