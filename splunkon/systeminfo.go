package splunkon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type SystemInfo struct {
	Entry []struct {
		Name    string    `json:"name"`
		Updated time.Time `json:"updated"`
		Content struct {
			CPUArch              string `json:"cpu_arch"`
			NumberOfCores        int    `json:"numberOfCores"`
			NumberOfVirtualCores int    `json:"numberOfVirtualCores"`
			OsBuild              string `json:"os_build"`
			OsName               string `json:"os_name"`
			OsNameExtended       string `json:"os_name_extended"`
			OsVersion            string `json:"os_version"`
			PhysicalMemoryMB     int    `json:"physicalMemoryMB"`
		} `json:"content"`
	} `json:"entry"`
	Messages []interface{} `json:"messages"`
}

// AuthCheck returns basic information about the user if creds are valid
func (s *Client) GetSystemInfo() (*SystemInfo, error) {
	var systemInfo SystemInfo

	response, err := s.Get("/services/admin/system-info?output_mode=json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&systemInfo)
	if err != nil {
		return nil, err
	}

	return &systemInfo, nil
}

func PrintSystemInfo(sysInfo SystemInfo) {
	fmt.Println("\nSYSTEM INFO:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "OSName", "OSNameExtended", "OSVersion", "OSBuild", "Arch", "Cores", "Mem"})

	for k, s := range sysInfo.Entry {
		t.AppendRows([]table.Row{
			{k, s.Name, s.Content.OsName, s.Content.OsNameExtended, s.Content.OsVersion, s.Content.OsBuild, s.Content.CPUArch, s.Content.NumberOfCores, s.Content.PhysicalMemoryMB},
		})
	}
	t.Render()
}
