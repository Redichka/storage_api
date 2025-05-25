package main

import (
	"context"
	"fmt"
	"storage_api/pkg/db"
	repo "storage_api/repository"

	"github.com/google/uuid"
)

func main() {
	db.Connect()
	var ctx context.Context
	strepo := repo.NewStorageRepo(db.DB)
	id, _ := uuid.Parse("bc28ee88-5b35-4d98-bf3c-34a602e51361")
	model, _ := strepo.GetByRef(ctx, id)
	fmt.Println(model)
}
