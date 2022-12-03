package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Tokens struct {
	Entry []struct {
		Name    string `json:"name"`
		Content struct {
			Claims struct {
				Aud   string   `json:"aud"`
				Exp   int      `json:"exp"`
				Iat   int      `json:"iat"`
				Idp   string   `json:"idp"`
				Iss   string   `json:"iss"`
				Nbr   int      `json:"nbr"`
				Roles []string `json:"roles"`
				Sub   string   `json:"sub"`
			} `json:"claims"`
			Headers struct {
				Alg  string `json:"alg"`
				Kid  string `json:"kid"`
				Ttyp string `json:"ttyp"`
				Ver  string `json:"ver"`
			} `json:"headers"`
			LastUsed   int    `json:"lastUsed"`
			LastUsedIP string `json:"lastUsedIp"`
			Status     string `json:"status"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

// GetTokens returns basic information about the server
func (s *Client) GetTokens() (*Tokens, error) {
	var tokens Tokens

	response, err := s.Get("/services/authorization/tokens?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&tokens)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

func PrintTokens(tokens Tokens) {
	fmt.Println("\nTOKENS:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Aud", "Exp", "Iat", "Iss", "Sub", "Last Used", "Last IP", "Status"})

	for k, tok := range tokens.Entry {
		t.AppendRows([]table.Row{
			{k, "..." + tok.Name[:8], tok.Content.Claims.Aud, tok.Content.Claims.Exp, tok.Content.Claims.Iat, tok.Content.Claims.Iss, tok.Content.Claims.Sub, tok.Content.LastUsed, tok.Content.LastUsedIP, tok.Content.Status},
		})
	}
	t.Render()
}
