package requests

import "strings"

type UserDisplayNameEdit struct {
	DisplayName string
}

func (obj *UserDisplayNameEdit) GetTrimDisplayName() string {
	return strings.TrimSpace(obj.DisplayName)
}

type UserSensitiveOptionEdit struct {
	SensitiveDirect bool
}
