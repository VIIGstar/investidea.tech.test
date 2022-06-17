package dtos

import (
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/pkg/password"
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
	user.Password = password.Hash(user.Password)
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
