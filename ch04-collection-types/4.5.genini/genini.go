package main

import (
	"fmt"
	"sort"
)

func main() {
	iniData := map[string]map[string]string{
		"App": {
			"Name":    "Iceweasel",
			"Profile": "mozilla/firefox",
			"Vendor":  "Mozilla",
			"Version": "3.5.16",
		},
		"Gecko": {
			"MaxVersion": "1.9.1.*",
			"MinVersion": "1.9.1",
		},
		"XRE": {
			"EnableExtensionManager": "1",
			"EnableProfileMigrator":  "0",
		},
	}

	PrintIni(iniData)
}

func PrintIni(iniData map[string]map[string]string) {
	groups := make([]string, 0, len(iniData))
	for g, _ := range iniData {
		groups = append(groups, g)
	}
	sort.Strings(groups)

	for _, group := range groups {
		fmt.Printf("\n[%s]\n", group)
		settings := make([]string, 0, len(iniData[group]))
		for setting, _ := range iniData[group] {
			settings = append(settings, setting)
		}
		for _, s := range settings {
			fmt.Printf("%s=%s\n", s, iniData[group][s])
		}
	}
}
