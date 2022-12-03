package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

// resourceInfo represents an api info response struct
type AuthenticationCurrentContext struct {
	Entry []struct {
		Name    string `json:"name"`
		Content struct {
			Capabilities        []string `json:"capabilities"`
			DefaultApp          string   `json:"defaultApp"`
			Email               string   `json:"email"`
			LastSuccessfulLogin int      `json:"last_successful_login"`
			LockedOut           bool     `json:"locked-out"`
			Realname            string   `json:"realname"`
			Roles               []string `json:"roles"`
			Tz                  string   `json:"tz"`
			Username            string   `json:"username"`
		} `json:"content"`
	} `json:"entry"`
	Messages []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"messages"`
}

// AuthCheck returns basic information about the user if creds are valid
func (s *Client) AuthCheck() (*AuthenticationCurrentContext, error) {
	var currentContext AuthenticationCurrentContext

	response, err := s.Get("/services/authentication/current-context?output_mode=json")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(bytes.NewReader(response)).Decode(&currentContext)
	if err != nil {
		return nil, err
	}

	for _, m := range currentContext.Messages {
		if strings.Contains(m.Type, "ERROR") {
			return nil, fmt.Errorf(string(response))
		}
	}

	return &currentContext, nil
}

func PrintCurrentUser(user AuthenticationCurrentContext) {
	fmt.Println("\nCURRENT USER:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Username", "Email", "Real name", "Last Login", "Locked", "Tz", "Roles"})

	for k, u := range user.Entry {
		t.AppendRows([]table.Row{
			{k, u.Content.Username, u.Content.Email, u.Content.Realname, u.Content.LastSuccessfulLogin, u.Content.LockedOut, u.Content.Tz, u.Content.Roles},
		})
	}
	t.Render()
}
