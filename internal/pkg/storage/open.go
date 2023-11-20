package storage

import (
	"github.com/pkg/errors"
	"os"
)

const remindersDBFileName = "reminders.json"

func OpenStorageFile() (string, error) {
	dataFile, err := getLocalFile(remindersDBFileName)
	if err != nil {
		return "", errors.Wrap(err, "OpenStorageFile")
	}

	return dataFile, nil
}

func getLocalFile(filename string) (string, error) {
	localConfigDir, err := ensureLocalConfigDir()
	if err != nil {
		return "", errors.Wrap(err, "checkFile - UserHomeDir")
	}

	localDataFile := localConfigDir + filename
	_, err = os.Stat(localDataFile)
	if os.IsNotExist(err) {
		createdFile, err := os.Create(localDataFile)
		if err != nil {
			return "", errors.Wrap(err, "checkFile - create")
		}

		_, err = createdFile.Write([]byte("[]"))
		if err != nil {
			return "", errors.Wrap(err, "checkFile - write")
		}

	}
	return localDataFile, nil
}

func ensureLocalConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := homeDir + "/.config/lembra/"
	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return configDir, nil
}
