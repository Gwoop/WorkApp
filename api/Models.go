package Structs

//СТруктура для создания тестового пользователя
type User struct {
	Name     string //Имя
	Lastname string //Фамилия
	Sex      int    //Пол
	Birdh    string //Дата рождения
	Tel      string //номер телефона (логин)
	Chatid   int    //Внешний ключ от таблицы чата
	Email    string //Электронная почта
	Password string //Пароль
}

//Структура ответа для создания пользоватлей
type ResponsesUser struct {
	Id       int64  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

//Структура для авторизации админа
type Admin struct {
	Id       int
	Login    string
	Password string
}

//Структура ответа для просмотра всех шаблонов документов
type ResponsesDockpattern struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Uuid        string `json:"uuid"`
	Create_date string `json:"create_date"`
}

//Структура запроса для создания шаблона документа также эта структура используется для обновления
type RequestDockpattern struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//Структура запроса через Id
type RequestDockid struct {
	Id int `json:"id"`
}

type RequestsearchDock struct {
	Namedoc string `json:"namedoc"`
}

//Структура для вывода статуса операции, ошибках и т.д
type ResponsesSytem struct {
	Responses string `json:"responses"`
}

//Стуктура ответа для просмотра текстов документов
type ResponsesDockstext struct {
	Id          int64  `json:"id"`
	Id_doc      int64  `json:"id_doc"`
	Text        string `json:"text"`
	Create_date string `json:"create_date"`
	Lang        string `json:"lang"`
	Uuid        string `json:"uuid"`
}

//Стуктура ответа для просмотра текстов документов
type RequestDockstext struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
	Uuid string `json:"uuid"`
}

type RequestHandler struct {
	Status int `json:"status"`
}
type RequestInsertHandler struct {
	NameHandler string `json:"namehandler"`
	Status      int    `json:"status"`
}

type RequestDeleteHandler struct {
	Id string `json:"id"`
}

type RequestGetHandlers struct {
	Id          int64  `json:"id"`
	NameHandler string `json:"namehandler"`
	Status      int    `json:"status"`
}
