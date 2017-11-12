package main

import (
	"os"
	"testing"
)

var testExec = os.Args[0]

func TestMain(t *testing.T) {
	t.Run("it shows usage", func(tt *testing.T) {
		args := []string{"ergo", "-h"}
		os.Args = args

		main()
		// Output: USAGE
	})
}

// TestMissingCommand will cause the main function to return a non-zero exit code
// so we need to do some wrapping to make the test not fail.
// For more info check ou to uset: https://talks.golang.org/2014/testing.slide#23
func TestMissingCommand(t *testing.T) {
	t.Run("missing a command so it shows usage", func(tt *testing.T) {
		args := []string{"ergo"}
		os.Args = args

		result := command()
		if result != nil {
			t.Errorf("Expected result to not be nil")
		}
	})
}

func TestListCommand(t *testing.T) {
	t.Run("it is list command", func(tt *testing.T) {
		args := []string{"ergo", "list"}
		os.Args = args

		result := command()
		if result == nil {
			t.Errorf("Expected result to not be nil")
		}

		result()
	})
}

func TestListNamesCommand(t *testing.T) {
	t.Run("it is list-names command", func(tt *testing.T) {
		args := []string{"ergo", "list-names"}
		os.Args = args

		result := command()
		if result == nil {
			t.Errorf("Expected result to not be nil")
		}

		result()
	})
}

func TestSetupCommand(t *testing.T) {
	t.Run("it shows usage", func(tt *testing.T) {
		args := []string{"ergo", "setup"}
		os.Args = args

		result := command()
		if result != nil {
			t.Errorf("Expected result to be nil")
		}
	})

	t.Run("it is setup command", func(tt *testing.T) {
		args := []string{"ergo", "setup", "osx"}
		os.Args = args

		result := command()
		if result == nil {
			t.Errorf("Expected result not to be nil")
		}
	})
}

func TestUrlCommand(t *testing.T) {
	t.Run("it shows usage", func(tt *testing.T) {
		args := []string{"ergo", "url"}
		os.Args = args

		result := command()
		if result != nil {
			t.Errorf("Expected result to be nil")
		}
	})

	t.Run("it is url command", func(tt *testing.T) {
		args := []string{"ergo", "url", "foo"}
		os.Args = args

		result := command()
		if result == nil {
			t.Errorf("Expected result not to be nil")
		}

		result()
	})
}

func TestRunCommand(t *testing.T) {
	t.Run("when has wrong domain flag", func(tt *testing.T) {
		args := []string{"ergo", "run", "-domain=foo"}
		os.Args = args

		result := command()
		if result == nil {
			tt.Errorf("Expected result not to be nil")
		}

		result()
	})
}

// t.Run("when an config wrong", func(tt *testing.T) {
// 	args := []string{"ergo", "-p foo", "run"}
// 	os.Args = args

// 	result := command()
// 	if result != nil {
// 		tt.Errorf("Expected result not to be nil")
// 	}
// })
// }

func TestAddCommand(t *testing.T) {
	t.Run("it shows usage", func(tt *testing.T) {
		args := []string{"ergo", "add"}
		os.Args = args

		result := command()
		if result != nil {
			t.Errorf("Expected result to be nil")
		}
	})

	t.Run("it is url command", func(tt *testing.T) {
		args := []string{"ergo", "add", "foo", "bar"}
		os.Args = args

		result := command()
		if result == nil {
			t.Errorf("Expected result not to be nil")
		}

		result()
	})
}

func TestRemoveCommand(t *testing.T) {
	t.Run("it returns no command", func(tt *testing.T) {
		args := []string{"ergo", "remove"}
		os.Args = args

		result := command()
		if result != nil {
			t.Errorf("Expected result to be nil")
		}
	})

	t.Run("it returns a command", func(tt *testing.T) {
		args := []string{"ergo", "remove", "test"}
		os.Args = args

		result := command()
		if result == nil {
			t.Errorf("Expected result to be a command, instead got nil")
		}
	})
}
