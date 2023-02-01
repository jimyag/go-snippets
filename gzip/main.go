package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"os"
	"runtime"
)

type Rule struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const (
	filename = "/Users/jimyag/code/repo/go-snippets/gzip/rule.conf"
)

func main() {
	var rules []Rule
	file, err := os.ReadFile(filename)
	Error(err)
	log.Printf("raw: %d bytes", len(file))
	err = json.Unmarshal(file, &rules)
	Error(err)
	// log.Print(rules)
	gziped, err := gzipEn(rules)
	log.Printf("len %d ,%#v", len(gziped), err)

	newRule, err := gzipDe(gziped)
	log.Printf("len %v ,%#v", newRule, err)
}

func gzipEn(data []Rule) ([]byte, error) {
	bs, err := json.Marshal(data)
	if err != nil {
		Error(err)
		return nil, err
	}
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, err = gz.Write(bs)
	if err != nil {
		Error(err)
		return nil, err
	}
	if err = gz.Flush(); err != nil {
		Error(err)
		return nil, err
	}
	if err = gz.Close(); err != nil {
		Error(err)
		return nil, err
	}
	return b.Bytes(), nil
}

func gzipDe(data []byte) ([]Rule, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		Error(err)
		return nil, err
	}
	defer reader.Close()
	bs, err := io.ReadAll(reader)
	if err != nil {
		Error(err)
		return nil, err
	}
	var res []Rule
	if err = json.Unmarshal(bs, &res); err != nil {
		Error(err)
		return nil, err
	}
	return res, nil
}

func Error(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			log.Printf("%s:%d %v", file, line, err)
		}
		log.Println(err)
	}
}
