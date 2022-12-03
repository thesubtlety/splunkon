package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Scripts struct {
	Entry []struct {
		Name    string    `json:"name"`
		Updated time.Time `json:"updated"`
		Acl     struct {
			App   string `json:"app"`
			Owner string `json:"owner"`
		} `json:"Acl"`
		Content struct {
			Disabled      bool        `json:"disabled"`
			Group         string      `json:"group"`
			HostResolved  string      `json:"host_resolved"`
			Index         string      `json:"index"`
			Interval      int         `json:"interval"`
			PassAuth      string      `json:"passAuth"`
			PythonVersion interface{} `json:"python.version"`
			RunOnlyOne    interface{} `json:"run_only_one"`
			Sourcetype    string      `json:"sourcetype"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

func (s *Client) GetScripts() (*Scripts, error) {
	var scripts Scripts

	response, err := s.Get("/services/admin/script?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&scripts)
	if err != nil {
		return nil, err
	}

	return &scripts, nil
}

func PrintScripts(scripts Scripts) {
	fmt.Println("\nSCRIPTS:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "App", "Author", "Disabled"})

	for k, s := range scripts.Entry {
		t.AppendRows([]table.Row{
			{k, s.Name, s.Acl.App, s.Acl.Owner, s.Content.Disabled},
		})
	}
	t.Render()
}
