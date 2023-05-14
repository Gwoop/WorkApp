package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	PortHandler = ""
	Handler     = ""
	PathBD      = ""
	db          *sql.DB
)

func Init() {
	var filearray [6]string
	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		filearray[i] = scanner.Text()
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	PortHandler = filearray[0]
	Handler = filearray[1]
	PathBD = filearray[2]
}

type Testst struct {
	Resposnse string `json:"resposnse"`
}

func Test(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Testst{Resposnse: "Успешно"})
}

func main() {
	Init()
	fmt.Println("Запущенно")
	r := mux.NewRouter()
	r.HandleFunc("/marlo/test", Test).Methods("Get")
	r.HandleFunc("/marlo/admin/adduser", AuthorizationAdmin(AddUser)).Methods("Get")                             // добавление тестовых пользователей
	r.HandleFunc("/marlo/admin/getdocpattern", AuthorizationAdmin(Getdockspattern)).Methods("Get")               // получения списка шаблонов
	r.HandleFunc("/marlo/admin/adddocpattern", AuthorizationAdmin(Adddockpattern)).Methods("Post")               // создание шаблона
	r.HandleFunc("/marlo/admin/deletedocpattern/{id}/", AuthorizationAdmin(Deletedockpattern)).Methods("Delete") // удаление шаблона
	r.HandleFunc("/marlo/admin/searchdockspattern", AuthorizationAdmin(Searchdockspattern)).Methods("Get")       // поиск шаблонов
	r.HandleFunc("/marlo/admin/updatedockpattern/{id}/", AuthorizationAdmin(Updatedockpattern)).Methods("Put")   //обновление шаблона
	r.HandleFunc("/marlo/admin/getdockstext", AuthorizationAdmin(Getdockstext)).Methods("Get")
	r.HandleFunc("/marlo/admin/getdockstextbydocid/{id}/", AuthorizationAdmin(Getdockstextbydocid)).Methods("Get")
	r.HandleFunc("/marlo/admin/getdockstextbyid/{id}/", AuthorizationAdmin(Getdockstextbyid)).Methods("Get")
	r.HandleFunc("/marlo/admin/getdockstextactyality/{id}/", AuthorizationAdmin(GetDocksTextActyality)).Methods("Get")
	r.HandleFunc("/marlo/admin/adddockstextactyality/{id}/", AuthorizationAdmin(AddDocksTextActyality)).Methods("Post")
	r.HandleFunc("/marlo/admin/update_status_handler/{name_handler}/", AuthorizationAdmin(UpdateStatusHandler)).Methods("PUT")
	r.HandleFunc("/marlo/admin/insert_handler/", AuthorizationAdmin(InsertHandler)).Methods("Post")
	r.HandleFunc("/marlo/admin/delete_handler/{id_handler}", AuthorizationAdmin(DeleteHandler)).Methods("Delete")
	r.HandleFunc("/marlo/admin/get_handler", AuthorizationAdmin(GetHandler)).Methods("Get")
	r.HandleFunc("/marlo/admin/apply_changes", AuthorizationAdmin(Apply_Changes)).Methods("Post")
	r.Use(loggingMiddleware)
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	s := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Базовый лог, в дальнейшем буду делать более подробным
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func AutorizeihenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Базовый лог, в дальнейшем буду делать более подробным
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func Sqlconnectionmarlo(namebd string) {
	//"root:1234@tcp(localhost:3306)/admin"
	cfg := mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: namebd,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
