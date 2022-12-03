package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Roles struct {
	Entry []struct {
		Name    string `json:"name"`
		Content struct {
			Capabilities                  []string      `json:"capabilities"`
			DefaultApp                    string        `json:"defaultApp"`
			FederatedProviders            []interface{} `json:"federatedProviders"`
			GrantableRoles                []interface{} `json:"grantable_roles"`
			ImportedCapabilities          []string      `json:"imported_capabilities"`
			ImportedRoles                 []string      `json:"imported_roles"`
			ImportedSrchFilter            string        `json:"imported_srchFilter"`
			ImportedSrchIndexesAllowed    []string      `json:"imported_srchIndexesAllowed"`
			ImportedSrchIndexesDisallowed []interface{} `json:"imported_srchIndexesDisallowed"`
			SrchFilter                    string        `json:"srchFilter"`
			SrchIndexesDefault            []string      `json:"srchIndexesDefault"`
			SrchIndexesAllowed            []string      `json:"srchIndexesAllowed"`
			SrchIndexesDisallowed         []interface{} `json:"srchIndexesDisallowed"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

// GetUsers returns basic information about users
func (s *Client) GetRoles() (*Roles, error) {
	var roles Roles

	response, err := s.Get("/services/authorization/roles?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&roles)
	if err != nil {
		return nil, err
	}

	return &roles, nil
}

func PrintRoles(roles Roles) {
	fmt.Println("\nROLES:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Default Search Idx", "Capabilities", "Imported Capabilities", "Admin Capabilities"})

	for k, r := range roles.Entry {
		t.AppendRows([]table.Row{
			{k, r.Name, r.Content.SrchIndexesDefault, len(r.Content.Capabilities), len(r.Content.ImportedCapabilities), strings.Contains(strings.Join(r.Content.Capabilities, ",")+strings.Join(r.Content.ImportedCapabilities, ","), "admin")},
		})
	}
	t.Render()
}
