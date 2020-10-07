package cobi

import (
	"testing"
)

func TestInvalidCommandCompletion(t *testing.T) {
	scenarioTable := []scenario{
		{
			command:            "invalid",
			expectedSuggestion: nil,
		},
		{
			command:            "test1invalid",
			expectedSuggestion: nil,
		},
	}

	runScenarios(t, scenarioTable)
}

func TestFlagCompletion(t *testing.T) {
	scenarioTable := []scenario{
		{
			command:            "test2 --d",
			expectedSuggestion: []string{"test2 --debug"},
		},
		{
			command:            "test2 -",
			expectedSuggestion: []string{"test2 --debug", "test2 -d", "test2 --help", "test2 -h"},
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
			expectedSuggestion: []string{"test1 Suggestion1", "test1 DiffSuggestion"},
		},
		{
			command:            "test2 dee",
			expectedSuggestion: []string{"test2 deep"},
		},
	}

	runScenarios(t, scenarioTable)
}
