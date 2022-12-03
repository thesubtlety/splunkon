package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type FiredAlerts struct {
	Entry []struct {
		Name    string    `json:"name"`
		Updated time.Time `json:"updated"`
		Author  string    `json:"author"`
		Content struct {
			TriggeredAlertCount int `json:"triggered_alert_count"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

// GetFiredAlerts returns basic information about users
func (s *Client) GetFiredAlerts() (*FiredAlerts, error) {
	var firedAlerts FiredAlerts

	response, err := s.Get("/services/alerts/fired_alerts?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&firedAlerts)
	if err != nil {
		return nil, err
	}

	return &firedAlerts, nil
}

func PrintFiredAlerts(firedAlerts FiredAlerts) {
	fmt.Println("\nFIRED ALERTS:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Author", "Updated", "TriggeredCount"})

	for k, f := range firedAlerts.Entry {
		t.AppendRows([]table.Row{
			{k, f.Name, f.Author, f.Updated, f.Content.TriggeredAlertCount},
		})
	}
	t.Render()
}
