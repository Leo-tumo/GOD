package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	fmt.Printf("%q\n\n", ships)

	from, to := 0, len(ships)

	switch len(os.Args) {
	default:
		fallthrough
	case 1:
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	case 3:
		to, _ = strconv.Atoi(os.Args[2])
		fallthrough
	case 2:
		from, _ = strconv.Atoi(os.Args[1])
	}

	if l := len(ships); from < 0 || to > l || from > to {
		fmt.Println("Wrong positions")
		return
	}

	fmt.Println(ships[from:to])

}
