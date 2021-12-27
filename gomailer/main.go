package main

import (
	"log"
	"net/http"

	"net/smtp"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response":"Wooho it worked...."}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	from := "your_email@gmail.com"
	password := "enable 2FA and generate app password from gmail"
	to := []string{"reciver_email@gmail.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("This mail from go application")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		log.Fatal(err)
	} else {
		w.Write([]byte(`{"res": "Email send successfull"}`))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/send-email", post).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8081", r))

}
