package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type DateTime struct {
	time.Time
}

func (cd *DateTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]

	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

type Columns struct {
	Id         int      `json:"id"`
	Company    string   `json:"company"`
	Model      string   `json:"model"`
	Product    string   `json:"product"`
	Partnumber string   `json:"partnumber"`
	CreateDate DateTime `json:"createdate"`
}

func main() {
	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var columns []Columns
	err = json.Unmarshal(byteValue, &columns)

	if err != nil {
		fmt.Println(err)
		return
	}

	csvFile, err := os.Create("data.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	header := []string{"Id", "Company", "Model", "Product", "Partnumber", "CreateDate"}
	writer.Write(header)

	for _, column := range columns {
		line := []string{
			fmt.Sprintf("%d", column.Id),
			column.Company,
			column.Model,
			column.Product,
			column.Partnumber,
			column.CreateDate.Format("2006-01-02"),
		}
		writer.Write(line)
	}

	fmt.Println("CSV file created with success!")
}
