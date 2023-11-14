package main

import (
	"flag"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {

	outPath := flag.String("outPath", hiveConfigs.UnsetStringValue, "Directory to write the generated installation resources")
	flag.Parse()

	// Rest of args should be HiveClusterConfigs
	hiveConfigFiles := flag.Args()

	hcBuilder, _ := hiveConfigs.NewHiveConfigBuilder()

	// For each HiveConfig
	for _, hiveConfigFile := range hiveConfigFiles {
		// Try to read it
		fileData, err := hiveConfigs.ReadFile(hiveConfigFile)
		if err != nil {
			log.Printf("Error: could not read %s: %s\n", hiveConfigFile, err)
		}

		// Try to yaml parse it
		hiveConfig := hiveConfigs.HiveConfig{}
		err = yaml.Unmarshal(fileData, &hiveConfig)
		if err != nil {
			log.Printf("Error: could not parse %s as yaml: %s\n", hiveConfigFile, err)
			continue
		}

		log.Printf("%s", fileData)

	}
}
