package utils

import (
	"fmt"
	"math/rand"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func InprinteAscii() {
	fmt.Println("  ___ _  _ ___ ___ ___ _  _ _____ ___")
	fmt.Println(" |_ _| \\| | _ \\ _ \\_ _| \\| |_   _| __|")
	fmt.Println("  | || .` |  _/   /| || .` | | | | _|")
	fmt.Println(" |___|_|\\_|_| |_|_\\___|_|\\_| |_| |___|\n\n	")
}
