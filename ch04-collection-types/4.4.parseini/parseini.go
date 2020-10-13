package main

import (
	"fmt"
	"strings"
)

func main() {
	iniData := []string{
		"; Cut down copy of Mozilla application.ini file", "",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0", "EnableExtensionManager=1",
	}

	fmt.Println(ParseIni(iniData))
}

func ParseIni(iniData []string) map[string]map[string]string {
	iniMap := make(map[string]map[string]string)
	var iniSubMap map[string]string
	for _, line := range iniData {
		if len(line) == 0 || line[0] == ';' {
			continue
		}

		if line[0] == '[' {
			iniMap[line[1:len(line)-1]] = make(map[string]string)
			iniSubMap = iniMap[line[1:len(line)-1]]
		} else if iniSubMap != nil {
			fields := strings.Split(line, "=")
			if len(fields) == 2 {
				iniSubMap[fields[0]] = fields[1]
			}
		}
	}
	return iniMap
}
