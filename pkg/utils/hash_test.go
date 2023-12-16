package utils

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	hashedPassword, err := HashPassword(RandomString(8))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
}

func TestCheckPassword(t *testing.T) {
	password := RandomString(8)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)

	password = password[:4] + "s2332ew" + password[5:]

	err = CheckPassword(password, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
