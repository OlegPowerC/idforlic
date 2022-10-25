package idforlic

import (
	"golang.org/x/sys/windows/registry"
)

const (
	windowsRegKey = `SOFTWARE\Microsoft\Cryptography`
)

func GetID() (string, error) {
	return winGetID()
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
