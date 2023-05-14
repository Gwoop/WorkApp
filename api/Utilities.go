package main

import (
	"AdminSimpleApi/Structs"
	"github.com/ddosify/go-faker/faker"
	"github.com/satori/go.uuid"
	"math/rand"
)

func FakeData() Structs.User {
	userst := Structs.User{}
	faker := faker.NewFaker()                       //создание экземпляра объекта faker для подставных данных
	userst.Name = faker.RandomPersonFirstName()     //генерация подставного имяни
	userst.Lastname = faker.RandomPersonFirstName() //генерация подставной фамилии
	userst.Sex = rand.Intn(11-0) + 0                //генерация подставного пола (значение от 0 до 11 включительно)
	userst.Birdh = "2008/10/23"                     //Подставление данной даты рождения
	userst.Tel = faker.RandomPhoneNumber()          //генерация подставного номера телефона (логина)
	userst.Chatid = rand.Intn(11-0) + 0             //генерация подставного id чата (значение от 0 до 11 включительно)
	userst.Email = faker.RandomEmail()              //генерация подставной почты
	userst.Password = faker.RandomPassword()        //генерация подставного пороля (здесь он в чистом ввиде без хэширования)
	return userst
}

func Uuid() uuid.UUID {
	var err error
	u1 := uuid.Must(uuid.NewV4(), err)
	if err != nil {

	}

	return u1
}
