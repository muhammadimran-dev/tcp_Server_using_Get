package csvdata

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

//Covid Struct
type Covid struct {
	TestPositive  string `json:"TestPositive"`
	TestPerformed string `json:"TestPerformed"`
	Date          string `json:"Date"`
	Discharged    string `json:"Discharged"`
	Expired       string `json:"Expired"`
	StillAdmitted string `json:"StillAdmitted"`
	Region        string `json:"Region"`
}

//Load Function
func Load(path string) []Covid {
	table := make([]Covid, 0)
	csvFile, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err.Error())
		}
		c := Covid{
			TestPositive:  row[0],
			TestPerformed: row[1],
			Date:          row[2],
			Discharged:    row[3],
			Expired:       row[4],
			Region:        row[5],
			StillAdmitted: row[6],
		}
		table = append(table, c)

	}
	return table

}

//Find Function to search data for the given query
func Find(table []Covid, filter string) []Covid {
	if filter == "" || filter == "*" {
		return table
	}

	result := make([]Covid, 0)
	filter = strings.ToUpper(filter)
	for _, cov := range table {
		if cov.Region == filter ||
			cov.Date == filter ||
			strings.Contains(strings.ToUpper(cov.Region), filter) ||
			strings.Contains(strings.ToUpper(cov.Date), filter) {
			result = append(result, cov)
		}
	}

	return result
}
