package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("13.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	f, err := os.Create("haha2.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		//fmt.Println(record[0]) // record has the type []string
		str := strings.Split(record[0], "\t")
		w.Write(str)
	}
	w.Flush()
}
