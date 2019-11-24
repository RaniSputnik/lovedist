package dist_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/RaniSputnik/lovedist/dist"
	"github.com/stretchr/testify/assert"
	"howett.net/plist"
)

const pathtolove = "../love/11.2.0/osx/love.app"

func buildDir() string {
	id := fmt.Sprintf("%d", time.Now().Unix())
	return filepath.Join("./build", fmt.Sprintf("build_%s", id))
}

func TestOSX(t *testing.T) {
	outputDir := buildDir()
	builder := dist.OSX(pathtolove, outputDir)
	p := dist.Project{
		Name:     "TestProject",
		BundleID: "com.example.TestProject",
	}
	input, err := os.Open("./testdata/helloworld.love")
	if err != nil {
		t.Fatalf("Failed to open test data: %v", err)
	}

	_, err = builder.Build(p, input)

	t.Run("Builds without error", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("Outputs to the correct location", func(t *testing.T) {
		expected := filepath.Join(outputDir, "TestProject.app")
		info, err := os.Stat(expected)
		if assert.NoError(t, err) {
			assert.True(t, info.IsDir())
		}
	})

	t.Run("Adds the correct name and bundle to the plist file", func(t *testing.T) {
		plistfile := filepath.Join(outputDir, "TestProject.app", "Contents", "Info.plist")
		file, err := os.Open(plistfile)
		if !assert.NoError(t, err) {
			return
		}
		defer file.Close()
		var got dist.Plist
		err = plist.NewDecoder(file).Decode(&got)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, p.Name, got.BundleName)
		assert.Equal(t, p.BundleID, got.BundleIdentifier)
	})
}
