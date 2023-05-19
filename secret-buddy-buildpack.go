package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConsolidatedSecret struct {
	Current  map[string]string `json:"current"`  // Current secret values
	Previous map[string]string `json:"previous"` // Previous secret values, useful for rollbacks
}

func ExportEnvVarsFromMap(env string, rules map[string]string) (map[string]string, error) {
	var consolidatedSecret ConsolidatedSecret
	err := json.Unmarshal([]byte(env), &consolidatedSecret)
	var envVars map[string]string
	envVars = consolidatedSecret.Current
	fmt.Println(envVars)

	for key, value := range rules {
		var previousValue = "previous." + key
		if value == previousValue {
			envVars[key] = consolidatedSecret.Previous[key]
		}
	}
	if err != nil {
		return nil, err
	}
	return envVars, nil
}

func GetEnvVar(envVarName string) (string, error) {
	envVarValue := os.Getenv(envVarName)
	return envVarValue, nil
}

func ParseRules(ruleString string) (map[string]string, error) {
	var rules map[string]string
	err := json.Unmarshal([]byte(ruleString), &rules)
	if err != nil {
		return nil, err
	}
	return rules, err
}

func main() {

	envVar, err := GetEnvVar("SECRETBUDDY_ENV")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	variableX, err := GetEnvVar("HEROKU_SECRETS_CONFIG")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rules, err := ParseRules(variableX)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	consolidatedSecret, err := ExportEnvVarsFromMap(envVar, rules)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for key, value := range consolidatedSecret {
		fmt.Printf("export %s=%v\n", key, value)
	}

}
