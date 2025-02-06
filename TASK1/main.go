package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
)

type Response struct {
	Number         int `json:"number"`
	IsPrime   bool `json:"is_prime"`
	IsPerfect bool `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum int `json:"digit_sum"`
	FunFact string `json:"fun_fact"`
}

type BadRequest struct{
	Number string `json:"number"`
	Error bool `json:"error"`
}

type NumbersAPIResponse struct {
    Text   string `json:"text"`
    Number int    `json:"number"`
    Type   string `json:"type"`
    Found  bool   `json:"found"`
}


func checkForIntOrFloat(value string) string {
	if value == "" {
		return "empty"
	}

	if _, err := strconv.Atoi(value); err == nil{
		return "int"
	}

	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "float"
	}

	if _, err := strconv.ParseBool(value); err == nil{
		return "bool"
	}

	return "alphabet"
}

func checkValidType(value string) *BadRequest{
	// check type
	typeOfNumber := checkForIntOrFloat(value);

	validTypes := []string{"alphabet", "empty", "bool"}

	for _, v := range validTypes{
		if v == typeOfNumber{
			return &BadRequest{
				Number: typeOfNumber,
				Error: true,
			}
		}
	}
	return nil
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumDigits(value string) int {
	total := 0

	for _, char := range value {
		if char >= '0' && char <= '9' {
			digit, _ := strconv.Atoi(string(char))
			total += digit
		}
	}
	return total
}
func isPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}
func isArmstrong(n int) bool {
	strNum := strconv.Itoa(n)
	digits := len(strNum)
	sum := 0

	for _, digit := range strNum {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += int(math.Pow(float64(digitInt), float64(digits)))
	}

	return sum == n
}
func categorizeNumberProperties(n int) []string {
	result := []string{}

	if isArmstrong(n) {
		result = append(result, "armstrong")
	}

	if n%2 == 0 {
		result = append(result, "even")
	} else {
		result = append(result, "odd")
	}

	return result
}

func getNumberFact(number int) (string, error){
	url := fmt.Sprintf("http://numbersapi.com/%d/math?json", number)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}
	var response NumbersAPIResponse
	_ = json.Unmarshal(body, &response)
	return response.Text, nil
}

func prepareValidResponse(value string) *Response {
	number, _ := strconv.Atoi(value)
	isPrime := isPrime(number)
	isPerfect := isPerfect(number)
	properties := categorizeNumberProperties(number)
	digitSum := sumDigits(value)
	funcFact, _ := getNumberFact(number)

	return &Response{
		Number: number,
		IsPrime: isPrime,
		IsPerfect: isPerfect,
		Properties: properties,
		DigitSum: digitSum,
		FunFact: funcFact,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// fetch query params
	number := r.URL.Query().Get("number")

	var response interface{}

	isValid := checkValidType(number)

	if isValid != nil {
		response = isValid
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(response)
		return
	}

	response = prepareValidResponse(number)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/classify-number", handler)
	// logger
	loggedRouter := handlers.LoggingHandler(log.Writer(), http.DefaultServeMux)

	log.Println("Server is starting...")
	http.ListenAndServe(":7001", loggedRouter)
}
