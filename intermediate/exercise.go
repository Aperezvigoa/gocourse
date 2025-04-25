package intermediate

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"strings"
	"time"
)

type LogDetection interface {
	DetectSuspicious() bool
	PrintSuspicious()
}
type LogEntry struct {
	Date         time.Time
	IP           string
	Method       string
	Path         string
	Code         int
	ResponseTime string
	Suspicious   bool
}

func (e *LogEntry) DetectSuspicious() bool {
	if e.Code >= 400 {
		e.Suspicious = true
	}
	fmt.Println("%s Detected suspicious status", e.IP)
	return e.Suspicious
}

func (e *LogEntry) PrintSuspicious() {
	templ := template.Must(template.New("suspicious").Parse(`
***************************************************
- Date {{.Date}}
- IP {{.IP}}
- Method {{.Method}}
`))
	err := templ.Execute(os.Stdout, e)
	if err != nil {
		fmt.Println(err)
	}
	logHash := sha256.Sum256([]byte(e.IP))
	fmt.Printf("- Log Hash: %x", logHash)
}

func main() {
	// --- Open File
	file, err := os.Open("exercise.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// --- Defer Closing File
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			return
		}
		fmt.Println("File closed", file.Name())
	}()

	gross, err := GetStringLogs(file)
	if err != nil {
		fmt.Println("Error getting logs:", err)
		return
	}

	// LOGS SLICE
	totalLogs := BuildLogs(gross)

	// GETTING SUSPICIOUS LOGS
	suspiciousLogs := make([]LogEntry, 0)
	for _, l := range totalLogs {
		if l.DetectSuspicious() {
			suspiciousLogs = append(suspiciousLogs, l)
		}
	}
	// PRINTING SUSPICIOUS REPORT
	for _, l := range suspiciousLogs {
		l.PrintSuspicious()
	}
}

func GetStringLogs(file *os.File) ([]string, error) {
	scanner := bufio.NewScanner(file)
	var builder strings.Builder
	gross := make([]string, 0)

	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		gross = append(gross, builder.String())
		builder.Reset()
	}
	if scanner.Err() != nil {
		fmt.Println("Error scanning file:", scanner.Err())
		return nil, scanner.Err()
	}
	return gross, nil
}

func BuildLogs(gross []string) []LogEntry {
	cleaned := make([]LogEntry, 0)
	var entry LogEntry
	for _, v := range gross {
		actualGrossLog := strings.Split(v, "|")
		for i, v := range actualGrossLog {
			switch i {
			case 0:
				entry.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST m=+0.000000000", strings.TrimSpace(v))
			case 1:
				entry.IP = strings.TrimSpace(v)
			case 2:
				splitedPathMethod := strings.Split(strings.TrimSpace(v), " ")
				entry.Method = strings.TrimSpace(splitedPathMethod[0])
				entry.Path = strings.TrimSpace(splitedPathMethod[1])
			case 3:
				entry.Code, _ = strconv.Atoi(strings.TrimSpace(v))
			case 4:
				entry.ResponseTime = strings.TrimSpace(v)
				cleaned = append(cleaned, entry)
			}
		}
	}
	return cleaned
}
