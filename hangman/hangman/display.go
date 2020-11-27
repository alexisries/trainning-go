package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`
	######     #####                 #####                
	#     # # #     #   ##     ##   #     #  ####   ####  
	#     # # #        #  #   #  #  #       #      #    # 
	######  # #       #    # #    #  #####   ####  #    # 
	#       # #       ###### ######       #      # #    # 
	#       # #     # #    # #    # #     # #    # #    # 
	#       #  #####  #    # #    #  #####   ####   ####  
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
  ,==========Y===
  || //      |
  ||//       |
  ||/        O
  ||        /|\
  ||        / \
  ||
 /||
//||
============
		`
	case 1:
		draw = `
  ,==========Y===
  || //      |
  ||//       |
  ||/        O
  ||        /|\
  ||        /
  ||
 /||
//||
============
		`
	case 2:
		draw = `
  ,==========Y===
  || //      |
  ||//       |
  ||/        O
  ||        /|\
  ||
  ||
 /||
//||
============
		`
	case 3:
		draw = `
  ,==========Y===
  || //      |
  ||//       |
  ||/        O
  ||
  ||
  ||
 /||
//||
============
		`
	case 4:
		draw = `
  ,==========Y===
  || //
  ||//
  ||/
  ||
  ||
  ||
 /||
//||
============
		`
	case 5:
		draw = `
  ,==========
  || //
  ||//
  ||/
  ||
  ||
  ||
 /||
//||
============
		`
	case 6:
		draw = `
  
  || // 
  ||//
  ||/
  ||
  ||
  ||
 /||
//||
============
		`
	case 7:
		draw = `
  








 /||
//||
============
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Println("Used: ")
	drawLetters(g.UseLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess !")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used", guess)
	case "badGuess":
		fmt.Printf("Bad guess '%s' is not the word", guess)
	case "lost":
		fmt.Print("You lost :(! The word was:")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("YOU WON! The word was:")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}
