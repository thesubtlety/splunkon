package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type SavedSearches struct {
	Entry []struct {
		Name    string    `json:"name"`
		Updated time.Time `json:"updated"`
		Author  string    `json:"author"`
		Content struct {
			ActionEmail            bool        `json:"action.email"`
			ActionEmailSendresults interface{} `json:"action.email.sendresults"`
			ActionEmailTo          string      `json:"action.email.to"`
			ActionScript           bool        `json:"action.script"`
			Actions                string      `json:"actions"`
			AlertManagedBy         string      `json:"alert.managedBy"`
			AlertSeverity          int         `json:"alert.severity"`
			CronSchedule           string      `json:"cron_schedule"`
			Description            string      `json:"description"`
			Disabled               bool        `json:"disabled"`
			IsScheduled            bool        `json:"is_scheduled"`
			IsVisible              bool        `json:"is_visible"`
			NextScheduledTime      string      `json:"next_scheduled_time"`
			QualifiedSearch        string      `json:"qualifiedSearch"`
			RunOnStartup           bool        `json:"run_on_startup"`
			ScheduleAs             string      `json:"schedule_as"`
			ScheduleWindow         string      `json:"schedule_window"`
			Search                 string      `json:"search"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

// GetsavedSearches returns basic information about users
func (s *Client) GetSavedSearches() (*SavedSearches, error) {
	var savedSearches SavedSearches

	response, err := s.Get("/services/admin/savedsearch?output_mode=json")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(bytes.NewReader(response)).Decode(&savedSearches)
	if err != nil {
		return nil, err
	}

	return &savedSearches, nil
}

func PrintSavedSearches(savedSearches SavedSearches) {
	fmt.Println("\nSAVED SEARCHES:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Updated", "Scheduled", "Search"})

	for k, s := range savedSearches.Entry {
		t.AppendRows([]table.Row{
			{k, s.Name, s.Updated, s.Content.IsScheduled, s.Content.Search},
		})
	}
	t.Render()
}
