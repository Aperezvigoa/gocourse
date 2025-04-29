package intermediate

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type GeneralExport struct {
	ExpDate       string
	TotalProyects int
	RiskProyects  int
	Percentage    string
}

type Proyect struct {
	Name       string
	ID         string
	Priority   string
	Percentage string
}

var (
	exportTitle = template.Must(template.New("export").Parse(`Informe de Proyectos - Generado: {{.ExpDate}}
🖋️==========================================🖋️
Total Proyects: {{.TotalProyects}}
Risk Proyects: {{.RiskProyects}}
Promedy: {{.Percentage}}
`))
	exportInfo = template.Must(template.New("export").Parse(`
Details:
- Proyect "{{.Name}}" ({{.ID}})
  Priority: {{.Priority}} | {{.Percentage}}
`))
)

func main() {

	var dir string
	var output string

	flag.StringVar(&dir, "dir", "", "Directory to analyze.")
	flag.StringVar(&output, "output", "Analysis_Result.txt", "Output file with extension.")

	flag.Parse()
	if dir == "" {
		panic(errors.New("dir is required"))
	}

	dirEntry, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	} else {
		log.SetPrefix("INFO: ")
		log.SetFlags(log.Ldate | log.Ltime)
		time.Sleep(2 * time.Second)
		log.Println("Directory readed.")
	}

	fileFilter := regexp.MustCompile(`proyecto_[a-zA-Z0-9]+_[0-9]+\.txt`)
	foundedProyects := make([]string, 0)

	for _, entry := range dirEntry {
		if fileFilter.MatchString(entry.Name()) {
			foundedProyects = append(foundedProyects, entry.Name())
			log.Printf("Found proyecto: %s", entry.Name())
		}
	}

	proyects := make(map[string]map[string]string)
	pcts := make([]float64, 0)
	totalRiskProyects := 0
	for _, entry := range foundedProyects {
		path := filepath.Join(dir, entry)
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		var pct float64
		defer file.Close()
		proyects[file.Name()], pct = AnalyzeProyect(file)
		pcts = append(pcts, pct)
		if proyects[file.Name()]["risk"] == "true" {
			totalRiskProyects++
		}
	}

	totalPercentage := 0.0
	for _, pct := range pcts {
		totalPercentage += pct
	}

	export, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := export.Close()
		if err != nil {
			panic(err)
		}
		log.Println("Report exported:", output)
	}()

	header := GeneralExport{
		ExpDate:       time.Now().Format("2006-01-02"),
		TotalProyects: len(foundedProyects),
		RiskProyects:  totalRiskProyects,
		Percentage:    fmt.Sprintf("%.2f%%", totalPercentage/float64(len(foundedProyects))),
	}

	err = exportTitle.Execute(export, header)
	if err != nil {
		panic(err)
	}

	for _, entry := range proyects {
		proyect := Proyect{
			Name:       entry["nombre"],
			Priority:   entry["prioridad"],
			Percentage: entry["percentage"],
		}
		err := exportInfo.Execute(export, proyect)
		if err != nil {
			panic(err)
		}
	}
}

func AnalyzeProyect(file *os.File) (map[string]string, float64) {
	scanner := bufio.NewScanner(file)
	values := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		splitedLine := strings.Split(line, ":")
		values[strings.TrimSpace(splitedLine[0])] = strings.TrimSpace(splitedLine[1])
	}
	completedTask, err := strconv.Atoi(values["tareas_completadas"])
	if err != nil {
		panic(err)
	}
	uncompletedTask, err := strconv.Atoi(values["tareas_pendientes"])
	if err != nil {
		panic(err)
	}
	percentage := (float64(completedTask) / float64(completedTask+uncompletedTask)) * 100

	values["percentage"] = fmt.Sprintf("%.2f%%", percentage)
	values["risk"] = strconv.FormatBool(percentage <= 50)

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return values, percentage
}
