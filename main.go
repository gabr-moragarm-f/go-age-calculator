package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("What year were you born?")
	reader := bufio.NewReader(os.Stdin)

	response, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	response = strings.TrimSuffix(response, "\n")

	year, err := strconv.Atoi(response)

	if err != nil {
		fmt.Println(response, "is not a valid year")
		log.Fatal(err)
	}

	currentYear := time.Time.Year(time.Now())

	age := currentYear - int(year)

	fmt.Println("You shoud have", age, "years!")

}
