package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	open   func(name string) (io.Reader, error)
	args   []string
	stderr io.Writer
	stdout io.Writer
	stdin  io.Reader
	exit   func(int)
	stop   = make(chan os.Signal, 1)
)

func setup() {
	rand.Seed(time.Now().UnixNano())
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
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	fs.SetOutput(stderr)
	csvFile := fs.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := fs.Uint("limit", 30, "the time limit for the quiz in seconds")
	shuffle := fs.Bool("shuffle", false, "shuffle the input")
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
	for i := 0; i < len(records); i++ {
		if *shuffle {
			if remain := len(records) - i - 1; remain != 0 {
				j := rand.Intn(remain) + i
				records[i], records[j] = records[j], records[i]
			}
		}
		fmt.Fprintf(stdout, "Problem #%d: %s = ", i+1, records[i][0])
		result := make(chan string, 1)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			var line string
			fmt.Fscanln(stdin, &line)
			result <- line
			wg.Done()
		}()
		select {
		case <-stop:
			cancel()
			wg.Wait()
			break RecordLoop
		case <-ctx.Done():
			wg.Wait()
			break RecordLoop
		case answer := <-result:
			wg.Wait()
			if strings.ToLower(strings.TrimSpace(answer)) == strings.ToLower(strings.TrimSpace(records[i][1])) {
				right++
			}
		}
	}
	fmt.Fprintf(stdout, "\nYou scored %d out of %d.\n", right, total)
	exit(0)
}
