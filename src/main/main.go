package main

import (
	"fmt"
	"html/template"
	"math/rand/v2"
	"net/http"
)

type Score struct {
	StateWINER  int
	StateLOSSER int
	StateDROW   int
}

type User struct {
	UserName string
	Password string
}

const ROCK = "rock"
const PAPER = "paper"
const SCISSORS = "scissors"
const WINER = "winer"
const LOSSER = "losser"
const DROW = "draw"

var stateWINER int = 0
var stateLOSSER int = 0
var stateDRAW int = 0
var userName string = "bot"

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

	//Вставка   Post  Оотправка rock, paper, scissors
	// /game/<rock/paper/scissors>
	http.HandleFunc("/game", addChoice)

	//Удаление  Delete Отчиска счета
	http.HandleFunc("/clear", clearCost)

	//Изменение Put  Изменить имя
	http.HandleFunc("/rename", rename)

	//Выбору продукта GET Выбрать

	// Запускаем сервер на порту 8080
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}

func addChoice(w http.ResponseWriter, r *http.Request) {
	choseOf := r.URL.Query().Get("chose")
	user := User{
		UserName: userName,
		Password: "123",
	}

	//parametrsOfRequest := r.URL.RawQuery
	switch choseOf {
	case ROCK:
		{
			_, t2 := gameOne(ROCK)
			scoreOfGame(t2)
			// fmt.Fprintf(w, "You = "+ROCK+" AI = "+t1+" State = "+t2)
			// fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))

			tmplSettings, _ := template.ParseFiles("templates/gamePage/settings.html")
			tmplSettings.Execute(w, user)
			data := Score{
				StateWINER:  stateWINER,
				StateLOSSER: stateLOSSER,
				StateDROW:   stateDRAW,
			}
			tmplStore, _ := template.ParseFiles("templates/gamePage/store.html")
			tmplStore.Execute(w, data)

			stateOf := struct {
				State string
			}{
				State: t2,
			}
			tmplState, _ := template.ParseFiles("templates/gamePage/state.html")
			tmplState.Execute(w, stateOf)

			tmplGame, _ := template.ParseFiles("templates/gamePage/game.html")
			tmplGame.Execute(w, data)
		}
	case PAPER:
		{
			_, t2 := gameOne(PAPER)
			scoreOfGame(t2)
			// fmt.Fprintf(w, "You = "+PAPER+" AI = "+t1+" State = "+t2)
			// fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))

			tmplSettings, _ := template.ParseFiles("templates/gamePage/settings.html")
			tmplSettings.Execute(w, user)

			data := Score{
				StateWINER:  stateWINER,
				StateLOSSER: stateLOSSER,
				StateDROW:   stateDRAW,
			}
			tmplStore, _ := template.ParseFiles("templates/gamePage/store.html")
			tmplStore.Execute(w, data)

			stateOf := struct {
				State string
			}{
				State: t2,
			}
			tmplState, _ := template.ParseFiles("templates/gamePage/state.html")
			tmplState.Execute(w, stateOf)

			tmplGame, _ := template.ParseFiles("templates/gamePage/game.html")
			tmplGame.Execute(w, data)
		}
	case SCISSORS:
		{
			_, t2 := gameOne(SCISSORS)
			scoreOfGame(t2)
			// fmt.Fprintf(w, "You = "+SCISSORS+" AI = "+t1+" State = "+t2)
			// fmt.Fprintf(w, fmt.Sprintf("\n stateWINER = %d \n stateLOSSER = %d \n stateDROW %d \n", stateWINER, stateLOSSER, stateDROW))

			tmplSettings, _ := template.ParseFiles("templates/gamePage/settings.html")
			tmplSettings.Execute(w, user)

			data := Score{
				StateWINER:  stateWINER,
				StateLOSSER: stateLOSSER,
				StateDROW:   stateDRAW,
			}
			tmplStore, _ := template.ParseFiles("templates/gamePage/store.html")
			tmplStore.Execute(w, data)

			stateOf := struct {
				State string
			}{
				State: t2,
			}
			tmplState, _ := template.ParseFiles("templates/gamePage/state.html")
			tmplState.Execute(w, stateOf)

			tmplGame, _ := template.ParseFiles("templates/gamePage/game.html")
			tmplGame.Execute(w, data)
		}
	default:
		tmplSettings, _ := template.ParseFiles("templates/gamePage/settings.html")
		tmplSettings.Execute(w, user)

		data := Score{
			StateWINER:  stateWINER,
			StateLOSSER: stateLOSSER,
			StateDROW:   stateDRAW,
		}

		tmplStore, _ := template.ParseFiles("templates/gamePage/store.html")
		tmplStore.Execute(w, data)

		tmplGame, _ := template.ParseFiles("templates/gamePage/game.html")
		tmplGame.Execute(w, data)
	}
}

func clearCost(w http.ResponseWriter, r *http.Request) {
	stateWINER = 0
	stateLOSSER = 0
	stateDRAW = 0
	http.Redirect(w, r, "/game", http.StatusMovedPermanently)
}

func rename(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userName = r.FormValue("name")
		// password := r.URL.Query().Get("password")
		http.Redirect(w, r, "/game", http.StatusMovedPermanently)
	} else if r.Method == http.MethodGet {
		user := User{
			UserName: userName,
			Password: "123",
		}
		tmplRename, _ := template.ParseFiles("templates/rename.html")
		tmplRename.Execute(w, user)
	}
	// parametrsOfRequest := r.URL.RawQuery
	// strings.TrimPrefix(parametrsOfRequest, "name=")
}

func scoreOfGame(state string) {
	if state == WINER {
		stateWINER++
	}
	if state == LOSSER {
		stateLOSSER++
	}
	if state == DROW {
		stateDRAW++
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
