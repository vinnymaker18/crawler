// Core package is the central module in the crawler program. All major
// datatypes and interfaces are defined here and core crawler function-
// ality is defined in terms of these concepts. Other packages like fe-
// tchers, storage etc... implement and adapt these interfaces for
// various article sources, storage provider types and so on.
// Entry point main simply reads configuration from a file, initializes
// various components and plugs in all required implementations and starts
// the crawler.
package core

import (
    "bufio"
    "os"
    "strings"
)

// Crawler program configuration, stored as key=value pairs, one 
// pair per line.
type Config struct {
    CfgKeyValues map[string]string
}

// Error in crawler configuration file.
type ConfigError struct {
	reason string
}

// Implement error interface for ConfigError
func (err *ConfigError) Error() string {
	return err.reason
}

// Config file parser.
func ParseConfig(configFilePath string) (*Config, *ConfigError) {
    f, err := os.Open(configFilePath)
    if err != nil {
        reason := "Error opening config file - " + err.Error()
        return nil, &ConfigError{ reason:reason }
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    keyValues := make(map[string]string)
    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Split(line, "=")
        keyValues[words[0]] = words[1]
    }

    if err := scanner.Err(); err != nil {
        reason := "Error scanning config file - " + err.Error()
        return nil, &ConfigError{ reason:reason}
    }

    return &Config {CfgKeyValues: keyValues}, nil
}
