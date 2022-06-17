package dtos

import (
	"crypto/md5"
	"investidea.tech.test/internal/entities"
	"net/mail"
)

const (
	EmailFormatErrReason = "[Email]: invalid format."
	InvalidRoleErrReason = "[Role]: invalid type."
)

type UserDTO struct {
	entities.User
}

func (dto UserDTO) ToEntity() (interface{}, error) {
	user := dto.User
	hashed := md5.Sum([]byte(user.Password))
	user.Password = string(hashed[:])
	return user, nil
}
func (dto UserDTO) Validate() []string {
	var listErr []string
	if _, mailErr := mail.ParseAddress(dto.Email); mailErr != nil {
		listErr = append(listErr, EmailFormatErrReason)
	}

	if len(dto.Role) == 0 {
		listErr = append(listErr, InvalidRoleErrReason)
	}
	return listErr
}
