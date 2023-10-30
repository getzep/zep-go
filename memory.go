package zep

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type MemoryManager struct {
	Client Client
}

func NewMemoryManager(client Client) *MemoryManager {
	return &MemoryManager{Client: client}
}

func (m *MemoryManager) GetSession(sessionID string) (*Session, error) {
	request, err := http.NewRequest("GET", m.Client.GetFullURL("/sessions/"+sessionID), nil)
	if err != nil {
		return nil, err
	}
	response, err := m.Client.HandleRequest(request, fmt.Sprintf("No session found for session %s", sessionID))
	if err != nil {
		return nil, err
	}

	var responseData Session
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &responseData, nil
}

func (m *MemoryManager) AddSession(session *Session) (*Session, error) {
	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", m.Client.GetFullURL("/sessions"), bytes.NewBuffer(sessionJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := m.Client.HandleRequest(request, fmt.Sprintf("Failed to add session %s", session.SessionID))

	if err != nil {
		return nil, err
	}

	var responseData Session
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &responseData, nil
}

func (m *MemoryManager) UpdateSession(session *Session) (*Session, error) {
	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("PATCH", m.Client.GetFullURL("/sessions/"+session.SessionID), bytes.NewBuffer(sessionJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := m.Client.HandleRequest(request, fmt.Sprintf("Failed to update session %s", session.SessionID))

	if err != nil {
		return nil, err
	}

	var responseData Session
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &responseData, nil
}

func (m *MemoryManager) AddMemory(sessionID string, memory *Memory) (*Memory, error) {
	memoryJSON, err := json.Marshal(memory)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", m.Client.GetFullURL("/sessions/"+sessionID+"/memory"), bytes.NewBuffer(memoryJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := m.Client.HandleRequest(request, fmt.Sprintf("Failed to add memory for session %s", sessionID))

	if err != nil {
		return nil, err
	}

	var responseData Memory
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &responseData, nil
}

func (m *MemoryManager) GetMemory(sessionID string, lastn *int) (*Memory, error) {
	url := m.Client.GetFullURL("/sessions/" + sessionID + "/memory")
	if lastn != nil {
		url += "?lastn=" + strconv.Itoa(*lastn)
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := m.Client.HandleRequest(request, fmt.Sprintf("No memory found for session %s", sessionID))

	if err != nil {
		return nil, err
	}

	var responseData Memory
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &responseData, nil
}

func (m *MemoryManager) SearchMemory(sessionID string, searchPayload *MemorySearchPayload, limit *int) ([]MemorySearchResult, error) {
	searchPayloadJSON, err := json.Marshal(searchPayload)
	if err != nil {
		return nil, err
	}
	url := m.Client.GetFullURL("/sessions/" + sessionID + "/search")
	if limit != nil {
		url += "?limit=" + strconv.Itoa(*limit)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(searchPayloadJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := m.Client.HandleRequest(request, fmt.Sprintf("Failed to search memory for session %s", sessionID))

	if err != nil {
		return nil, err
	}

	var responseData []MemorySearchResult
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return responseData, nil
}

func (m *MemoryManager) ListSessions(limit *int, cursor *int) ([]Session, error) {
	url := m.Client.GetFullURL("/sessions")
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

	response, err := m.Client.HandleRequest(request, "Failed to get sessions")
	if err != nil {
		return nil, err
	}

	var responseData []Session
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return responseData, nil
}

func (m *MemoryManager) DeleteMemory(sessionID string) (string, error) {
	request, err := http.NewRequest("DELETE", m.Client.GetFullURL("/sessions/"+sessionID+"/memory"), nil)
	if err != nil {
		return "", err
	}
	_, err = m.Client.HandleRequest(request, fmt.Sprintf("Failed to delete memory for session %s", sessionID))
	if err != nil {
		return "", err
	}

	return "Memory deleted successfully", nil
}
