package config

import (
	"github.com/Jeffail/gabs/v2"
	"os"
	"testing"
)

func saveGabsObj(t *testing.T, name string, g *gabs.Container, value string, hierarchy ...string) error {
	t.Helper()

	file, errCreate := os.Create(name)
	if errCreate != nil {
		t.Fatalf("Error opening file: %v", errCreate)
	}
	defer file.Close()

	if _, err := g.Set(value, hierarchy...); err != nil {
		t.Fatalf("Error creating \"%s\" in json: %v", hierarchy, err)
		return err
	}
	data, errMarshal := g.MarshalJSON()
	if errMarshal != nil {
		t.Fatalf("Error marshalling gabs container: %v", errMarshal)
	}

	if _, err := file.Write(data); err != nil {
		t.Fatalf("Error writing to \"config_test.json\": %v", err)
	}
	return nil
}

func clearFile(t *testing.T, name string) {
	t.Helper()

	errRemove := os.Remove(name)
	if errRemove != nil {
		t.Logf("Error deleting file: %v", errRemove)
	}
}

func TestNew(t *testing.T) {
	fileName := "config_test.json"
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("Error creating \"config_test.json\": %v", err)
	}
	if err = file.Close(); err != nil {
		t.Fatalf("Error closing file: %v", err)
	}

	g := gabs.New()

	tests := []struct {
		name    string
		prepare func() error
		wantErr bool
	}{
		{
			"test-incomplete-json",
			func() error {
				return saveGabsObj(t, fileName, g, "host=localhost user=postgres password=password dbname=tasks port=5432", "database", "dsn")
			},
			true,
		},
		{
			"test-new-works",
			func() error {
				if err = saveGabsObj(t, fileName, g, "host=localhost user=postgres password=password dbname=tasks port=5432", "database", "dsn"); err != nil {
					return err
				}
				if err = saveGabsObj(t, fileName, g, "dbkeyshouldbesecure@!.", "database", "db_key"); err != nil {
					return err
				}
				if err = saveGabsObj(t, fileName, g, ":8080", "api", "host"); err != nil {
					return err
				}
				return nil
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err = tt.prepare(); err != nil {
				t.Fatalf("Error preparing for test: %v", err)
			}

			cfg, errNew := New("config_test.json")
			hasError := errNew != nil
			if hasError != tt.wantErr {
				t.Fatalf("Error test result: got: %v, expected: %v", hasError, tt.wantErr)
			}

			if err != nil {
				t.Log(errNew)
			}

			if cfg != nil {
				t.Log(cfg)
			}

			clearFile(t, fileName)
		})
	}
}
