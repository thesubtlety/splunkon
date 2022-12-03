package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Users struct {
	Entry []struct {
		Name    string    `json:"name"`
		Updated time.Time `json:"updated"`
		Content struct {
			Capabilities        []string `json:"capabilities"`
			Email               string   `json:"email"`
			LastSuccessfulLogin int      `json:"last_successful_login"`
			LockedOut           bool     `json:"locked-out"`
			Realname            string   `json:"realname"`
			Roles               []string `json:"roles"`
			Type                string   `json:"type"`
			Tz                  string   `json:"tz"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

/*
https://<cloudinstance>.splunkcloud.com:8089/services/admin/users
https://<cloudinstance>.splunkcloud.com:8089/services/authentication/users
*/

// GetUsers returns basic information about users
func (s *Client) GetUsers() (*Users, error) {
	var users Users

	response, err := s.Get("/services/authentication/users?output_mode=json")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(bytes.NewReader(response)).Decode(&users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func PrintUsers(users Users) {
	fmt.Println("\nUSERS:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Username", "Email", "Real name", "Last Login", "Locked", "Tz", "Roles", "Type"})

	for k, u := range users.Entry {
		t.AppendRows([]table.Row{
			{k, u.Name, u.Content.Email, u.Content.Realname, u.Content.LastSuccessfulLogin, u.Content.LockedOut, u.Content.Tz, u.Content.Roles, u.Content.Type},
		})
	}
	t.Render()
}
