package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Log struct {
	IP           string
	Date         time.Time
	HTTPMethod   string
	Path         string
	Protocol     string
	StatusCode   int
	ResponseSize int
	UserAgent    string
	Hash         string
}

func main() {

	var input string
	var output string
	var filter_status int
	var specific_ip string
	var top_ips bool

	flag.StringVar(&input, "input", "", "Access log dir")
	flag.StringVar(&output, "output", "output.txt", "Output path")
	flag.IntVar(&filter_status, "filter", 0, "Filter using HTTP status codes")
	flag.StringVar(&specific_ip, "specific-ip", "", "Filter a single IP")
	flag.BoolVar(&top_ips, "top-ips", false, "Filter top IPs")

	flag.Parse()

	if strings.TrimSpace(input) == "" {
		log.SetPrefix("FATAL ERROR: ")
		log.Fatalln("Input file can't be empty!")
		return
	}

	ipValidator := regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

	if strings.TrimSpace(specific_ip) != "" {
		valid := ipValidator.MatchString(strings.TrimSpace(specific_ip))
		if !valid {
			log.SetPrefix("FATAL ERROR: ")
			log.Fatalln("Invalid specific IP")
			return
		}
	}

	inputFile, err := os.Open(strings.TrimSpace(input))
	CheckError(err)

	defer func() {
		err = inputFile.Close()
		CheckError(err)
		log.SetPrefix("INFO: ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
		log.Printf("📦 File (%s) closed. 📦\n", inputFile.Name())
	}()

	readedLines := ReadLines(inputFile)

	RegisteredLogs := make([]Log, 0)

	for i := range readedLines {
		splitedData := SplitLogLines(readedLines[i])
		if splitedData != nil {
			RegisteredLogs = append(RegisteredLogs, RegisterLog(splitedData))
		} else {
			log.Fatalln("Something went wrong...")
		}
	}

	outputFile, err := os.Create(output)
	CheckError(err)

	defer func() {
		err := outputFile.Close()
		CheckError(err)
	}()

	if filter_status != 0 {
		for _, log := range RegisteredLogs {
			if log.StatusCode == filter_status {
				ExportFullData(outputFile, log)
			}
		}
	} else if specific_ip != "" {
		for _, log := range RegisteredLogs {
			if log.IP == specific_ip {
				ExportFullData(outputFile, log)
			}
		}
	} else {
		for _, log := range RegisteredLogs {
			ExportFullData(outputFile, log)
		}
	}
}

func ReadLines(input *os.File) []string {
	scanner := bufio.NewScanner(input)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err := scanner.Err()
	CheckError(err)

	return lines
}

func SplitLogLines(logLine string) []string {
	logRegexp := regexp.MustCompile(`^(\S+)\s+\[([^\]]+)\]\s+"(\S+)\s+([^"]+)\s+(\S+)"\s+(\d{3})\s+(\d+)\s+"([^"]+)"$`)
	return logRegexp.FindStringSubmatch(logLine)
}

func RegisterLog(splitedLine []string) Log {

	getTime, err := time.Parse("02/Jan/2006:15:04:05 -0700", splitedLine[2])
	CheckError(err)
	getStatusCode, err := strconv.Atoi(strings.TrimSpace(splitedLine[6]))
	CheckError(err)
	getResponseSize, err := strconv.Atoi(strings.TrimSpace(splitedLine[7]))
	CheckError(err)

	return Log{
		IP:           strings.TrimSpace(splitedLine[1]),
		Date:         getTime,
		HTTPMethod:   strings.TrimSpace(splitedLine[3]),
		Path:         strings.TrimSpace(splitedLine[4]),
		Protocol:     strings.TrimSpace(splitedLine[5]),
		StatusCode:   getStatusCode,
		ResponseSize: getResponseSize,
		UserAgent:    strings.TrimSpace(splitedLine[8]),
		Hash:         fmt.Sprintf("%X", sha256.Sum256([]byte(splitedLine[0]))),
	}
}

func ExportFullData(output *os.File, FLog Log) {
	exportTemplate := template.Must(template.New("full-export").Parse(`╔════════════════════════════════════════════╗
   		🔍✨  FOUND ACCESS LOG ENTRY ✨🔍  
╚════════════════════════════════════════════╝
🌐 Client IP          ┃   {{.IP}}
📅 Date & Time        ┃   {{.Date}}
⚡ HTTP Method        ┃   {{.HTTPMethod}}
📌 Requested Resource ┃   {{.Path}}
🔗 Protocol           ┃   {{.Protocol}}
⚠️ Status Code        ┃   {{.StatusCode}}
📦 Response Size      ┃   {{.ResponseSize}} bytes
🖥️ User-Agent         ┃   {{.UserAgent}}
🧬 Log Hash           ┃   {{.Hash}} (SHA256)
	
`))
	err := exportTemplate.Execute(output, FLog)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
