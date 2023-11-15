package zep

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserManager_Add(t *testing.T) {
	email := "test@example.com"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(User{
			UserID: "testUser",
			Email:  email,
		})
	}))
	defer server.Close()

	client := &http.Client{}
	zepClient := NewClient(server.URL, "", client)
	userManager := NewUserManager(zepClient)

	user := &CreateUserRequest{
		UserID: "testUser",
		Email:  email,
	}

	createdUser, err := userManager.Add(user)
	assert.Nil(t, err)
	assert.Equal(t, "testUser", createdUser.UserID)
	assert.Equal(t, email, createdUser.Email)
}

func TestUserManager_Add_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := &http.Client{}
	zepClient := NewClient(server.URL, "", client)
	userManager := NewUserManager(zepClient)

	user := &CreateUserRequest{
		UserID: "testUser",
	}

	_, err := userManager.Add(user)
	assert.NotNil(t, err)
	assert.IsType(t, &APIError{}, err)
}

func TestUserManager_Get(t *testing.T) {
	// Define a user for the test
	email := "test@example.com"
	testUser := User{
		UserID: "testUser",
		Email:  email,
	}

	// Set up the test server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(testUser)
	}))
	defer server.Close()

	// Set up the user manager
	userManager := NewUserManager(NewClient(server.URL, "", nil))

	// Call the Get method
	user, err := userManager.Get("testUser")

	// Assert there was no error
	assert.Nil(t, err)

	// Assert the returned user matches the test user
	assert.Equal(t, testUser.UserID, user.UserID)
	assert.Equal(t, testUser.Email, user.Email)
}

func TestUserManager_Get_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := &http.Client{}
	zepClient := NewClient(server.URL, "", client)
	userManager := NewUserManager(zepClient)

	_, err := userManager.Get("testUser")
	assert.NotNil(t, err)
	assert.IsType(t, &APIError{}, err)
}

func TestUserManager_Update(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		user := User{
			UserID: "testUser",
			Email:  "",
		}
		user.Email = "updated@example.com"
		json.NewEncoder(rw).Encode(user)
	}))
	defer server.Close()

	userManager := NewUserManager(NewClient(server.URL, "", nil))
	email := "updated@example.com"
	user := &UpdateUserRequest{
		UserID: "testUser",
		Email:  email,
	}

	updatedUser, err := userManager.Update(user)
	assert.Nil(t, err)
	assert.Equal(t, "testUser", updatedUser.UserID)
	assert.Equal(t, email, updatedUser.Email)
}

func TestUserManager_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	userManager := NewUserManager(NewClient(server.URL, "", nil))

	err := userManager.Delete("testUser")
	assert.Nil(t, err)
}

func TestUserManager_List(t *testing.T) {
	// Define a user for the test
	userID := "testUser"
	email := "test@example.com"
	testUser := User{
		UserID: userID,
		Email:  email,
	}

	// Set up the test server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode([]User{testUser})
	}))
	defer server.Close()

	userManager := NewUserManager(NewClient(server.URL, "", nil))

	users, err := userManager.List(nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, userID, users[0].UserID)
	assert.Equal(t, email, users[0].Email)
}

func TestUserManager_GetSessions(t *testing.T) {
	// Define a session for the test
	sessionID := "testSession"
	testSession := Session{
		SessionID: sessionID,
	}

	// Set up the test server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode([]Session{testSession})
	}))
	defer server.Close()

	userManager := NewUserManager(NewClient(server.URL, "", nil))

	sessions, err := userManager.GetSessions("testUser")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(sessions))
	assert.Equal(t, sessionID, sessions[0].SessionID)
}

func TestUserManager_ListChunked(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode([]User{})
	}))
	defer server.Close()

	userManager := NewUserManager(NewClient(server.URL, "", nil))

	ch := userManager.ListChunked(1)
	_, ok := <-ch
	assert.False(t, ok)
}
