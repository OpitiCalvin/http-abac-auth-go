package entity_test

import (
	"errors"
	"testing"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/entity"

	"github.com/stretchr/testify/assert"
)

// TestNewUser test creation of new user entity
func TestNewUser(t *testing.T) {
	u, err := entity.NewUser("testUser@mail.com", "newUser", "someSecretPass", int64(1))

	assert.Nil(t, err)
	assert.Equal(t, u.Email, "testUser@mail.com")
	assert.Equal(t, u.Username, "newUser")
	assert.NotEqual(t, u.Password, "someSecretPass")
	assert.Equal(t, u.ClientID, int64(1))
}

func TestUserValidate(t *testing.T) {
	type test struct {
		testName string
		email    string
		username string
		password string
		clientID int64
		want     error
	}

	tests := []test{
		{
			testName: "User Validate - No username",
			email:    "testUser@mail.com",
			username: "",
			password: "someSecretPass",
			clientID: int64(1),
			want:     errors.New("username is required"),
		},
		{
			testName: "User Validate - No email",
			email:    "",
			username: "newUser",
			password: "someSecretPass",
			clientID: int64(1),
			want:     errors.New("email address is required"),
		},
		{
			testName: "User Validate - No Password",
			email:    "testUser@mail.com",
			username: "newUser",
			password: "",
			clientID: int64(1),
			want:     errors.New("password is required"),
		},
		{
			testName: "User Validate - No client id",
			email:    "testUser@mail.com",
			username: "newUser",
			password: "someSecretPass",
			clientID: int64(0),
			want:     errors.New("a valid client id is required"),
		},
	}
	for _, tc := range tests {

		_, err := entity.NewUser(tc.email, tc.username, tc.password, tc.clientID)
		assert.Equal(t, err, tc.want)
	}

	// for _, tc := range tests {
	// 	t.Run(tc.testName, func(t *testing.T) {
	// 		_, err := entity.NewUser(tc.email, tc.username, tc.password, tc.clientID)
	// 		if err != tc.want {
	// 			t.Errorf("NewUser() error = %v, want %v", err, tc.want)
	// 		}
	// 	})
	// }

}

// TestUserPasswordCheck test compare hash and password
func TestUserPasswordCheck(t *testing.T) {
	u, _ := entity.NewUser("testUser@mail.com", "newUser", "someSecretPass", int64(1))

	assert.Equal(t, entity.CheckPassword(u.Password, "someSecretPass"), true)
}
