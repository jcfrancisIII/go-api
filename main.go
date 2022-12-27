package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// CompareArrs returns bool, takes two string arrays and checks if they match by join ""
func CompareArrs(solution []string, scramble []string) bool {
	return strings.Join(solution, "") == strings.Join(scramble, "")
}

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func main() {
	// solution := []string{"t", "o", "a", "s", "t"}
	// scramble := []string{"", "o", "", "s", "t"}
	// tries := 0
	// abc := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	// fmt.Printf("Start" + "\n")

	// for i := 0; i < 26; i++ {
	// 	tries++

	// 	// for j := 0; j < 26; j++ {
	// 	// 	tries++

	// 	// 	if strings.Join(solution, "") == strings.Join(scramble, "") {
	// 	// 		fmt.Printf("j loop: Solved %s in %s tries\n", strings.Join(scramble, ""), strconv.Itoa(tries))
	// 	// 		break
	// 	// 	}
	// 	// }
	// 	fmt.Printf("i loop: Solved %s in %s tries\n", strings.Join(scramble, ""), strconv.Itoa(tries))
	// 	if strings.Join(solution, "") == strings.Join(scramble, "") {
	// 		fmt.Printf("i loop: Solved %s in %s tries\n", strings.Join(scramble, ""), strconv.Itoa(tries))
	// 		break
	// 	}
	// }
	args := os.Args

	if len(args) < 2 { // args always has one element for the exe
		fmt.Printf("Usage: go-api <url>\n")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in an invalid format: %s\n", err)
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	// test go run main.go https://jsonplaceholder.typicode.com/todos/1
	fmt.Printf("HTTP Status Code: %s\nBody: %s\n", response.StatusCode, body)

}
