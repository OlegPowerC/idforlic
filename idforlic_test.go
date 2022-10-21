package idforlic

import "testing"

const YOUR_GUID = "1f9c66be-a8d7-4963-0000-3c9db5000000"

func TestGetID(t *testing.T) {
	Guid, GuidErr := GetID()
	if GuidErr != nil {
		t.Errorf("Err: %s", GuidErr)
	} else {
		if Guid != YOUR_GUID {
			t.Errorf("Got wrong guid: %s", Guid)
		}
	}
}
