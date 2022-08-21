package main

import (
	"encoding/json"
	"log"
	"os"
)

type SampleStruct struct {
	SampleText string `json:"sample"`
}

func main() {
	st := SampleStruct{}
	f, err := os.ReadFile("sample.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &st)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", st)

	st.SampleText = "Goodbye"

	file, err := json.MarshalIndent(st, "", " ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("sample.json", file, 0644)
	if err != nil {
		panic(err)
	}
}
