package zep

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryManager_GetSession(t *testing.T) {
	// Define a session for the test
	sessionID := "testSession"
	testSession := Session{
		SessionID: sessionID,
	}

	// Set up the test server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(testSession)
	}))
	defer server.Close()

	memoryManager := NewMemoryManager(NewClient(server.URL, "", nil))

	// Call the GetSession method
	session, err := memoryManager.GetSession("testSession")

	// Assert there was no error
	assert.Nil(t, err)

	// Assert the returned session matches the test session
	assert.Equal(t, testSession.SessionID, session.SessionID)
}

func TestMemoryManager_AddSession(t *testing.T) {
	// TODO: Setup test server and MemoryManager
	// TODO: Generate test data

	// Call the AddSession method
	// session, err := memoryManager.AddSession(testSession)

	// Assert there was no error
	// assert.Nil(t, err)

	// Assert the returned session matches the test session
	// assert.Equal(t, testSession, session)
}

func TestMemoryManager_UpdateSession(t *testing.T) {
	// TODO: Setup test server and MemoryManager
	// TODO: Generate test data

	// Call the UpdateSession method
	// session, err := memoryManager.UpdateSession(testSession)

	// Assert there was no error
	// assert.Nil(t, err)

	// Assert the returned session matches the test session
	// assert.Equal(t, testSession, session)
}

func TestMemoryManager_AddMemory(t *testing.T) {
	// TODO: Setup test server and MemoryManager
	// TODO: Generate test data

	// Call the AddMemory method
	// memory, err := memoryManager.AddMemory("testSessionID", testMemory)

	// Assert there was no error
	// assert.Nil(t, err)

	// Assert the returned memory matches the test memory
	// assert.Equal(t, testMemory, memory)
}

func TestMemoryManager_GetMemory(t *testing.T) {
	// TODO: Setup test server and MemoryManager
	// TODO: Generate test data

	// Call the GetMemory method
	// memory, err := memoryManager.GetMemory("testSessionID", nil)

	// Assert there was no error
	// assert.Nil(t, err)

	// Assert the returned memory matches the test memory
	// assert.Equal(t, testMemory, memory)
}

func TestMemoryManager_SearchMemory(t *testing.T) {
	// TODO: Setup test server and MemoryManager
	// TODO: Generate test data

	// Call the SearchMemory method
	// results, err := memoryManager.SearchMemory("testSessionID", testSearchPayload, nil)

	// Assert there was no error
	// assert.Nil(t, err)

	// Assert the returned results matches the test results
	// assert.Equal(t, testResults, results)
}

func TestMemoryManager_ListSessions(t *testing.T) {
	// TODO: Setup test server and MemoryManager
	// TODO: Generate test data

	// Call the ListSessions method
	// sessions, err := memoryManager.ListSessions(nil, nil)

	// Assert there was no error
	// assert.Nil(t, err)

	// Assert the returned sessions matches the test sessions
	// assert.Equal(t, testSessions, sessions)
}

func TestMemoryManager_DeleteMemory(t *testing.T) {
	// TODO: Setup test server and MemoryManager
	// TODO: Generate test data

	// Call the DeleteMemory method
	// message, err := memoryManager.DeleteMemory("testSessionID")

	// Assert there was no error
	// assert.Nil(t, err)

	// Assert the returned message is correct
	// assert.Equal(t, "Memory deleted successfully", message)
}
