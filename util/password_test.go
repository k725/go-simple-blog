package util

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	tests := []struct{
		Input string
		Result bool
	}{
		{
			Input:  "",
			Result: false,
		},
		{
			Input:  "test",
			Result: true,
		},
		{
			Input:  "日本語",
			Result: false,
		},
		{
			Input:  "123456789012345678901234567890123456789012345678901234567890123456789012",
			Result: true,
		},
		{
			Input:  "1234567890123456789012345678901234567890123456789012345678901234567890123",
			Result: false,
		},
	}
	for _, test := range tests {
		r := IsValidPassword(test.Input)
		assert.Equal(t, test.Result, r, test)
	}
}

func TestPasswordHash(t *testing.T) {
	tests := []struct{
		Input string
		ResultLen int
		Error error
	}{
		{
			Input:  "dolphin",
			ResultLen: 60,
			Error: nil,
		},
		{
			Input:  "password",
			ResultLen: 60,
			Error: nil,
		},
	}
	for _, test := range tests {
		r, err := PasswordHash(test.Input)
		assert.Len(t, r, test.ResultLen, test)
		if test.Error == nil {
			assert.Nil(t, err, test)
		} else {
			assert.EqualError(t, err, test.Error.Error(), test)
		}
	}
}

func TestPasswordVerify(t *testing.T) {
	tests := []struct{
		InputHash string
		InputPass string
		Result error
	}{
		{
			InputHash: "$2a$10$RIogJ.mUYubmyxGXVXsb8ugPiiwhu6hkDKv5z5VJPELDHyin6ZtWq",
			InputPass: "dolphin",
			Result: nil,
		},
		{
			InputHash: "$2a$10$RIogJ.mUYubmyxGXVXsb8ugPiiwhu6hkDKv5z5VJPELDHyin6ZtWq",
			InputPass: "foobar",
			Result: errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password"),
		},
		{
			InputHash: "invalid",
			InputPass: "",
			Result: errors.New("crypto/bcrypt: hashedSecret too short to be a bcrypted password"),
		},
	}
	for _, test := range tests {
		err := PasswordVerify(test.InputHash, test.InputPass)
		if test.Result == nil {
			assert.Nil(t, err, test)
		} else {
			assert.EqualError(t, err, test.Result.Error(), test)
		}
	}
}
