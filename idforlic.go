package idforlic

import (
	"errors"
	"golang.org/x/sys/windows/registry"
	"runtime"
)

func GetID() (string, error) {
	if runtime.GOOS == "windows" {
		return winGetID()
	} else {
		return "", errors.New("Unknown OS")
	}
}

func winGetID() (wGuid string, wGuidErr error) {
	RegKey, RegKeyErr := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE|registry.WOW64_64KEY)
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
