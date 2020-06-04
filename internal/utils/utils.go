package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// getWorkingPath gets the current working directory
func GetWorkingPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return wd, nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func MakeDirForFile(fileName string) error {
	dir := filepath.Dir(fileName)
	if err := MakeDir(dir); err != nil {
		return fmt.Errorf("unable to create dir for file: %s %w", fileName, err)
	}
	return nil
}

func MakeDir(dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("unable to create dir: %s %w", dir, err)
	}
	return nil
}
func MakeDirAndChDir(dir string) error {
	if err := MakeDir(dir); err != nil {
		return err
	}
	if err := os.Chdir(dir); err != nil {
		return err
	}
	return nil
}
func CreateFile(fileName string, data []byte) error {
	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		return fmt.Errorf("unable to create file: %s %w", fileName, err)
	}
	return nil
}
