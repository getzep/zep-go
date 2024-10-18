package main

import "github.com/getzep/zep-go/v2"

var history = [][]zep.Message{
	// Conversation 1: Japan
	{
		{
			Role:     zep.String("Paul"),
			RoleType: "user",
			Content:  "Hi, I'm planning a vacation to Japan. Can you give me some advice?",
		},
		{
			Role:     zep.String("assistant"),
			RoleType: "assistant",
			Content:  "Of course! Japan is a fascinating destination. Are you more interested in modern cities like Tokyo, or historical sites like Kyoto?",
		},
		{
			Role:     zep.String("Paul"),
			RoleType: "user",
			Content:  "I think I'd like to experience both. Can you suggest an itinerary?",
		},
		{
			Role:     zep.String("assistant"),
			RoleType: "assistant",
			Content:  "Certainly! You could start with 3 days in Tokyo, then take the bullet train to Kyoto for 3 days. This way, you'll experience both the modern and traditional aspects of Japan.",
		},
		{
			Role:     zep.String("Paul"),
			RoleType: "user",
			Content:  "That sounds perfect! What are some must-visit places in Tokyo?",
		},
	},
	// Conversation 2: Italy
	{
		{
			Role:     zep.String("Paul"),
			RoleType: "user",
			Content:  "I'm thinking about visiting Italy next summer. Any recommendations?",
		},
		{
			Role:     zep.String("assistant"),
			RoleType: "assistant",
			Content:  "Italy is a wonderful choice! Are you more interested in art and history, or would you prefer to focus on food and wine experiences?",
		},
		{
			Role:     zep.String("Paul"),
			RoleType: "user",
			Content:  "I love both, but I think I'm leaning towards the food and wine experiences.",
		},
		{
			Role:     zep.String("assistant"),
			RoleType: "assistant",
			Content:  "Great! In that case, you might want to consider regions like Tuscany or Emilia-Romagna. Would you like more information about these areas?",
		},
		{
			Role:     zep.String("Paul"),
			RoleType: "user",
			Content:  "Yes, please tell me more about Tuscany. What are some must-try dishes and wines there?",
		},
	},
	{
		{
			Role:     zep.String("Paul"),
			RoleType: "user",
			Content:  "Apples are my favorite fruit",
		},
		{
			Role:     zep.String("Paul"),
			RoleType: "user",
			Content:  "now bananas are my favorite fruit",
		},
		{
			Role:     zep.String("Paul"),
			RoleType: "user",
			Content:  "Eric Clapton is my favorite guitarist",
		},
	},
	// Conversation 3: US Road Trip
}
