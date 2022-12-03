package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Passwords struct {
	Entry []struct {
		Name    string    `json:"name"`
		Updated time.Time `json:"updated"`
		Author  string    `json:"author"`
		Content struct {
			ClearPassword string `json:"clear_password"`
			EncrPassword  string `json:"encr_password"`
			Realm         string `json:"realm"`
			Username      string `json:"username"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

// GetPasswords returns basic information about users
func (s *Client) GetPasswords() (*Passwords, error) {
	var passwords Passwords

	response, err := s.Get("/services/storage/passwords?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&passwords)
	if err != nil {
		return nil, err
	}

	return &passwords, nil
}

func PrintPasswords(passwords Passwords) {
	fmt.Println("\nPasswords:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	//SLPUNK_DB SPLUNK_HOME appServerPorts httpport kvStorePort mgmtHostPort pass4SymmKey python.version trustedIP
	t.AppendHeader(table.Row{"#", "Username", "Clear", "Encrypted", "Realm", "Author", "Updated"})

	for k, p := range passwords.Entry {
		t.AppendRows([]table.Row{
			{k, p.Content.Username, p.Content.ClearPassword, p.Content.EncrPassword, p.Content.Realm, p.Author, p.Updated},
		})
	}
	t.Render()
}
