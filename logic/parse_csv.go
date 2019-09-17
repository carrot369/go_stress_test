package logic

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"strings"
)

func ParseCSVFile(fileName string) [][]string {
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, err := r2.ReadAll()
	if err != nil {
		log.Fatalln("ParseCSVFile Err:", err)

		return nil
	}

	results := make([][]string, len(ss))

	for i := 0; i < len(ss); i++ {
		results[i] = ss[i]
	}

	return results
}
