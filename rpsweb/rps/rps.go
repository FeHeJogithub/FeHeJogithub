package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // LOGICA PIEDRA. VENCE A LAS TIJERAS. (TIJERAS +1)%3=0
	PAPER    = 1 // PAPEL. VENCE A LA PIEDRA. (PIEDRA +1)%3=1
	SCISSORS = 2 // TIJERAS. VENCE A PAPEL. (PAPEL +1)%=2

)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

// Mensaje para cuando pierde
var winMessages = []string{
	"¡Felicidades!",
	"¡Lo lograste!",
	"¡Victoria aplastante!",
}
var loseMessages = []string{
	"Que lastima!",
	"Intentalo de nuevo!",
	"Hoy simplemente no es tu dia.",
}

//Mensaje de empate

var drawMessages = []string{
	"Las grandes mentes piensan igual.!",
	"Oh oh. Intentalo de nuevo!",
	"Nadie gana, pero puedes intentarlo de nuevo.",
}

//variables para el puntaje

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	//Mensaje dependiendo de lo que ha elegido la computadora

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligio PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligio TIJERAS"

	}

	messageInt := rand.Intn(3)

	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"

		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana!"
		message = winMessages[messageInt]
	} else {
		ComputerScore++

		roundResult = "La computadora gana!"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
