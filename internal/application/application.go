package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/pkg/calculation"
)

type Request struct {
	Expression string `json:"expression"`
}

//запуск сервера
func StartServer() {

	fmt.Println("start server port:8080")

	http.HandleFunc("/api/v1/calculate", RecoveryMiddleware(MidlewareLog(CalcHandler)))

	err := http.ListenAndServe(":8080", nil)
if err != nil{

}
}

func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Восстановление после паники: %v", err)
				http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

//записываем лог
func MidlewareLog(next http.HandlerFunc)http.HandlerFunc{

	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

		time := time.Now().Format(time.RFC3339)
		logger := log.New(log.Writer(),time, log.LstdFlags)
		logger.Output(2, "request received")

		next.ServeHTTP(w, r)
	})


}





type BadRequest struct {
	Error string `json:"error"`
}

type OkRequest struct {
	Result float64 `json:"result"`
}

// обрабатываем запрос
func CalcHandler(w http.ResponseWriter, r *http.Request) {

	request := new(Request)

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	

	res, errCalc := calculation.Calc(request.Expression)

	if errCalc != nil {
		br := BadRequest{Error: "Expression is not valid"}

		j, err := json.Marshal(br)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, string(j), http.StatusUnprocessableEntity)
		
	} else {

		okr := OkRequest{Result: res}
		jr, err := json.Marshal(okr)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // статус 200
		w.Write(jr)   

			}

}
