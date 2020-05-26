package commands

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestRegister is a unit test to ensure that, when called, the commands are registered correctly.
// at the moment we do not pass a Run function as the pointers caused the tests to fail
func TestRegister(t *testing.T) {
	type test struct {
		input    []Command
		expected map[string]Command
	}

	var tests = map[string]test{
		"command name is passed correctly": {
			[]Command{
				{
					Name:        "ping",
					Usage:       "ping",
					Description: "see how long the bot takes to respond.",
					Category:    "General",
					NeedArgs:    false,
					Args:        map[string]bool{},
					OwnerOnly:   false,
					Enabled:     true,
				},
			},
			map[string]Command{
				"ping": {
					Name:        "ping",
					Usage:       "ping",
					Description: "see how long the bot takes to respond.",
					Category:    "General",
					NeedArgs:    false,
					Args:        map[string]bool{},
					OwnerOnly:   false,
					Enabled:     true,
				},
			},
		},
		"command is not added if disabled": {
			[]Command{
				{
					Name:        "ping",
					Usage:       "ping",
					Description: "see how long the bot takes to respond.",
					Category:    "General",
					NeedArgs:    false,
					Args:        map[string]bool{},
					OwnerOnly:   false,
					Enabled:     false,
				},
			},
			map[string]Command{},
		},
		"disabled command is not added to list": {
			[]Command{
				{
					Name:        "ping",
					Usage:       "ping",
					Description: "see how long the bot takes to respond.",
					Category:    "General",
					NeedArgs:    false,
					Args:        map[string]bool{},
					OwnerOnly:   false,
					Enabled:     true,
				},
				{
					Name:        "pingTwo",
					Usage:       "ping",
					Description: "see how long the bot takes to respond.",
					Category:    "General",
					NeedArgs:    false,
					Args:        map[string]bool{},
					OwnerOnly:   false,
					Enabled:     true,
				},
			},
			map[string]Command{
				"ping": {
					Name:        "ping",
					Usage:       "ping",
					Description: "see how long the bot takes to respond.",
					Category:    "General",
					NeedArgs:    false,
					Args:        map[string]bool{},
					OwnerOnly:   false,
					Enabled:     true,
				},
				"pingTwo": {
					Name:        "pingTwo",
					Usage:       "ping",
					Description: "see how long the bot takes to respond.",
					Category:    "General",
					NeedArgs:    false,
					Args:        map[string]bool{},
					OwnerOnly:   false,
					Enabled:     true,
				},
			},
		},
		"all passed commands are added": {
			[]Command{
				{
					Name:        "ping",
					Usage:       "ping",
					Description: "see how long the bot takes to respond.",
					Category:    "General",
					NeedArgs:    false,
					Args:        map[string]bool{},
					OwnerOnly:   false,
					Enabled:     false,
				},
				{
					Name:        "pingTwo",
					Usage:       "ping",
					Description: "see how long the bot takes to respond.",
					Category:    "General",
					NeedArgs:    false,
					Args:        map[string]bool{},
					OwnerOnly:   false,
					Enabled:     true,
				},
			},
			map[string]Command{
				"pingTwo": {
					Name:        "pingTwo",
					Usage:       "ping",
					Description: "see how long the bot takes to respond.",
					Category:    "General",
					NeedArgs:    false,
					Args:        map[string]bool{},
					OwnerOnly:   false,
					Enabled:     true,
				},
			},
		},
	}

	// Loop through and run each test case
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			Commands = map[string]Command{}
			for _, command := range test.input {
				command.Register()
			}
			if diff := cmp.Diff(Commands, test.expected); diff != "" {
				t.Error(fmt.Sprintf("Test Failed: %v mismatch (-want +got):\n%s", name, diff))
			}
		})
	}
}
