package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func LoadFloats(fn string) ([]float64, error) {
	var numbers []float64
	fs, err := os.Open(fn)
	if err != nil {
		return numbers, err
	}
	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
		// fmt.Println(scanner.Text())
	}
	err = fs.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, err
}

func main() {
	var sum float64 = 0
	numbers, err := LoadFloats("scores.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, number := range numbers {
		sum = sum + number
	}
	fmt.Printf("평균 : %.2f\n", sum/float64(len(numbers)))
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// )

// func LoadFloats(fn string) ([4]float64, error) {
// 	var numbers [4]float64
// 	var i int = 0
// 	fs, err := os.Open(fn)
// 	if err != nil {
// 		return numbers, err
// 	}
// 	scanner := bufio.NewScanner(fs)
// 	for scanner.Scan() {
// 		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
// 		if err != nil {
// 			return numbers, err
// 		}
// 		i++
// 		// fmt.Println(scanner.Text())
// 	}
// 	err = fs.Close()
// 	if err != nil {
// 		return numbers, err
// 	}
// 	if scanner.Err() != nil {
// 		return numbers, scanner.Err()
// 	}
// 	return numbers, err
// }

// func main() {
// 	var sum float64 = 0
// 	numbers, err := LoadFloats("scores.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, number := range numbers {
// 		sum = sum + number
// 	}
// 	fmt.Printf("평균 : %.2f\n", sum/float64(len(numbers)))
// }
