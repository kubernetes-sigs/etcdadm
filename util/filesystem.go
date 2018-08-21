package util

import (
	"os"
	"os/exec"
)

// Exists checks whether the file or directory exists.
func Exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CopyFile copies file from src to dest
func CopyFile(srcFile, destFile string) error {
	if err := exec.Command("cp", "-f", srcFile, destFile).Run(); err != nil {
		return err
	}
	return nil
}
