package server

import (
	"stack/web/backend/internal/database"
	"stack/web/backend/internal/store"
)

func Start() {
	store.SetDBConnection(database.NewDBOptions())
	router := setRouter()

	router.Run(":3000")
}