package main

import (
	"fmt"
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
)

type numberGuessModel struct {
    secretNumber int     // Game data
    guess        string  // User input
    message      string  // UI feedback
    attempts     int     // Game progress
    gameOver     bool    // State flag
}


func main() {
	var guessingNumber = rand.Intn(10)
	var m = numberGuessModel{
		secretNumber: guessingNumber,
		guess:        "",
		message:      "Welcome to the Number Guessing Game! I'm thinking of a number between 1 and 10. Can you guess what it is?",
		attempts:     0,
		gameOver:     false,
	}
	tea.NewProgram(m).Run()
}

func (m numberGuessModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
			if msg.String() == "ctrl+c" || msg.String() == "q" {
				return m, tea.Quit
			}

			switch msg.Type {
			case tea.KeyEnter:
				if m.gameOver {
					return m, tea.Quit
				}

				var guessInt int
				_, err := fmt.Sscanf(m.guess, "%d", &guessInt)
				if err != nil {
					m.message = "Please enter a valid number."
					m.guess = ""
					return m, nil
				}
				m.attempts++

				processGuessing(guessInt, &m)

			case tea.KeyBackspace:
				if len(m.guess) > 0 {
					m.guess = m.guess[:len(m.guess)-1]
				}
			case tea.KeyRunes:
				m.guess += string(msg.Runes)
			}
	}
	return m, nil
}

func processGuessing(guessInt int, m *numberGuessModel) {
	if guessInt < m.secretNumber {
		m.message = "Too low! Try again."
	} else if guessInt > m.secretNumber {
		m.message = "Too high! Try again."
	} else {
		m.message = fmt.Sprintf("Congratulations! You've guessed the number %d in %d attempts!", m.secretNumber, m.attempts)
		m.gameOver = true
	}
	m.guess = ""
}

func (m numberGuessModel) View() string {
	s := m.message + "\n\n"
	s += fmt.Sprintf("Guess: %s\n", m.guess)
	s += fmt.Sprintf("Attempts: %d\n\n", m.attempts)

	if m.gameOver {
		s += "Press Enter or q to quit.\n"
	} else {
		s += "Enter your guess and press Enter. Press Ctrl+C or q to quit.\n"
	}

	return s
}

func (m numberGuessModel) Init() tea.Cmd {
    return nil
}