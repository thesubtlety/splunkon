package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/thesubtlety/splunkon/splunkon"
)

func main() {

	if len(os.Args) != 4 {
		log.Fatalln("Usage: ./splunkon https://target:8089/ user pass")
	}
	baseURL := os.Args[1]
	user := os.Args[2]
	password := os.Args[3]

	uri, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Beginning recon on ", uri)

	c := splunkon.New(user, password, *uri)

	checkAuth(c)
	getUsers(c)
	getRoles(c)
	getSAMLRoles(c)
	getSystemInfo(c)
	getServerInfo(c)
	getServerSettings(c)
	getPassword(c)
	getTokens(c)
	getScripts(c)
	getIndexes(c)
	getSourceTypes(c)
	getLocalApps(c)
	getSavedSearches(c)
	getFiredAlerts(c)

	log.Println("Completed. Wrote output to splunkon.json")
}

func checkAuth(c *splunkon.Client) {
	info, err := c.AuthCheck()
	if err != nil {
		log.Fatal(err)
	}
	splunkon.PrintCurrentUser(*info)
}

func getUsers(c *splunkon.Client) {
	users, err := c.GetUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintUsers(*users)
}

func getRoles(c *splunkon.Client) {
	roles, err := c.GetRoles()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintRoles(*roles)
}

func getSAMLRoles(c *splunkon.Client) {
	saml, err := c.GetSAML()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintSAMLMap(*saml)
}

func getSystemInfo(c *splunkon.Client) {
	sytemInfo, err := c.GetSystemInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintSystemInfo(*sytemInfo)
}

func getServerInfo(c *splunkon.Client) {
	serverInfo, err := c.GetServerInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintServerInfo(*serverInfo)
}

func getServerSettings(c *splunkon.Client) {
	serverSettings, err := c.GetServerSettings()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintServerSettings(*serverSettings)
}

func getPassword(c *splunkon.Client) {
	passwords, err := c.GetPasswords()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintPasswords(*passwords)
}

func getTokens(c *splunkon.Client) {
	tokens, err := c.GetTokens()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintTokens(*tokens)
}

func getScripts(c *splunkon.Client) {
	scripts, err := c.GetScripts()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintScripts(*scripts)
}

func getIndexes(c *splunkon.Client) {
	indexes, err := c.GetIndexes()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintIndexes(*indexes)
}

func getSourceTypes(c *splunkon.Client) {
	sourceTypes, err := c.GetSourceTypes()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintSourceTypes(*sourceTypes)
}

func getLocalApps(c *splunkon.Client) {
	localApps, err := c.GetLocalApps()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintLocalApps(*localApps)
}

func getSavedSearches(c *splunkon.Client) {
	savedSearches, err := c.GetSavedSearches()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintSavedSearches(*savedSearches)
}

func getFiredAlerts(c *splunkon.Client) {
	firedAlerts, err := c.GetFiredAlerts()
	if err != nil {
		fmt.Println(err)
		return
	}
	splunkon.PrintFiredAlerts(*firedAlerts)
}
