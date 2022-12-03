package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type UsrRoleSamlMap struct {
	Entry []struct {
		Title   string    `json:"title"`
		Updated time.Time `json:"updated"`
		Content struct {
			Email    string   `json:"email"`
			Realname string   `json:"realname"`
			Roles    []string `json:"roles"`
			Type     string   `json:"type"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

func (s *Client) GetSAML() (*UsrRoleSamlMap, error) {
	var samlMap UsrRoleSamlMap

	response, err := s.Get("/services/properties/authentication/userToRoleMap_SAML?output_mode=json")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(bytes.NewReader(response)).Decode(&samlMap)
	if err != nil {
		return nil, err
	}

	return &samlMap, nil
}

func PrintSAMLMap(samlMap UsrRoleSamlMap) {
	fmt.Println("\nSAML USERS:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Email", "Name", "Type", "Roles"})

	for k, u := range samlMap.Entry {
		t.AppendRows([]table.Row{
			{k, u.Title, u.Content.Email, u.Content.Realname, u.Content.Type, u.Content.Roles},
		})
	}
	t.Render()
}
