package main

import (
	"fmt"
	"sorter/sorter/algorithms/bubblesort"
	"sorter/sorter/algorithms/qsort"
	"flag"
	"os"
	"bufio"
	"io"
	"strconv"
	"time"
)
/*
func main() {
	arr := []int{1, 2, 10, 9, 4, 5}
	arr1 := []int{1, 2, 10, 9, 8, 7}
	fmt.Println(arr1)
	bubblesort.BubbleSort(arr)
	qsort.QuickSort(arr1)
	fmt.Println("sorted arr :", arr)
	fmt.Println()
	fmt.Println("sorted arr1 :", arr1)
}
*/

var infile *string = flag.String("i", "unsorted.dat","File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ", algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complate")
		writeValues(values, *outfile)
	}else {
		fmt.Println(err)
	}
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file", file)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected")
			return
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)

	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil{
		fmt.Println("Failed to create the outfile ", outfile)
		return err
	}
	defer  file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}



