package main

import "github.com/google/uuid"

func generateSecurityToken() string {
	return uuid.New().String()
}
