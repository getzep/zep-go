// This file was auto-generated by Fern from our API Definition.

package zep

type MemoryGetRequest struct {
	// memoryType: perpetual or message_window
	MemoryType *string `json:"-" url:"memoryType,omitempty"`
	// Last N messages. Overrides memory_window configuration
	Lastn *int `json:"-" url:"lastn,omitempty"`
}

type MemorySynthesizeQuestionRequest struct {
	// Last N messages
	LastNMessages *int `json:"-" url:"lastNMessages,omitempty"`
}