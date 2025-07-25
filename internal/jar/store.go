package jar

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var storePath = filepath.Join(os.Getenv("HOME"), ".jar-ctl", "store.json")

func LoadStore() ([]JarEntry, error) {
	var jars []JarEntry

	// Ensure parent directory exists
	os.MkdirAll(filepath.Dir(storePath), os.ModePerm)

	// If file doesn't exist, create it with empty array
	if _, err := os.Stat(storePath); os.IsNotExist(err) {
		if err := os.WriteFile(storePath, []byte("[]"), 0644); err != nil {
			return nil, err
		}
	}

	// Read file
	file, err := os.Open(storePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&jars)
	return jars, err
}

func SaveStore(jars []JarEntry) error {
	os.MkdirAll(filepath.Dir(storePath), os.ModePerm)

	file, err := os.Create(storePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(jars)
}
