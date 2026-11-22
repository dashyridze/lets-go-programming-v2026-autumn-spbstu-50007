package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int  `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	ValueStr string  `xml:"Value"`
	Value    float64 `xml:"-"`
}

type OutputValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to the configuration file")
	flag.Parse()

	configData, err := os.ReadFile(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file: %v", err))
	}
	var cfg Config
	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse YAML: %v", err))
	}
	xmlData, err := os.ReadFile(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to read XML file: %v", err))
	}
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}
	var valCurs ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse XML: %v", err))
	}
	for i := range valCurs.Valutes {
		cleanValue := strings.Replace(valCurs.Valutes[i].ValueStr, ",", ".", -1)
		val, err := strconv.ParseFloat(cleanValue, 64)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse value %q: %v", valCurs.Valutes[i].ValueStr, err))
		}
		valCurs.Valutes[i].Value = val
	}
	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})
	result := make([]OutputValute, 0, len(valCurs.Valutes))
	for _, v := range valCurs.Valutes {
		result = append(result, OutputValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    v.Value,
		})
	}
	dir := filepath.Dir(cfg.OutputFile)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		panic(fmt.Sprintf("Failed to create directory: %v", err))
	}
	jsonData, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		panic(fmt.Sprintf("Failed to encode to JSON: %v", err))
	}
	err = os.WriteFile(cfg.OutputFile, jsonData, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to write output file: %v", err))
	}
	fmt.Println("File saved to:", cfg.OutputFile)
}
