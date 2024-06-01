package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("When were you born? (yyyy-mm-dd)")
	reader := bufio.NewReader(os.Stdin)

	response, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	response = strings.TrimSuffix(response, "\n")

	birthDate, err := time.Parse("2006-01-02", response)

	if err != nil {
		fmt.Println(response, "is not a valid date(yyyy-mm-dd)")
		log.Fatal(err)
	}

	age := int(time.Since(birthDate).Hours() / 24 / 365.25)

	fmt.Println("If I'm not wrong you have", age, "years!")

}
