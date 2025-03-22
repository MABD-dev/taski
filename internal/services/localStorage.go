package services

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type LocalStorage[T any] struct {
	FileName string
}

func NewLocalStorage[T any](filename string) *LocalStorage[T] {
	return &LocalStorage[T]{FileName: filename}
}

func (s *LocalStorage[T]) Save(data T) error {
	filePath, err := createFilePath(s.FileName)
	if err != nil {
		return err
	}

	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	os.WriteFile(filePath, fileData, 0644)
	return nil
}

func (s *LocalStorage[T]) Load(data *T) error {
	filePath, err := createFilePath(s.FileName)
	if err != nil {
		return err
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}

func createFilePath(filename string) (string, error) {
	taskiDir, err := getOrCreateRootDir()
	if err != nil {
		return "", nil
	}
	fileFullPath := filepath.Join(taskiDir, filename)
	return fileFullPath, nil
}

func getOrCreateRootDir() (string, error) {
	// TODO: make this configurable
	// path from root directory
	taskiPath := ".taski"

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}

	taskiDir := filepath.Join(homeDir, taskiPath)
	if _, err := os.Stat(taskiDir); os.IsNotExist(err) {
		err := os.MkdirAll(taskiDir, 0700)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	return taskiDir, nil
}
