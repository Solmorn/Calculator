package application

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/Solmorn/Calculator/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFormEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFormEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type Result struct {
	Result string `json:"result"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode((&request))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fastr, err := calculation.Calc(request.Expression)
	if err == nil {
		result := Result{
			Result: strconv.FormatFloat(fastr, 'f', 6, 64),
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {

			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)

		}

	} else {
		if err == calculation.ErrDivisionByZero {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else if err == calculation.ErrInvalidExpression {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
