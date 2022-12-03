package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type ServerInfo struct {
	Entry []struct {
		Name    string `json:"name"`
		Content struct {
			ActiveLicenseGroup    string `json:"activeLicenseGroup"`
			ActiveLicenseSubgroup string `json:"activeLicenseSubgroup"`
			AddOns                struct {
				Hadoop struct {
					Parameters struct {
						ErpType  string `json:"erp_type"`
						GUID     string `json:"guid"`
						MaxNodes string `json:"maxNodes"`
					} `json:"parameters"`
					Type string `json:"type"`
				} `json:"hadoop"`
			} `json:"addOns"`
			Build                  string      `json:"build"`
			CPUArch                string      `json:"cpu_arch"`
			EaiACL                 interface{} `json:"eai:acl"`
			FederatedSearchEnabled bool        `json:"federated_search_enabled"`
			FipsMode               bool        `json:"fips_mode"`
			GUID                   string      `json:"guid"`
			HealthInfo             string      `json:"health_info"`
			HealthVersion          int         `json:"health_version"`
			Host                   string      `json:"host"`
			HostFqdn               string      `json:"host_fqdn"`
			HostResolved           string      `json:"host_resolved"`
			IsForwarding           bool        `json:"isForwarding"`
			IsFree                 bool        `json:"isFree"`
			IsTrial                bool        `json:"isTrial"`
			KvStoreStatus          string      `json:"kvStoreStatus"`
			LicenseKeys            []string    `json:"licenseKeys"`
			LicenseSignature       string      `json:"licenseSignature"`
			LicenseState           string      `json:"licenseState"`
			LicenseLabels          []string    `json:"license_labels"`
			ManagerGUID            string      `json:"manager_guid"`
			ManagerURI             string      `json:"manager_uri"`
			MasterGUID             string      `json:"master_guid"`
			MasterURI              string      `json:"master_uri"`
			MaxUsers               int64       `json:"max_users"`
			Mode                   string      `json:"mode"`
			NumberOfCores          int         `json:"numberOfCores"`
			NumberOfVirtualCores   int         `json:"numberOfVirtualCores"`
			OsBuild                string      `json:"os_build"`
			OsName                 string      `json:"os_name"`
			OsNameExtended         string      `json:"os_name_extended"`
			OsVersion              string      `json:"os_version"`
			PhysicalMemoryMB       int         `json:"physicalMemoryMB"`
			ProductType            string      `json:"product_type"`
			RtsearchEnabled        bool        `json:"rtsearch_enabled"`
			ServerName             string      `json:"serverName"`
			ServerRoles            []string    `json:"server_roles"`
			ShuttingDown           string      `json:"shutting_down"`
			StartupTime            int         `json:"startup_time"`
			StaticAssetID          string      `json:"staticAssetId"`
			Version                string      `json:"version"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

// GetServerInfo returns basic information about the server
func (s *Client) GetServerInfo() (*ServerInfo, error) {
	var serverInfo ServerInfo

	response, err := s.Get("/services/admin/server-info?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&serverInfo)
	if err != nil {
		return nil, err
	}

	return &serverInfo, nil
}

func PrintServerInfo(serverInfo ServerInfo) {
	fmt.Println("\nSERVER INFO:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Host", "Splunk Version", "License", "Roles", "Start Time"})

	for k, s := range serverInfo.Entry {
		t.AppendRows([]table.Row{
			{k, s.Content.HostFqdn, s.Content.Version, s.Content.LicenseLabels, s.Content.ServerRoles, s.Content.StartupTime},
		})
	}
	t.Render()
}
