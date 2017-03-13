package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/transitorykris/hypnos"
)

const usage = `
Usage: Just like cron!

        ┌───────── minute (0 - 59)
        │ ┌───────── hour (0 - 23)
        │ │ ┌───────── day of month (1 - 31)
        │ │ │ ┌───────── month (1 - 12)
        │ │ │ │ ┌───────── day of week (0 - 6) (Sunday to Saturday;
        │ │ │ │ │                                   7 is also Sunday)
hypnos "* * * * *"

For more details, see https://en.wikipedia.org/wiki/Cron
`

func main() {
	logger := log.New(os.Stderr, "", 0)

	flag.Usage = func() {
		logger.Println(usage)
		flag.PrintDefaults()
	}

	var silent bool
	flag.BoolVar(&silent, "q", false, "Make hypnos quiet")
	flag.Parse()

	if silent {
		logger = log.New(ioutil.Discard, "", 0)
	}

	if len(flag.Args()) != 1 {
		logger.Println(usage)
		os.Exit(1)
	}

	duration, date, err := hypnos.Sleep(flag.Args()[0])
	if err != nil {
		logger.Println(usage)
		os.Exit(1)
	}

	logger.Println("Next run at", date, "Sleeping for", duration)
	time.Sleep(duration)
}
