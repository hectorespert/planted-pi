package main

import (
	"fmt"
	"os"

	"github.com/reef-pi/reef-pi/controller/daemon"
	"github.com/reef-pi/reef-pi/controller/storage"
	"github.com/reef-pi/reef-pi/controller/utils"
)

func resetPassword(db, u, p string) {
	store, err := storage.NewStore(db)
	if u == "" {
		fmt.Println("username can not be empty")
		os.Exit(1)
	}
	if p == "" {
		fmt.Println("password can not be empty")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("Failed to open database. Check if reef-pi is already running. Error:", err)
		os.Exit(1)
	}
	defer store.Close()
	credentials := utils.Credentials{
		User:     u,
		Password: p,
	}
	if err := utils.NewCredentialsManager(store, daemon.Bucket).Update(credentials); err != nil {
		fmt.Println("Failed to save new credential. Error:", err)
		os.Exit(1)
	}
	fmt.Println("Credentials successfully updated")
}
