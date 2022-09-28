package gorm

import (
	"sesi8-latihan/server/models"
	"sesi8-latihan/server/repositories"

	"gorm.io/gorm"
)

type personRepo struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) repositories.PersonRepo {
	return &personRepo{
		db: db,
	}

}

func (p *personRepo) CreatePeople(person *models.Person) error {
	return p.db.Create(person).Error
}

func (p *personRepo) GetPeople() (*[]models.Person, error) {
	var person []models.Person
	err := p.db.Find(&person).Error
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (p *personRepo) GetPeopleById(id int) (*models.Person, error) {
	people := models.Person{}
	err := p.db.First(&people, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &people, nil
}

func (p *personRepo) DeletePeople(id int) error {
	people := models.Person{}
	err := p.db.Where("id = ?", id).Delete(&people).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *personRepo) UpdatePeople(Id int, person *models.Person) error {
	people := models.Person{}
	err := p.db.Model(&people).Where("id = ?", Id).Updates(models.Person{Name: person.Name, Address: person.Address}).Error
	if err != nil {
		return err
	}
	return nil
}
