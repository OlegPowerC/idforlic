package idforlic

import (
	"errors"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"os"
	"runtime"
)

const (
	linuxMfile    = `/var/lib/dbus/machine-id`
	linuxRHfile   = `/etc/machine-id`
	windowsRegKey = `SOFTWARE\Microsoft\Cryptography`
)

func GetID() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return winGetID()
		break
	case "linux":
		return linuxGetID()
		break
	}
	return "", errors.New("Unknown OS")
}

func winGetID() (wGuid string, wGuidErr error) {
	RegKey, RegKeyErr := registry.OpenKey(registry.LOCAL_MACHINE, windowsRegKey, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if RegKeyErr != nil {
		return "", RegKeyErr
	}
	defer RegKey.Close()
	KeyValue, _, KeyErr := RegKey.GetStringValue("MachineGuid")
	if KeyErr != nil {
		return "", KeyErr
	}
	return KeyValue, nil
}

func linuxGetID() (linuxId string, linuxIdErr error) {
	filecheck := linuxMfile
	_, fileExistError := os.Stat(filecheck)
	if os.IsNotExist(fileExistError) {
		filecheck = linuxRHfile
		_, fileExistError = os.Stat(filecheck)
		if os.IsNotExist(fileExistError) {
			return "", errors.New("Can not get ID on this OS")
		}
	}

	LinuxIDFile, LinuxIDFileErr := os.Open(filecheck)
	if LinuxIDFileErr != nil {
		return "", LinuxIDFileErr
	}
	defer LinuxIDFile.Close()
	ID, idReadErr := ioutil.ReadAll(LinuxIDFile)
	if idReadErr != nil {
		return "", idReadErr
	}
	return string(ID), nil
}
