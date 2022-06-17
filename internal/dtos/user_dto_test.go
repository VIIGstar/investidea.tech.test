package dtos

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/pkg/auth"
	"investidea.tech.test/pkg/password"
	"testing"
)

func TestUserDTO_IsValid_False(t *testing.T) {
	dto := UserDTO{
		User: entities.User{
			Role:     auth.BuyerRole.String(),
			Email:    "example",
			Name:     "Trung",
			Password: "123456",
			Address:  "ha noi, viet nam",
		},
	}
	errs := dto.Validate()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, EmailFormatErrReason, errs[0])
}

func TestUserDTO_IsValid_OK(t *testing.T) {
	dto := UserDTO{
		User: entities.User{
			Role:     auth.BuyerRole.String(),
			Email:    "investidea@gmail.com",
			Name:     "Trung",
			Password: "123456",
			Address:  "ha noi, viet nam",
		},
	}
	errs := dto.Validate()
	assert.Nil(t, errs)
}

func TestUserDTO_ToEntity(t *testing.T) {
	dto := UserDTO{
		User: entities.User{
			Role:     auth.BuyerRole.String(),
			Email:    "investidea@gmail.com",
			Name:     "Trung",
			Password: "123456",
			Address:  "ha noi, viet nam",
		},
	}
	fmt.Println(password.Hash(dto.Password))
	user, err := dto.ToEntity()
	assert.Nil(t, err)
	assert.True(t, password.Hash(dto.Password) == user.(entities.User).Password)
}
