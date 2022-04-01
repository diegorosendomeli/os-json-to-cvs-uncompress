package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	//"fmt"
	"log"
	"os"
	"time"
)

type FruitAndVegetableRank struct {
	// 1. Create a new struct for storing read JSON objects
	Vegetable string `json:"vegetable"`
	Fruit     string `json:"fruit"`
	Rank      int64  `json:"rank"`
}

type FavoritesCompress struct {
	Item struct {
		Key struct {
			S string `json:"S"`
		} `json:"key"`
		Metadata struct {
			S string `json:"S"`
		} `json:"metadata"`
		Version struct {
			N string `json:"N"`
		} `json:"version"`
		LastUpdated struct {
			S time.Time `json:"S"`
		} `json:"last_updated"`
		DateCreated struct {
			S time.Time `json:"S"`
		} `json:"date_created"`
		CompressedValue struct {
			B string `json:"B"`
		} `json:"compressed_value"`
		DecompressedData struct {
			S DecompressedData `json:"S"`
		} `json:"decompressed_value"`
		LastUpdatedMicros struct {
			N string `json:"N"`
		} `json:"last_updated_micros"`
	} `json:"Item"`
}

type DecompressedData struct {
	UserID           string    `json:"user_id"`
	ShortcutIds      []string  `json:"shortcut_ids"`
	Preset           string    `json:"preset"`
	LastModifiedBy   string    `json:"last_modified_by"`
	LastModifiedDate time.Time `json:"last_modified_date"`
}

func convertJSONToCSV(source, destination string) error {
	// 2. Read the JSON file into the struct array
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	// remember to close the file at the end of the function
	defer sourceFile.Close()

	//var ranking []FruitAndVegetableRank
	var ranking []FavoritesCompress
	if err := json.NewDecoder(sourceFile).Decode(&ranking); err != nil {
		return err
	}

	// 3. Create a new file to store CSV data
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 4. Write the header of the CSV file and the successive rows by iterating through the JSON struct array
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	//header := []string{"key",
	//	"date_created",
	//	"last_updated",
	//	"vLastModifiedBy",
	//	"vPreset",
	//	"vUserID",
	//	"vLastModifiedDate",
	//	"vShortcutIds"}
	//if err := writer.Write(header); err != nil {
	//	return err
	//}

	count := 0
	for _, r := range ranking {
		count++

		r.Item.DecompressedData.S = stringCompressToStruct(r.Item.CompressedValue.B)

		var csvRow []string
		csvRow = append(csvRow, r.Item.Key.S,
			fmt.Sprint(r.Item.DateCreated.S)[0:19],
			fmt.Sprint(r.Item.LastUpdated.S)[0:19],
			r.Item.DecompressedData.S.LastModifiedBy,
			r.Item.DecompressedData.S.Preset,
			r.Item.DecompressedData.S.UserID,
			fmt.Sprint(r.Item.DecompressedData.S.LastModifiedDate)[0:19],
			fmt.Sprint(r.Item.DecompressedData.S.ShortcutIds))
		if err := writer.Write(csvRow); err != nil {
			log.Println(r.Item)
			return err
		}
	}
	return nil
}

func stringCompressToStruct(strCompress string) DecompressedData {
	strDescompress := stringDecompress(strCompress)
	res := DecompressedData{}
	json.Unmarshal([]byte(strDescompress), &res)
	return res
}

func main() {

	arqsJson := []string{
	"4oimjmcmeeypzn6j274scf3dxi.json",
		"4sbmkylm6uyjjcom7c4m6bmhtq.json",
		"5xexcqaqcu27zlm6ii6nzbb334.json",
		"dmlsisk2gy5sjjwdxobgmqwrnu.json",
		"ibtignisnezuhp6gmx5dmhaoda.json",
		"kmsp34kcsa7irp7vhh3gphwcty.json",
		"lygi4nq3ne4tzehbyrtlgvnxru.json",
		"rfqz3zndnm3e7cvnqoadopnovq.json",
		"v3jcesgslu57dlgml3g3znnb4a.json",
		"zjfmqu7ooy35piq4fwdxz47epi.json",
	}

	var arqCsv string

	for _, a := range arqsJson {
		arqCsv = strings.ReplaceAll(a, ".json", "_sem_header.csv")
		log.Println("início - de: " + a + " para: " + arqCsv)
		if err := convertJSONToCSV(a, arqCsv); err != nil {
			log.Fatal(err)
		}
		log.Println("fim - de: " + a + " para: " + arqCsv)
	}





}

func stringDecompress(str string) string {
	data, _ := base64.StdEncoding.DecodeString(str)
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	return string(s)
}
