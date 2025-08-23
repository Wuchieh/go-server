package config

import "testing"

func TestGetMapStructure(t *testing.T) {
	result := GetMapStructure()
	for _, s := range result {
		t.Log(s)
	}
	t.Log()
	envResult := GetMapStructureForEnv(result)
	for _, s := range envResult {
		t.Log(s)
	}
}
