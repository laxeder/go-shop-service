package date

import (
	"strconv"
	"time"
)

// devolve UTC
func NowUTC() string {
	return time.Now().Format(time.RFC3339Nano)
}

// string timespamp unix para UTC
func TimestampToUTC(timestamp string) string {
	timestampInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return ""
	}
	return time.Unix(timestampInt, 0).Format(time.RFC3339Nano)
}

// comverte string UUC para time.Time
func UTCToTime(dateUTC string) time.Time {
	timestamp, err := time.Parse(time.RFC3339Nano, dateUTC)
	if err != nil {
		return time.Time{}
	}
	return timestamp
}

// transforma uma string UTC para saber se é válida
func UTCValid(dateUTC string) bool {
	_, err := time.Parse(time.RFC3339Nano, dateUTC)
	if err != nil {
		return false
	}
	return true
}

// transforma uma string BR para objeto time
func BRToTime(dateBR string) time.Time {
	timestamp, err := time.Parse("02/01/2006", dateBR)
	if err != nil {
		return time.Time{}
	}
	return timestamp
}

// transforma uma string BR para string UTC
func BRToUTC(dateBR string) string {
	timestamp, err := time.Parse("02/01/2006", dateBR)
	if err != nil {
		return ""
	}
	return timestamp.Format(time.RFC3339Nano)
}

// transforma uma string BR para saber se é válida
func BRValid(dateBR string) bool {
	_, err := time.Parse("02/01/2006", dateBR)
	if err != nil {
		return false
	}
	return true
}

func TimeToUTC(date time.Time) string {
	return date.Format(time.RFC3339Nano)
}

func NowUTCAddMinutes(minutes time.Duration) string {
	return time.Now().Add(time.Minute * minutes).Format(time.RFC3339Nano)
}
