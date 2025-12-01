package aoc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/landanqrew/advent-of-code-go/internal/client"
)


var inputClient = client.NewHttpClient("https://adventofcode.com")
var solutionClient = client.NewHttpClient("https://aoc.fornwall.workers.dev")

func LoadEnv(envPath string) error {
	err := godotenv.Load(envPath)
	if err != nil {
		return err
	}
	return nil
}

/**
* Get the session cookie from the .env file
* @return string, error

* Log in to Advent of Code and access and puzzle input page (e.g. http://adventofcode.com/2022/day/1/input)
* Right click the page and click "inspect"
* Navigate to the "Network" tab
* Click on any request, and go to the "Headers" tab
* Search through the "Request Headers" for a header named cookie.
* You should find one value that starts with session=, followed by a long string of hexadecimal characters. Copy the whole value, starting with session= and including all the hex characters until you hit a semicolon.
* Save this value as an environment variable on your system using the name COOKIE.
*/
func GetSessionCookie(envPath string) (string, error) {
	cookie := os.Getenv("COOKIE")
	if cookie == "" {
		LoadEnv(envPath)
	}
	cookie = os.Getenv("COOKIE")
	if cookie == "" {
		return "", fmt.Errorf("cookie not found")
	}
	return cookie, nil
}

func GetInput(year int, day int, envPath string) ([]byte, error) {
	cookie, err := GetSessionCookie(envPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", inputClient.BaseURL + fmt.Sprintf("/%d/day/%d/input", year, day), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "session=" + cookie)
	resp, err := inputClient.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetSolution(year int, day int, part int, input string) (string, error) {
	buffer := bytes.NewBuffer([]byte(input))
	resp, err := solutionClient.Client.Post(solutionClient.BaseURL + fmt.Sprintf("/%d/day/%d/part/%d", year, day, part), "text/plain", buffer)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}