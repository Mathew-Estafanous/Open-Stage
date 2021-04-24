package main

import (
	"os"
	"strings"
)

var profile = os.Getenv("PROFILE")

// PortByProfile will return the desired port to bind to depending on
// the profile that the server was started with. If
func PortByProfile() string {
	if isProfileProd() {
		return ":" + os.Getenv("PORT")
	}
	return ":8080"
}

func DbUrlByProfile() string {
	dbUrl := os.Getenv("DATABASE_URL")
	if isProfileProd() {
		return dbUrl
	}
	return dbUrl + "?sslmode=disabled"
}

func isProfileProd() bool {
	return strings.ToLower(profile) == "prod"
}