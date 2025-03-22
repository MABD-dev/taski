package services

import (
	"encoding/json"
	"os"
)

type LocalStorage[T any] struct {
	FileName string
}

func NewLocalStorage[T any](filename string) *LocalStorage[T] {
	return &LocalStorage[T]{FileName: filename}
}

func (s *LocalStorage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	os.WriteFile(s.FileName, fileData, 0644)
	return nil
}

func (s *LocalStorage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
