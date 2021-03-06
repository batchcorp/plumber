package main

import (
	"bufio"
	"context"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/batchcorp/plumber/cli"
	"github.com/batchcorp/plumber/plumber"
	"github.com/batchcorp/plumber/printer"
	"github.com/batchcorp/plumber/stats"
)

func main() {
	cmd, opts, err := cli.Handle(os.Args[1:])
	if err != nil {
		logrus.Fatalf("Unable to handle CLI input: %s", err)
	}

	readFromStdin(opts)

	if opts.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if opts.Quiet {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	serviceCtx, serviceShutdownFunc := context.WithCancel(context.Background())
	mainCtx, mainShutdownFunc := context.WithCancel(context.Background())

	// We only want to intercept these in relay mode
	if strings.HasPrefix(cmd, "relay") {
		logrus.Debug("Intercepting signals")
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGTERM)

		go func() {
			signal := <-c
			logrus.Debugf("Received system call: %+v", signal)

			serviceShutdownFunc()
		}()

		// Create prometheus counters/gauges
		stats.InitPrometheusMetrics()
	}

	if opts.Stats {
		stats.Start(opts.StatsReportInterval)
	}

	// In container mode, force JSON and don't print logo
	if !terminal.IsTerminal(int(os.Stderr.Fd())) || opts.Batch.OutputType == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		printer.PrintLogo()
	}

	p, err := plumber.New(&plumber.Config{
		ServiceShutdownCtx: serviceCtx,
		MainShutdownFunc:   mainShutdownFunc,
		MainShutdownCtx:    mainCtx,
		Cmd:                cmd,
		Options:            opts,
	})

	if err != nil {
		logrus.Fatal(err)
	}

	p.Run()
}

// readFromStdin reads data piped into stdin
func readFromStdin(opts *cli.Options) {
	info, err := os.Stdin.Stat()
	if err != nil {
		logrus.Fatal(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return
	}

	inputData := make([]string, 0)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err == io.EOF {
			break
		}
		inputData = append(inputData, strings.Trim(string(line), "\n"))
	}

	// Treat input as a JSON array
	if opts.WriteInputIsJsonArray {
		opts.WriteInputData = convertJSONInput(inputData[0])
		return
	}

	// Treat input as new object on each line
	opts.WriteInputData = inputData
}

// convertJSONInput converts a JSON array to a slice of strings for the writer to consume
func convertJSONInput(value string) []string {
	inputData := make([]string, 0)
	jsonArray := gjson.Parse(value)
	if !jsonArray.IsArray() {
		logrus.Fatal("--json-array option was passed, but input data is not a valid JSON array")
	}

	jsonArray.ForEach(func(key, value gjson.Result) bool {
		inputData = append(inputData, value.Raw)
		return true
	})

	return inputData
}
