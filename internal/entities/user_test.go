package entities

import (
	"github.com/stretchr/testify/require"
	"github.com/teten-nugraha/fiber-form-clean-architecture/internal/pkg/id"
	"testing"
	"time"
)

func TestEntitiesUser_IsValid(t *testing.T) {
	user := User{}
	user.ID = id.NewID()
	user.Username = "username"
	user.Password = "password"
	user.CreatedAt = time.Now()

	isValid, err := user.IsValid()
	require.True(t, isValid, err)
	require.Nil(t, err, err)
}

func TestEntityUser_NewUser(t *testing.T) {
	user, err := NewUser("username", "123456")
	require.Nil(t, err, err)
	require.NotEmpty(t, user.CreatedAt, err)
	require.NotEmpty(t, user.ID, err)

	err = user.ValidatePassword("123456")
	require.Nil(t, err, err)

	err = user.ValidatePassword("wrong_password")
	require.Error(t, err, err)
}

func TestEntityUser_ValidatePassword(t *testing.T) {
	u, _ := NewUser("username", "123456")
	err := u.ValidatePassword("123456")
	require.Nil(t, err, err)
	err = u.ValidatePassword("wrong_password")
	require.NotNil(t, err, err)
}
