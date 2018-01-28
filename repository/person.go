package repository

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/juanmanuelromeraferrio/rest-api/config/database"
	"github.com/juanmanuelromeraferrio/rest-api/model"
)

func GetPersonById(userId string) model.Person {
	var person model.Person
	if err := database.DB.First(&person, userId).Error; err != nil {
		panic(err)
	}
	database.DB.Model(person).Related(&person.Address)
	return person
}

func GetPeople() []model.Person {
	var people []model.Person
	if err := database.DB.Find(&people).Error; err != nil {
		panic(err)
	}

	for i, _ := range people {
		database.DB.Model(people[i]).Related(&people[i].Address)
	}

	return people
}

func CreatePerson(person model.Person) model.Person {
	if err := database.DB.Create(&person).Error; err != nil {
		panic(err)
	}
	database.DB.Model(person).Related(&person.Address)
	return person
}

func DeletePersonById(userId string) model.Person {
	person := GetPersonById(userId)
	if err := database.DB.Delete(&person).Error; err != nil {
		panic(err)
	}
	database.DB.Model(person).Related(&person.Address)
	return person
}
