package mathSkills

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

var (
	errGetFiles = errors.New("err: run() ---> getFiles()")
	errGetText  = errors.New("err: run() ---> getText()")
	errFileName = errors.New("err: run() ---> validateFiles()")
	errPutFile  = errors.New("err: run() ---> putFile()")
)

func Run() error {
	file, err := getFiles()
	if err != nil {
		return fmt.Errorf("%s: %w", errGetFiles, err)
	}

	if err := validateFiles(file); err != nil {
		return fmt.Errorf("%s: %w", errGetText, err)
	}

	text, err := getText(file)
	if err != nil {
		return fmt.Errorf("%s: %w", errFileName, err)
	}

	if text == "" {
		return nil
	}

	process(text)

	return nil
}

func Error() error {
	err1 := errors.New("did you README")
	err2 := errors.New("hmmmmm, I think you did not README")
	err3 := errors.New("why would you skip my README")
	err4 := errors.New("bruh, something is wrong")
	err5 := errors.New("i want an ice cream")

	errors := []error{
		err1,
		err2,
		err3,
		err4,
		err5,
	}

	return errors[randInt(0, len(errors))]
}

// Helper of Error().
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func getFiles() (string, error) {
	args := os.Args[1:]
	if one := 1; len(args) != one {
		return "", errors.New("enter ONLY one file path")
	}

	return args[0], nil
}

func validateFiles(file string) error {
	fileExt := filepath.Ext(file)
	errorText := "file extension is invalid. Provide ONLY .txt files"

	if fileExt != ".txt" && fileExt != "" {
		return fmt.Errorf("input, %s", errorText)
	}

	return nil
}

func getText(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("%s no such file", file)
	}

	return string(data), nil
}

func process(text string) {
	numbers := make([]int, 0, len(text))
	var num string
	for _, v := range text {
		if v == '\n' {
			number, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Not valid")
				return
			}
			num = ""
			numbers = append(numbers, number)
			continue
		}
		num += string(v)
	}
	fmt.Printf("Average: %d\n", average(numbers))
	fmt.Printf("Median: %d\n", median(numbers))
	fmt.Printf("Variance: %d\n", int(round(variance(numbers))))
	fmt.Printf("Standard Deviation: %d\n", int(round(math.Sqrt(float64(variance(numbers))))))
}

func average(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum / len(numbers)
}

func median(numbers []int) int {
	sort.Ints(numbers)

	if len(numbers)%2 != 0 {
		return numbers[len(numbers)/2]
	}

	l1 := (len(numbers) / 2) - 1
	l2 := (len(numbers) / 2)
	return average([]int{numbers[l1], numbers[l2]})
}

func variance(numbers []int) float64 {
	mean := average(numbers)
	var result float64
	for _, v := range numbers {
		dif := (v - mean)
		result += float64(dif) * float64(dif)
	}

	return result / float64(len(numbers))
}

func round(x float64) float64 {
	return math.Round(x)
}
