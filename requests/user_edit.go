package requests

import "strings"

type UserDisplayNameEdit struct {
	DisplayName string `json:"display_name"`
}

func (obj *UserDisplayNameEdit) GetTrimDisplayName() string {
	return strings.TrimSpace(obj.DisplayName)
}
