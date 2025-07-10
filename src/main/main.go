package main

import (
	"fmt"
	"html/template"
	"math/rand/v2"
	"net/http"
	"strings"
)

type Score struct {
	StateWINER int
	StateLOSSER  int
    StateDROW int
}

type User struct{
    Name string
    Password string
}

const ROCK = "rock"
const PAPER = "paper"
const SCISSORS = "scissors"
const WINER = "winer"
const LOSSER = "losser"
const DROW = "drow"

var stateWINER int = 0
var stateLOSSER int = 0
var stateDROW int = 0
var name string = "bot"

func main() {
	// Регистрируем обработчик для всех запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			data := ""
			tmpl, _ := template.ParseFiles("templates/root.html")
			tmpl.Execute(w, data)
		} else if r.Method == http.MethodPost {
			http.Redirect(w, r, "/game", http.StatusMovedPermanently)
		}

	})

	//TODO: Вставка   Post  Оотправка rock, paper, scissors
	// /game/<rock/paper/scissors>
	http.HandleFunc("/game", addChoice)

	//TODO: Удаление  Delete Отчиска счета
	http.HandleFunc("/clear", clearCost)

	//TODO: Изменение Put  Изменить имя
	http.HandleFunc("/rename", rename)

	//TODO: Выбору продукта GET Выбрать

	// Запускаем сервер на порту 8080
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
    
}

func addChoice(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("chose")

	//parametrsOfRequest := r.URL.RawQuery
	switch name {
	case ROCK:
		{
			t1, t2 := gameOne(ROCK)
			scoreOfGame(t2)
			fmt.Fprintf(w, "You = "+ROCK+" AI = "+t1+" State = "+t2)
			fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))
		}
	case PAPER:
		{
			t1, t2 := gameOne(PAPER)
			scoreOfGame(t2)
			fmt.Fprintf(w, "You = "+PAPER+" AI = "+t1+" State = "+t2)
			fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))
		}
	case SCISSORS:
		{
			t1, t2 := gameOne(SCISSORS)
			scoreOfGame(t2)
			fmt.Fprintf(w, "You = "+SCISSORS+" AI = "+t1+" State = "+t2)
			fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))
		}
	default:
        data := Score{
            StateWINER:stateWINER,
	        StateLOSSER:stateLOSSER,
            StateDROW:stateDROW,
        }
        tmpl, _ := template.ParseFiles("templates/game.html")
        tmpl.Execute(w, data)

		//fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))
	}

	// if strings.Contains(parametrsOfRequest, "chose=rock") {
	// 	t1, t2 := gameOne(ROCK)
	// 	scoreOfGame(t2)
	// 	fmt.Fprintf(w, "You = "+ROCK+" AI = "+t1+" State = "+t2)
	// 	fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))
	// } else if strings.Contains(parametrsOfRequest, "chose=paper") {
	// 	t1, t2 := gameOne(PAPER)
	// 	scoreOfGame(t2)
	// 	fmt.Fprintf(w, "You = "+PAPER+" AI = "+t1+" State = "+t2)
	// 	fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))
	// } else if strings.Contains(parametrsOfRequest, "chose=scissors") {
	// 	t1, t2 := gameOne(SCISSORS)
	// 	scoreOfGame(t2)
	// 	fmt.Fprintf(w, "You = "+SCISSORS+" AI = "+t1+" State = "+t2)
	// 	fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))
	// } else {
	// 	fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))
	// }
}

func clearCost(w http.ResponseWriter, r *http.Request) {
	stateWINER = 0
	stateLOSSER = 0
	stateDROW = 0
}

func rename(w http.ResponseWriter, r *http.Request) {
	parametrsOfRequest := r.URL.RawQuery
	strings.TrimPrefix(parametrsOfRequest, "name=")

}

func scoreOfGame(state string) {
	if state == WINER {
		stateWINER++
	}
	if state == LOSSER {
		stateLOSSER++
	}
	if state == DROW {
		stateDROW++
	}
}

func gameOne(chosce string) (aIChoice string, state string) {

	generator := rand.Int32N(3)
	switch generator {
	case 0:
		aIChoice = ROCK
	case 1:
		aIChoice = PAPER
	case 2:
		aIChoice = SCISSORS
	}

	if chosce == aIChoice {
		return aIChoice, DROW
	} else if chosce == ROCK {
		if aIChoice == PAPER {
			return aIChoice, LOSSER
		} else if aIChoice == SCISSORS {
			return aIChoice, WINER
		}
	} else if chosce == PAPER {
		if aIChoice == ROCK {
			return aIChoice, LOSSER
		} else if aIChoice == SCISSORS {
			return aIChoice, WINER
		}
	} else if chosce == SCISSORS {
		if aIChoice == PAPER {
			return aIChoice, WINER
		} else if aIChoice == ROCK {
			return aIChoice, LOSSER
		}
	}
	return aIChoice, WINER
}
