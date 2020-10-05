package cobi

import (
	"testing"
)

func TestInvalidCommandCompletion(t *testing.T) {
	scenarioTable := []scenario{
		{
			command:            "invalid",
			expectedSuggestion: []string{},
		},
		{
			command:            "test1invalid",
			expectedSuggestion: []string{},
		},
	}

	runScenarios(t, scenarioTable)
}

func TestFlagCompletion(t *testing.T) {
	scenarioTable := []scenario{
		{
			command:            "test2 --d",
			expectedSuggestion: []string{"--debug"},
		},
		{
			command:            "test2 -",
			expectedSuggestion: []string{"--debug", "-d", "--help", "-h"},
		},
	}

	runScenarios(t, scenarioTable)
}

func TestValidCommandCompletion(t *testing.T) {
	scenarioTable := []scenario{
		{
			command:            "tes",
			expectedSuggestion: []string{"test1", "test2"},
		},
	}

	runScenarios(t, scenarioTable)
}

func TestGeneratedSuggestions(t *testing.T) {

	scenarioTable := []scenario{
		{
			command:            "test1 S",
			expectedSuggestion: []string{"Suggestion1"},
		},
		{
			command:            "test2 dee",
			expectedSuggestion: []string{"deep"},
		},
		{
			command:            "ex",
			expectedSuggestion: []string{"exit"},
		},
	}

	runScenarios(t, scenarioTable)
}
