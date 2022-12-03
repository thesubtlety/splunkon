package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Settings struct {
	Links struct {
	} `json:"links"`
	Origin    string    `json:"origin"`
	Updated   time.Time `json:"updated"`
	Generator struct {
		Build   string `json:"build"`
		Version string `json:"version"`
	} `json:"generator"`
	Entry []struct {
		Name    string `json:"name"`
		Content struct {
			SplunkDb                        string `json:"SPLUNK_DB"`
			SplunkHome                      string `json:"SPLUNK_HOME"`
			AppServerPorts                  []int  `json:"appServerPorts"`
			DataFederationDisabled          bool   `json:"dataFederationDisabled"`
			EnableSplunkWebSSL              bool   `json:"enableSplunkWebSSL"`
			Host                            string `json:"host"`
			HostResolved                    string `json:"host_resolved"`
			Httpport                        int    `json:"httpport"`
			InvalidateSessionTokensOnLogout bool   `json:"invalidateSessionTokensOnLogout"`
			KvStoreDisabled                 bool   `json:"kvStoreDisabled"`
			KvStorePort                     int    `json:"kvStorePort"`
			LogoutCacheRefreshInterval      string `json:"logoutCacheRefreshInterval"`
			MgmtHostPort                    int    `json:"mgmtHostPort"`
			MinFreeSpace                    int    `json:"minFreeSpace"`
			Pass4SymmKey                    string `json:"pass4SymmKey"`
			PythonVersion                   string `json:"python.version"`
			ServerName                      string `json:"serverName"`
			SessionTimeout                  string `json:"sessionTimeout"`
			Startwebserver                  bool   `json:"startwebserver"`
			TrustedIP                       string `json:"trustedIP"`
		} `json:"content"`
	} `json:"entry"`
	Messages []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"messages"`
}

func (s *Client) GetServerSettings() (*Settings, error) {
	var settings Settings

	response, err := s.Get("/services/server/settings?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&settings)
	if err != nil {
		return nil, err
	}

	for _, m := range settings.Messages {
		if strings.Contains(m.Type, "ERROR") {
			return &settings, fmt.Errorf(string(response))
		}
	}

	return &settings, nil
}

func PrintServerSettings(serverSettings Settings) {
	fmt.Println("\nSERVER SETTINGS:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "DB Path", "Splunk Home", "App Port", "HTTP Port", "KV Port", "Mgmt Port", "SymmKeyPass", "Python Ver", "Trusted IP"})

	for k, s := range serverSettings.Entry {
		t.AppendRows([]table.Row{
			{k, s.Content.SplunkDb, s.Content.SplunkHome, s.Content.AppServerPorts, s.Content.Httpport, s.Content.KvStorePort, s.Content.MgmtHostPort, s.Content.Pass4SymmKey, s.Content.PythonVersion, s.Content.TrustedIP},
		})
	}
	t.Render()
}
