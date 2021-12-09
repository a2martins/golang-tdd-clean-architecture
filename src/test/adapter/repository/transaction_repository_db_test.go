package repository

import (
	"github.com/a2martins/golang-tdd-clean-architecture/src/main/adapter"
	"github.com/a2martins/golang-tdd-clean-architecture/src/test/adapter/repository/fixture"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTransactionRepositoryDB_Insert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := adapter.NewTransactionRepositoryDB(db)
	err := repository.Insert("1", "1", 2, "approved", "")
	assert.Nil(t, err)
}
