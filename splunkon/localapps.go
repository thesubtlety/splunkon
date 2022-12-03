package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type LocalApps struct {
	Entry []struct {
		Name    string    `json:"name"`
		Updated time.Time `json:"updated"`
		Author  string    `json:"author"`
		Content struct {
			Author                     string `json:"author"`
			CheckForUpdates            bool   `json:"check_for_updates"`
			Configured                 bool   `json:"configured"`
			Core                       bool   `json:"core"`
			Description                string `json:"description"`
			Disabled                   bool   `json:"disabled"`
			Label                      string `json:"label"`
			ManagedByDeploymentClient  bool   `json:"managed_by_deployment_client"`
			ShowInNav                  bool   `json:"show_in_nav"`
			StateChangeRequiresRestart bool   `json:"state_change_requires_restart"`
			Version                    string `json:"version"`
			Visible                    bool   `json:"visible"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

// GetLocalApps returns basic information about LocalApps
func (s *Client) GetLocalApps() (*LocalApps, error) {
	var localApps LocalApps

	response, err := s.Get("/services/apps/local?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&localApps)
	if err != nil {
		return nil, err
	}

	return &localApps, nil
}

func PrintLocalApps(localApps LocalApps) {
	fmt.Println("\nAPPS:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Label", "Version", "Disabled", "Core"})

	for k, l := range localApps.Entry {
		t.AppendRows([]table.Row{
			{k, l.Name, l.Content.Label, l.Content.Version, l.Content.Disabled, l.Content.Core},
		})
	}
	t.Render()
}
