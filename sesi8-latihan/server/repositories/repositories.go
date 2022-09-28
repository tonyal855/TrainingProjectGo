package repositories

import "sesi8-latihan/server/models"

type PersonRepo interface {
	GetPeople() (*[]models.Person, error)
	GetPeopleById(id int) (*models.Person, error)
	CreatePeople(people *models.Person) error
	DeletePeople(Id int) error
	UpdatePeople(Id int, person *models.Person) error
}
