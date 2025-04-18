package main

import (
	"flag"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A simple test that generates a docker linter for testing
func TestGenerateALinter(t *testing.T) {
	dirPath, err := os.MkdirTemp("", "gabo-git-test")
	assert.NoError(t, err)
	t.Logf("Dir name is %s", dirPath)
	// Cleanup
	defer func() {
		require.NoError(t, os.RemoveAll(dirPath))
	}()

	setFlagOrFail(t, "dir", dirPath)
	setFlagOrFail(t, "mode", _modeAnalyze)
	err = flag.Set("mode", _modeAnalyze)
	assert.NoError(t, err)

	// Expected as this dir does not have ".git"
	assert.Error(t, validateGitDir())
	// Now make it a git dir
	err = os.Mkdir(filepath.Join(dirPath, ".git"), 0o755)
	assert.NoError(t, err)
	assert.NoError(t, validateGitDir())

	// Now, do analyze again, verify it won't fail
	main()
	// Now, add a basic generate test for docker linting
	setFlagOrFail(t, "mode", _modeGenerate)
	setFlagOrFail(t, "for", "lint-docker")
	main()
	assert.FileExists(t, filepath.Join(dirPath, ".github/workflows/lint-docker.yaml"))
}

func setFlagOrFail(t *testing.T, flagName string, value string) {
	err := flag.Set(flagName, value)
	assert.NoError(t, err)
}
