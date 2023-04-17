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

func ExportEnvVarsFromMap(env string) error {
	var consolidatedSecret ConsolidatedSecret
	err := json.Unmarshal([]byte(env), &consolidatedSecret)
	if err != nil {
		return err

	}
	for key, value := range consolidatedSecret.Current {
		fmt.Printf("%s=%v\n", key, value)

	}
	return nil
}

func GetEnvVar(envVarName string) (string, error) {
	envVarValue := os.Getenv(envVarName)
	return envVarValue, nil
}

func main() {

	envVar, err := GetEnvVar("SECRET_BUDDY")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = ExportEnvVarsFromMap(envVar)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
