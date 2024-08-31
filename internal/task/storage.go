package task

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func getStorageFilePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(cwd, "tasks.json"), nil
}

func createStorageFileIfNotExists(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if _, err := os.Create(path); err != nil {
			return err
		}

		if err := os.WriteFile(path, []byte("[]"), os.ModeAppend.Perm()); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

func LoadTasks() (*[]Task, error) {
	storageFilePath, err := getStorageFilePath()
	if err != nil {
		return nil, err
	}

	if err := createStorageFileIfNotExists(storageFilePath); err != nil {
		return nil, err
	}

	storageData, err := os.ReadFile(storageFilePath)
	if err != nil {
		return nil, err
	}
	if len(storageData) <= 0 {
		return nil, errors.New("storage file is empty")
	}

	var tasks []Task
	if err := json.Unmarshal(storageData, &tasks); err != nil {
		return nil, err
	}

	return &tasks, nil
}

func SaveTasks(tasks *[]Task) error {
	storageFilePath, err := getStorageFilePath()
	if err != nil {
		return err
	}

	if err := createStorageFileIfNotExists(storageFilePath); err != nil {
		return err
	}

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(storageFilePath, jsonData, 0755)
	if err != nil {
		return err
	}

	return nil
}
