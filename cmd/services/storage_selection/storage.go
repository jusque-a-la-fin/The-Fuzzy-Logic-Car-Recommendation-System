package main

import (
	"car-recommendation-service/cmd/services/storage"
)

func main() {
	configName := "storage_selection"
	storage.MakeStorage(configName)
}
