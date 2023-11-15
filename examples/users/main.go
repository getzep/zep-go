package main

import (
	"fmt"

	"github.com/getzep/zep-go/zep"
	"github.com/google/uuid"
)

func main() {
	baseURL := "http://localhost:8000" // TODO: Replace with Zep API URL
	apiKey := ""                       // TODO: Replace with your API key

	client := zep.NewClient(baseURL, apiKey, nil)

	// Create multiple users
	for i := 0; i < 3; i++ {
		userID := fmt.Sprintf("user%d%s", i, uuid.New().String())
		userRequest := zep.CreateUserRequest{
			UserID:    userID,
			Email:     fmt.Sprintf("user%d@example.com", i),
			FirstName: fmt.Sprintf("John%d", i),
			LastName:  fmt.Sprintf("Doe%d", i),
			Metadata:  map[string]interface{}{"foo": "bar"},
		}
		user, err := client.User.Add(&userRequest)
		if err != nil {
			fmt.Printf("Failed to create user %d: %v\n", i+1, err)
		} else {
			fmt.Printf("Created user %d: %s\n", i+1, user.UserID)
		}
	}

	// Update the first user
	users, err := client.User.List(&[]int{1}[0], nil)
	if err != nil {
		fmt.Printf("Failed to list users: %v\n", err)
	} else {
		userID := users[0].UserID
		userRequest := zep.UpdateUserRequest{
			UserID:    userID,
			Email:     "updated_user@example.com",
			FirstName: "UpdatedJohn",
			LastName:  "UpdatedDoe",
			Metadata:  map[string]interface{}{"foo": "updated_bar"},
		}
		updatedUser, err := client.User.Update(&userRequest)
		if err != nil {
			fmt.Printf("Failed to update user: %v\n", err)
		} else {
			fmt.Printf("Updated user: %s\n", updatedUser.UserID)
		}
	}

	// Create a Session for the first user
	sessionID := uuid.New().String()
	session := zep.Session{
		SessionID: sessionID,
		UserID:    users[0].UserID,
		Metadata:  map[string]interface{}{"session": 1},
	}
	_, err = client.Memory.AddSession(&session)
	if err != nil {
		fmt.Printf("Failed to create session: %v\n", err)
	} else {
		fmt.Printf("Created session: %s\n", session.SessionID)
	}

	// Delete the second user
	users, err = client.User.List(&[]int{1}[0], &[]int{1}[0])
	if err != nil {
		fmt.Printf("Failed to list users: %v\n", err)
	} else {
		userID := users[0].UserID
		err = client.User.Delete(userID)
		if err != nil {
			fmt.Printf("Failed to delete user: %v\n", err)
		} else {
			fmt.Printf("Deleted user: %s\n", userID)
		}
	}

	// List all users
	users, err = client.User.List(nil, nil)
	if err != nil {
		fmt.Printf("Failed to list users: %v\n", err)
	} else {
		fmt.Println("All users:")
		for _, user := range users {
			fmt.Println(user.UserID)
		}
	}
}
