package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

var PORT = ":5050"

type User struct {
	ID      int
	Name    string
	Email   string
	Address string
	Job     string
	MBTI    string
	Photo   string
}

var users = []User{
	{ID: 1, Name: "Minji", Email: "minji@newjeans.kr", Address: "Chuncheon, Gangwon, South Korea", Job: "Vocalist", MBTI: "ESTJ", Photo: "https://i.pinimg.com/564x/db/25/5a/db255abf12171f5bf18aba4337bfbe56.jpg"},
	{ID: 2, Name: "Hanni", Email: "hanni@newjeans.kr", Address: "Melbourne, Victoria, Australia", Job: "Main Vocal, Main Dance", MBTI: "INFP", Photo: "https://i.pinimg.com/564x/f3/24/fb/f324fbcb5d14d87588dde8f0c03fb068.jpg"},
	{ID: 3, Name: "Danielle", Email: "danille@newjeans.kr", Address: "Newcastle, New South Wales, Australia", Job: "Vocalist", MBTI: "ENFJ", Photo: "https://i.pinimg.com/564x/04/47/52/0447521f46cc96c0edcc8496685b0e5a.jpg"},
	{ID: 4, Name: "Haerin", Email: "haerin@newjeans.kr", Address: "Seoul, South Korea", Job: "Vocalist", MBTI: "INTP", Photo: "https://i.pinimg.com/564x/a3/84/8e/a3848ec6e3dc2cdc6ef90be66c08e10b.jpg"},
	{ID: 5, Name: "Hyein", Email: "hyein@newjeans.kr", Address: "Incheon, South Korea", Job: "Vocalist, maknae", MBTI: "ESFP", Photo: "https://i.pinimg.com/564x/af/24/cf/af24cf8737f87e02bca1fba549e862fa.jpg"},
}

var user_found = User{}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			loginPage(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/login", userCheck)
	http.HandleFunc("/profile", userProfil)
	http.HandleFunc("/logout", logout)

	fmt.Println("Server started at localhost", PORT)
	http.ListenAndServe(PORT, nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	var filePath = path.Join("views", "login.html")
	var tmpl, err = template.ParseFiles(filePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func userCheck(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		email := r.FormValue("email")
		for _, user := range users {
			if user.Email == email {
				user_found = user
				http.Redirect(w, r, "/profile", http.StatusFound)
				return
			}
		}
		var filePath = path.Join("views", "404.html")
		var tmpl, err = template.ParseFiles(filePath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user_found = User{Email: email}
		err = tmpl.Execute(w, user_found)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func userProfil(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var filePath = path.Join("views", "biodata.html")
		var tmpl, err = template.ParseFiles(filePath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, user_found)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user_found = User{}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
