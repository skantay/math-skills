package mathSkills

import (
	"bufio"
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
	errFileName = errors.New("err: run() ---> validateFiles()")
	errGetText  = errors.New("err: run() ---> getText()")
	errPutFile  = errors.New("err: run() ---> putFile()")
)

func Run() error {
	file, err := getFiles()
	if err != nil {
		return fmt.Errorf("%s: %w", errGetFiles, err)
	}
	if err := validateFiles(file); err != nil {
		return fmt.Errorf("%s: %w", errFileName, err)
	}
	numbers, err := getText(file)
	if err != nil {
		return fmt.Errorf("%s: %w", errGetText, err)
	}
	process(numbers)
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

func getText(file string) ([]int, error) {
	fileO, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fileO)
	var numbers []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func process(numbers []int) {
	if len(numbers) == 0 {
		return
	}
	fmt.Printf("Average: %d\n", int(round(average(numbers))))
	fmt.Printf("Median: %d\n", int((median(numbers))))
	fmt.Printf("Variance: %d\n", int(round(variance(numbers))))
	fmt.Printf("Standard Deviation: %d\n", int(round(math.Sqrt(float64(variance(numbers))))))
}

func average(numbers []int) float64 {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return float64(sum) / float64(len(numbers))
}

func median(numbers []int) float64 {
	sort.Ints(numbers)
	if len(numbers)%2 != 0 {
		return float64(numbers[len(numbers)+1/2])
	}
	l1 := (len(numbers) / 2)
	l2 := (len(numbers) / 2) + 1
	ave := []int{
		numbers[l1],
		numbers[l2],
	}
	return average(ave)
}

func variance(numbers []int) float64 {
	mean := average(numbers)
	var result float64
	for _, v := range numbers {
		dif := (float64(v) - mean)
		result += float64(dif) * float64(dif)
	}
	return result / float64(len(numbers))
}

func round(x float64) float64 {
	return math.Round(x)
}
