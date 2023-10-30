package zep

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserManager struct {
	Client Client
}

func NewUserManager(client Client) *UserManager {
	return &UserManager{Client: client}
}

func (u *UserManager) Add(user *CreateUserRequest) (*User, error) {
	request, err := http.NewRequest("POST", u.Client.GetFullURL("/user"), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := u.Client.HandleRequest(request, fmt.Sprintf("Failed to add user %s", user.UserID))

	if err != nil {
		return nil, err
	}

	var responseData User
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (u *UserManager) Get(userID string) (*User, error) {
	request, err := http.NewRequest("GET", u.Client.GetFullURL("/user/"+userID), nil)
	if err != nil {
		return nil, err
	}
	response, err := u.Client.HandleRequest(request, fmt.Sprintf("No user found for userID %s", userID))

	if err != nil {
		return nil, err
	}

	var responseData User
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (u *UserManager) Update(user *UpdateUserRequest) (*User, error) {
	request, err := http.NewRequest("PATCH", u.Client.GetFullURL("/user/"+user.UserID), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := u.Client.HandleRequest(request, fmt.Sprintf("Failed to update user %s", user.UserID))

	if err != nil {
		return nil, err
	}

	var responseData User
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (u *UserManager) Delete(userID string) error {
	request, err := http.NewRequest("DELETE", u.Client.GetFullURL("/user/"+userID), nil)
	if err != nil {
		return err
	}
	_, err = u.Client.HandleRequest(request, fmt.Sprintf("Failed to delete user %s", userID))

	if err != nil {
		return err
	}

	return nil
}

func (u *UserManager) List(limit *int, cursor *int) ([]User, error) {
	url := u.Client.GetFullURL("/user")
	if limit != nil {
		url += "?limit=" + strconv.Itoa(*limit)
	}
	if cursor != nil {
		url += "&cursor=" + strconv.Itoa(*cursor)
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := u.Client.HandleRequest(request, "Failed to list users")

	if err != nil {
		return nil, err
	}

	var responseData []User
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (u *UserManager) GetSessions(userID string) ([]Session, error) {
	request, err := http.NewRequest("GET", u.Client.GetFullURL("/user/"+userID+"/sessions"), nil)
	if err != nil {
		return nil, err
	}
	response, err := u.Client.HandleRequest(request, fmt.Sprintf("No sessions found for userID %s", userID))

	if err != nil {
		return nil, err
	}

	var responseData []Session
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (u *UserManager) ListChunked(chunkSize int) chan []User {
	ch := make(chan []User)

	go func() {
		defer close(ch)

		cursor := 0
		for {
			users, err := u.List(&chunkSize, &cursor)
			if err != nil || len(users) == 0 {
				break
			}

			ch <- users
			cursor += chunkSize
		}
	}()

	return ch
}
