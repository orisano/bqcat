package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/orisano/gproject"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

func main() {
	if err := run(); err != nil {
		log.Fatal("bqcat: ", err)
	}
}

func run() error {
	var projectID string
	flag.StringVar(&projectID, "p", "", "project id")

	var queryPath string
	flag.StringVar(&queryPath, "f", "", "query file path")

	flag.Parse()

	var query string
	switch {
	case queryPath != "":
		b, err := ioutil.ReadFile(queryPath)
		if err != nil {
			return err
		}
		query = string(b)
	case flag.NArg() >= 1:
		query = flag.Arg(0)
	default:
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		query = string(b)
	}

	if projectID == "" {
		projectID = gproject.Default()
	}

	ctx := context.Background()
	bq, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return errors.Wrap(err, "failed to create BigQuery client")
	}

	q := bq.Query(query)
	job, err := q.Run(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to run query")
	}
	rows, err := job.Read(ctx)
	if err != nil {
		return err
	}

	for {
		var row []bigquery.Value
		err := rows.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for i, col := range row {
			fmt.Print(col)
			if i == len(row)-1 {
				fmt.Println()
			} else {
				fmt.Print(",")
			}
		}
	}
	return nil
}
