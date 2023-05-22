package main

import (
	"WorkApp/api/security"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	zag   = 0
	n     = 1
	maxn  = 0
	count = 0
)

// хэндлел для авторизации
func AuthorizationAdmin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login, password, ok := r.BasicAuth() //инициализация базовой авторизации
		if !ok {
			(w).WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Ошибка обработки сессии")
			return
		}
		Sqlconnectionmarlo("admin")
		var err error
		if err != nil {
			panic(err)
		}
		defer db.Close()
		//проверка логина и пароля администратора из базы данных
		rows, err := db.Query("select * from admin.aunt")
		for rows.Next() {
			p := Admin{}
			erro := rows.Scan(&p.Id, &p.Login, &p.Password, &p.Token)
			if erro != nil {
				fmt.Println(erro)
				continue
			}
			//if login == p.Login && password == p.Password {
			//
			//}
			if login != p.Login || password != p.Password {
				(w).WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode("Ошибка ввода данных (логин или пароль не верны)")
				return
			}
		}
		(w).WriteHeader(http.StatusOK)
		handler.ServeHTTP(w, r)
	}
}

func Testingbd(w http.ResponseWriter, r *http.Request) {
	Sqlconnectionmarlo("marlo")

	(w).WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	procentflouat, _ := strconv.ParseFloat(vars["proc"], 8)
	println(vars["proc"])
	var s string = vars["sost"]
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	zag = i
	go bd(procentflouat)
}

func CloseTesting(w http.ResponseWriter, r *http.Request) {
	(w).WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	var s string = vars["sost"]
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	zag = i
}

//2023-01-26 09:35:29.544323 active_thread: 10 max_thread: 10 sended_request: 44
func logingtxt() {

	today := time.Now().String()
	//fmt.Println(today)
	dt := time.Now()

	t := strconv.Itoa(n)
	tmax := strconv.Itoa(maxn)
	tcount := strconv.Itoa(count)

	text := today + " active_thread: " + t + " max_thread: " + tmax + " sended_request: " + tcount + "\n"
	//file, err := os.Create("hello.txt")  D:\GGWP\WorkApp\logs
	//file.Close()

	var path string = "logs/" + dt.Format("01-02-2006") + ".txt"

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer f.Close()
	f.WriteString(text) //
}

func bd(procentflouat float64) {
	println(procentflouat)
	n = 2
	count = 0
	maxn = 0
	for zag == 1 {
		v, _ := mem.VirtualMemory()
		if v.UsedPercent == procentflouat {
			n--
			print("--")
		}
		if v.UsedPercent <= procentflouat {
			n++
			print("++")
		}
		if n == 0 {
			n++
			print("00")
		}
		if maxn < n {
			maxn = n
		}
		for i := 0; i < n; i++ {

			//println(v.UsedPercent)

			go test()

			//go theads(true)
		}
	}

}
func test() {

	println(n)
	logingtxt()
	//logingtxt()
}

func theads(zag bool) {

	count = 0
	for zag {
		//println("Горутины работают!!!!!!!")

		var err error
		defer db.Close()

		// Create a table
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS test_table (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(50) NOT NULL, PRIMARY KEY (id));")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Insert some data
		stmt, err := db.Prepare("INSERT INTO test_table (name) VALUES (?)")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer stmt.Close()

		for i := 0; i < 100; i++ {
			_, err = stmt.Exec("Name")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		rows, err := db.Query("SELECT * FROM test_table")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			err = rows.Scan(&id, &name)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(id, name)
		}
		count++
		logingtxt()
	}
}

func GetHardwareData(w http.ResponseWriter, r *http.Request) {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	// convert to JSON. String() is also implemented
	fmt.Println(v)

}

// хэндлер для добавления тестовых пользователей
func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userst := FakeData() //подставление случаных данных для создания тестового пользователя

	Sqlconnectionmarlo("marlo")
	var err error
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//запрос для бд на добавления тестового пользователя (пароль хешируется под MD5 с помощью функции GetMD5Hash)
	result, err := db.Exec("insert into marlo.users (name, lastname, sex,birdh,tel,chatid,email,password) values ( ?, ?, ?, ?, ?, ?, ?, ?)",
		userst.Name, userst.Lastname, userst.Sex, userst.Birdh, userst.Tel, userst.Chatid, userst.Email, security.GetMD5Hash(userst.Password))
	if err != nil {
		panic(err)
	}
	var id, _ = result.LastInsertId()
	json.NewEncoder(w).Encode(ResponsesUser{Id: id, Login: userst.Tel, Password: userst.Password}) // отправка ответа пользователю
	fmt.Println(result.LastInsertId())                                                             // id добавленного объекта
	fmt.Println(result.RowsAffected())                                                             // количество затронутых строк
}

// хэндлер для получения всех шаблонов
func Getdockspattern(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Sqlconnectionmarlo("marlo")
	rows, err := db.Query("SELECT * FROM document")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	docc := ResponsesDockpattern{}
	doc := []ResponsesDockpattern{}
	for rows.Next() {
		rows.Scan(&docc.Id, &docc.Name, &docc.Description, &docc.Uuid, &docc.Create_date)
		//json.NewEncoder(w).Encode(&doc)
		doc = append(doc, docc)
	}
	json.NewEncoder(w).Encode(&doc)
}

func AuthorizationReturn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	message := map[string]interface{}{
		"answer": true,
	}

	json.NewEncoder(w).Encode(message)
}

// хэндлер для добавления Шаблона документов
func Adddockpattern(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := ""
	var requestdockpattern RequestDockpattern
	_ = json.NewDecoder(r.Body).Decode(&requestdockpattern)

	fmt.Println(requestdockpattern.Name)
	fmt.Println(requestdockpattern.Description)
	Sqlconnectionmarlo("marlo")
	var err error
	defer db.Close()

	_, err = db.Query("insert into marlo.document (name, description, uuid) values (?,?,?)", requestdockpattern.Name, requestdockpattern.Description, Uuid())
	if err != nil {

		res = "Ошибка данных" + err.Error()
		ResponsesUserfunc(w, res)

		//response.Responses = "Ошибка данных " + err.Error()
		//json.NewEncoder(w).Encode(&response)
		return
	}
	res = "Данные успешно добавлены"
	ResponsesUserfunc(w, res)
	//response.Responses = "Данные успешно добавлены"
	//json.NewEncoder(w).Encode(&response)
}

// хэндлер для удаления Шаблона документов
func Deletedockpattern(w http.ResponseWriter, r *http.Request) {
	res := ""
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	Sqlconnectionmarlo("marlo")
	var err error
	defer db.Close()
	_, err = db.Query("DELETE FROM `marlo`.`document` WHERE (`id` = ?);", vars["id"])
	//fmt.Fprintf(w, "Category: %v\n", vars["category"])
	fmt.Println(vars["id"])
	if err != nil {
		res = "Ошибка данных" + err.Error()
		ResponsesUserfunc(w, res)
		return
	}
	res = "Данные успешно удалены или их не было"
	ResponsesUserfunc(w, res)
}

func Searchdockspattern(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Sqlconnectionmarlo("marlo")

	var requestsearchdock RequestsearchDock
	_ = json.NewDecoder(r.Body).Decode(&requestsearchdock)

	rows, err := db.Query("SELECT * FROM document WHERE name= ?", requestsearchdock.Namedoc)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		doc := ResponsesDockpattern{}
		rows.Scan(&doc.Id, &doc.Name, &doc.Description, &doc.Uuid, &doc.Create_date)
		json.NewEncoder(w).Encode(&doc)
	}
}

func Updatedockpattern(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//response := Structs.ResponsesSytem{}
	res := ""
	var requestdockpattern RequestDockpattern
	_ = json.NewDecoder(r.Body).Decode(&requestdockpattern)
	Sqlconnectionmarlo("marlo")
	var err error
	defer db.Close()

	_, err = db.Query("UPDATE document SET name = ?, description = ? WHERE (`id` = ?);", requestdockpattern.Name, requestdockpattern.Description, vars["id"])
	if err != nil {
		res = "Ошибка данных" + err.Error()
		ResponsesUserfunc(w, res)
		return
	}
	res = "Данные успешно Обновленны"
	ResponsesUserfunc(w, res)
}

func Getdockstext(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Sqlconnectionmarlo("marlo")
	rows, err := db.Query("SELECT * FROM document_text")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	docc := ResponsesDockstext{}
	doc := []ResponsesDockstext{}
	for rows.Next() {
		rows.Scan(&docc.Id, &docc.Id_doc, &docc.Text, &docc.Create_date, &docc.Lang, &docc.Uuid)
		//json.NewEncoder(w).Encode(&doc)
		doc = append(doc, docc)
	}
	json.NewEncoder(w).Encode(&doc)
}

func Getdockstextbydocid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Sqlconnectionmarlo("marlo")
	vars := mux.Vars(r)
	rows, err := db.Query("SELECT * FROM document_text WHERE id_doc= ?", vars["id"])
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		doc := ResponsesDockstext{}
		rows.Scan(&doc.Id, &doc.Id_doc, &doc.Text, &doc.Create_date, &doc.Lang, &doc.Uuid)
		json.NewEncoder(w).Encode(&doc)

	}
}

func Getdockstextbyid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Sqlconnectionmarlo("marlo")
	vars := mux.Vars(r)
	rows, err := db.Query("SELECT * FROM document_text WHERE id= ?", vars["id"])
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		doc := ResponsesDockstext{}
		rows.Scan(&doc.Id, &doc.Id_doc, &doc.Text, &doc.Create_date, &doc.Lang, &doc.Uuid)
		json.NewEncoder(w).Encode(&doc)
	}
}

func GetDocksTextActyality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Sqlconnectionmarlo("marlo")
	vars := mux.Vars(r)
	var err error
	rows, err := db.Query("SELECT * FROM document_text  WHERE id_doc= ? ORDER BY id DESC LIMIT 1", vars["id"])
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		doc := ResponsesDockstext{}
		rows.Scan(&doc.Id, &doc.Id_doc, &doc.Text, &doc.Create_date, &doc.Lang, &doc.Uuid)
		json.NewEncoder(w).Encode(&doc)
	}

}

func AddDocksTextActyality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Sqlconnectionmarlo("marlo")
	vars := mux.Vars(r)
	//res := ""
	var requestdockstext RequestDockstext
	_ = json.NewDecoder(r.Body).Decode(&requestdockstext)

	var err error
	rows, err := db.Exec("insert into document_text (id_doc, text, lang,uuid) values (?,?,?,?)", vars["id"], requestdockstext.Text, requestdockstext.Lang, Uuid())
	if err != nil {
		panic(err)

	}
	defer db.Close()
	var id, _ = rows.LastInsertId()
	res := "Данные успешно добавленны - id записи " + strconv.FormatInt(id, 10)
	ResponsesUserfunc(w, res)

}

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	Sqlconnectionmarlo("admin")
	vars := mux.Vars(r)
	//res := ""
	var requesthandler RequestHandler
	_ = json.NewDecoder(r.Body).Decode(&requesthandler)

	var err error
	rows, err := db.Exec("UPDATE handlers SET status = ? WHERE (name_handler = ?);", vars["name_handler"], requesthandler.Status)
	if err != nil {
		panic(err)

	}
	defer db.Close()
	var id, _ = rows.LastInsertId()
	res := "Данные хэндлера изменены - id изменённого хэнлера  " + strconv.FormatInt(id, 10)
	ResponsesUserfunc(w, res)
}

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	Sqlconnectionmarlo("admin")
	var requestinserthandler RequestInsertHandler
	_ = json.NewDecoder(r.Body).Decode(&requestinserthandler)

	var err error
	rows, err := db.Exec("insert into handlers (name_handler, status) values (?,?);", requestinserthandler.NameHandler, requestinserthandler.Status)
	if err != nil {
		panic(err)

	}
	defer db.Close()
	var id, _ = rows.LastInsertId()
	res := "Данные хэндлера добавленны - id созданного хэнлера  " + strconv.FormatInt(id, 10)
	ResponsesUserfunc(w, res)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	Sqlconnectionmarlo("admin")
	vars := mux.Vars(r)
	var err error
	rows, err := db.Exec("DELETE FROM handlers WHERE (`id` = ?);", vars["id_handler"])
	if err != nil {
		panic(err)

	}
	defer db.Close()
	var id, _ = rows.RowsAffected()

	res := "Данные хэндлера удаленны или их не было, задейсвованно строк - " + strconv.FormatInt(id, 10)
	ResponsesUserfunc(w, res)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Sqlconnectionmarlo("admin")
	var err error
	rows, err := db.Query("SELECT * FROM handlers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	gethandarr := []RequestGetHandlers{}
	gethand := RequestGetHandlers{}

	for rows.Next() {

		rows.Scan(&gethand.Id, &gethand.NameHandler, &gethand.Status)
		gethandarr = append(gethandarr, gethand)

	}
	json.NewEncoder(w).Encode(&gethandarr)

}

func Apply_Changes(w http.ResponseWriter, r *http.Request) {
	Sqlconnectionmarlo("admin")
	var err error
	rows, err := db.Query("SELECT * FROM handlers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	gethandarr := []RequestGetHandlers{}
	gethand := RequestGetHandlers{}

	for rows.Next() {

		rows.Scan(&gethand.Id, &gethand.NameHandler, &gethand.Status)
		gethandarr = append(gethandarr, gethand)

	}
	bytesRepresentation, err := json.Marshal(&gethandarr)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://пиздец_нахуй_блэт/post", "application/json", bytes.NewBuffer(bytesRepresentation)) //сюда путь до клиентской api
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	json.NewEncoder(w).Encode(&result)
}

func ResponsesUserfunc(w http.ResponseWriter, res string) {
	response := ResponsesSytem{}
	w.Header().Set("Content-Type", "application/json")
	response.Responses = res
	json.NewEncoder(w).Encode(&response)
}
