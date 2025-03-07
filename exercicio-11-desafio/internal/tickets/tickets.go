package tickets

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Night     = "night"
	Morning   = "morning"
	Afternoon = "afternoon"
	Evening   = "evening"
)

var (
	ErrInvalidPeriod = errors.New("not a valid period")
)

type Ticket struct {
	ID          int
	Name        string
	Email       string
	Destination string
	Time        time.Time
	Price       int
}

var tickets []Ticket

func GetTicketsFromCSV(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		details := strings.Split(line, ",")

		id, errId := strconv.Atoi(strings.TrimSpace(details[0]))
		airplaneTime, errTime := time.Parse("15:04", details[4])
		airplane, errPrice := strconv.Atoi(strings.TrimSpace(details[5]))

		var parseError error = nil
		if errId != nil {
			parseError = fmt.Errorf("error while parsing id in line: %w", errId)
		} else if errTime != nil {
			parseError = fmt.Errorf("error while parsing time in line %d, with time %s, error: %w", id, details[4], errTime)
		} else if errPrice != nil {
			parseError = fmt.Errorf("error while parsing price in line: %d, %w", id, errPrice)
		}

		if parseError != nil {
			fmt.Println(parseError)
		} else {
			tickets = append(tickets, Ticket{
				ID:          id,
				Name:        details[1],
				Email:       details[2],
				Destination: details[3],
				Time:        airplaneTime,
				Price:       airplane,
			})
		}
	}

	if scanner.Err() != nil {
		panic("Não foi possível ler o arquivo adequadamente")
	}

	// fmt.Println(tickets)
}

// ejemplo 1
func GetTotalTicketsByDestination(destination string) (total int, err error) {
	for _, ticket := range tickets {
		if ticket.Destination == destination {
			total += 1
		}
	}
	return total, nil
}

// ejemplo 2
func GetCountByPeriod(period string) (total int, err error) {
	switch period {
	case Night:
		total, err = searchByPeriod("00:00", "06:59")
	case Morning:
		total, err = searchByPeriod("07:00", "12:59")
	case Afternoon:
		total, err = searchByPeriod("13:00", "19:59")
	case Evening:
		total, err = searchByPeriod("20:00", "23:59")
	default:
		return 0, ErrInvalidPeriod
	}
	return
}

func searchByPeriod(begin string, end string) (int, error) {
	total := 0

	beginTime, beginErr := time.Parse("15:04", begin)
	if beginErr != nil {
		return total, beginErr
	}

	endTime, endErr := time.Parse("15:04", end)
	if endErr != nil {
		return total, endErr
	}

	for _, ticket := range tickets {
		check := ticket.Time
		if inTimeSpan(beginTime, endTime, check) {
			total += 1
		}

	}
	return total, nil
}

func inTimeSpan(start, end, check time.Time) bool {
	return !check.After(end) && !check.Before(start)
}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	totalByDestination, err := GetTotalTicketsByDestination(destination)
	if err != nil {
		return 0, err
	}

	var percentage float64 = float64(totalByDestination) / float64(total)
	return int(percentage * 100), nil
}
