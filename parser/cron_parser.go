package parser

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	minuteField = iota
	hourField
	dayOfMonthField
	monthField
	dayOfWeekField
	commandField
)

type CronParser interface {
	Parse(cronString string) (Cron, error)
}

type Cron struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Command    string
}

type DefaultCronParser struct{}

func (p *DefaultCronParser) Parse(cronString string) (Cron, error) {
	fields := strings.Fields(cronString)
	if len(fields) < 6 {
		return Cron{}, fmt.Errorf("invalid cron string format")
	}

	return Cron{
		Minute:     expandField(fields[minuteField], 0, 59),
		Hour:       expandField(fields[hourField], 0, 23),
		DayOfMonth: expandField(fields[dayOfMonthField], 1, 31),
		Month:      expandField(fields[monthField], 1, 12),
		DayOfWeek:  expandField(fields[dayOfWeekField], 0, 6),
		Command:    fields[commandField],
	}, nil
}

func expandField(field string, min, max int) string {
	var result []string
	if field == "*" {
		for i := min; i <= max; i++ {
			result = append(result, strconv.Itoa(i))
		}
	} else if strings.Contains(field, ",") {
		parts := strings.Split(field, ",")
		for _, part := range parts {
			result = append(result, expandPart(part, min, max)...)
		}
	} else {
		result = append(result, expandPart(field, min, max)...)
	}
	return strings.Join(result, " ")
}

func expandPart(part string, min, max int) []string {
	var result []string
	if strings.Contains(part, "-") {
		rangeParts := strings.Split(part, "-")
		start, _ := strconv.Atoi(rangeParts[0])
		end, _ := strconv.Atoi(rangeParts[1])
		for i := start; i <= end; i++ {
			result = append(result, strconv.Itoa(i))
		}
	} else if strings.Contains(part, "/") {
		stepParts := strings.Split(part, "/")
		step, _ := strconv.Atoi(stepParts[1])
		for i := min; i <= max; i += step {
			result = append(result, strconv.Itoa(i))
		}
	} else {
		result = append(result, part)
	}
	return result
}
