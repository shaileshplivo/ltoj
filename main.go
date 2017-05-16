package main

import (
	"bytes"
	"container/list"
	"encoding/json"
	"fmt"
	"os"
)

type DataStruct struct {
	Data int `json:"data"`
}

func listToJson(l *list.List) []byte {
	buffer := bytes.NewBufferString("[")
	length := l.Len()
	count := 0
	for e := l.Front(); e != nil; e = e.Next() {
		out, _ := json.Marshal(e.Value)
		buffer.WriteString(fmt.Sprintf("%s", string(out)))
		count++
		if count < length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return buffer.Bytes()
}

func main() {
	l := list.New()
	var d DataStruct
	d.Data = 1
	l.PushBack(d)
	d.Data = 2
	l.PushBack(d)
	d.Data = 3
	l.PushBack(d)
	d.Data = 4
	l.PushBack(d)
	d.Data = 5
	l.PushBack(d)

	jsonBytes := listToJson(l)
	println(string(jsonBytes))

	f, _ := os.Create("/tmp/dat2.json")
	defer f.Close()
	n2, _ := f.Write(jsonBytes)
	fmt.Printf("wrote %d bytes\n", n2)
}
