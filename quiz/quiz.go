package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var open func(name string) (io.Reader, error)
var args []string
var stderr io.Writer
var stdout io.Writer
var stdin io.Reader
var exit func(int)

func setup() {
	open = func(name string) (io.Reader, error) {
		return os.Open(name)
	}
	args = os.Args
	stderr = os.Stderr
	stdout = os.Stdout
	stdin = os.Stdin
	exit = os.Exit
}

func init() {
	setup()
}

func main() {
	logger := log.New(stderr, fmt.Sprintf("%s-", args[0]), log.LstdFlags)
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	fs.SetOutput(stderr)
	csvFile := fs.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := fs.Uint("limit", 30, "the time limit for the quiz in seconds")
	parseError := fs.Parse(args[1:])
	if parseError != nil {
		exit(1)
		return
	}
	limitDuration, _ := time.ParseDuration(fmt.Sprintf("%ds", *limit))
	var reader io.Reader
	var err error
	reader, err = open(*csvFile)
	if err != nil {
		logger.Println("Could not open csv file", err)
		exit(1)
		return
	}
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		logger.Println("error in reading data file", err)
		exit(1)
		return
	}
	right, total := 0, len(records)
	ctx, cancel := context.WithTimeout(context.Background(), limitDuration)
	defer cancel()
RecordLoop:
	for problemNum, record := range records {
		fmt.Fprintf(stdout, "Problem #%d: %s = ", problemNum+1, record[0])
		result := make(chan string, 1)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			var line string
			fmt.Fscanf(stdin, "%s", &line)
			result <- line
			wg.Done()
		}()
		select {
		case <-ctx.Done():
			wg.Wait()
			break RecordLoop
		case answer := <-result:
			wg.Wait()
			if strings.TrimSpace(answer) == strings.TrimSpace(record[1]) {
				right++
			}
		}
	}
	fmt.Fprintf(stdout, "\nYou scored %d out of %d.\n", right, total)
	exit(0)
}
