package main

import (
	"github.com/TaitA2/bloggregator/internal/config"
	"github.com/TaitA2/bloggregator/internal/database"
)

type State struct {
	db     *database.Queries
	config *config.Config
}
