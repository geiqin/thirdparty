package utils

import (
	"github.com/satori/go.uuid"
	"strings"
)

func GetUUID() string {
	u := uuid.NewV4()
	return strings.Replace(u.String(), "-", "", -1)
}
