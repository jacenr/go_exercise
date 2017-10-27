package main

import (
	"fmt"
	"github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	timeNow := time.Now()
	aMonth := time.Date(timeNow.Year(), timeNow.Month()-1, timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond(), timeNow.Location())
	aYear := time.Date(timeNow.Year()-1, timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond(), timeNow.Location())
	// fmt.Println(timeNow)
	// fmt.Println(aMonth)
	// fmt.Println(aYear)
	// MapItems := make(map[string][]*github.Issue)
	LtAMonth := []*github.Issue{}
	AYearToAMonth := []*github.Issue{}
	MtAYear := []*github.Issue{}
	// MapItems["LtAMonth"] = LtAMonth
	// MapItems["AYearToAMonth"] = AYearToAMonth
	// MapItems["MtAYear"] = MtAYear
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.Before(aYear):
			MtAYear = append(MtAYear, item)
		case item.CreatedAt.Before(aMonth) && item.CreatedAt.After(aYear):
			AYearToAMonth = append(AYearToAMonth, item)
		default:
			LtAMonth = append(LtAMonth, item)
		}
	}
	fmt.Println("Issue of mort than a year: ", len(MtAYear))
	for _, item := range MtAYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Issue of less than a year but more then a month: ", len(AYearToAMonth))
	for _, item := range AYearToAMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Issue of less than a month: ", len(LtAMonth))
	for _, item := range LtAMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
