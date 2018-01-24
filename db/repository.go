package db

import (
	"go-api-scaffolding/models"
)

//ModelRepository ...
type ModelRepository struct{}

//Find ...
func (rep ModelRepository) Find(filter string, from int, size int) *[]models.Model {
	//your code here.
	return nil
}

//Get ...
func (rep ModelRepository) Get(id string) *models.Model {
	//your code here.
	return nil
}
