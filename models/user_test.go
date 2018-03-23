package models

import (
	"testing"

	"github.com/joshuakwan/time-tracker/utils"
)

func TestCreateUser(t *testing.T) {
	users := createTestUsers()
	number := len(users)
	dup := false
	for i := 0; i < number; i++ {
		for j := i + 1; j < number; j++ {
			if users[i].ID == users[j].ID {
				dup = true
				break
			}
		}
	}

	utils.AssertNotEqual(t, dup, true, "UUID not unique")
}

func TestDeleteUsers(t *testing.T) {
	for _, user := range testInputUsers {
		users, err := GetUsersByEmail(user.email)
		utils.AssertEqual(t, len(users), 1, "user not found or duplicated")
		if err != nil {
			t.Fail()
		}
		users[0].Delete()

		users, err = GetUsersByEmail(user.email)
		utils.AssertEqual(t, len(users), 0, "user not deleted")
	}

}
