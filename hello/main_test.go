package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type LogSystem struct {
	Field1 string      `json:"field1,omitempty"`
	Field2 string      `json:"field2,omitempty"`
	Field3 interface{} `json:"field3,omitempty"`
}

type LogCondition struct {
	Program []LogSubField `json:"program,omitempty"`
	System  []LogSubField `json:"system,omitempty"`
	System2 []LogSubField `json:"system2,omitempty"`
}

type LogSubField struct {
	Name   string `json:"name,omitempty"`
	Cond   string `json:"cond,omitempty"`
	Nocase string `json:"nocase,omitempty"`
}

func TestLogQuery(t *testing.T) {
	jsonvalue := ` {
		"field1" : "test1",
		"field2" : "test2",
		"field3" : {
			"program" : [
				{
					"name":"name1",
					"cond":"cond1",
					"nocase":"yes"
				},
				{
					"name":"name2",
					"cond":"cond2",
					"nocase":"no"					
				}
			],
			"system" : [
				{
					"name":"system1",
					"cond":"cond1",
					"nocase":"yes"
				},
				{
					"name":"system2",
					"cond":"cond2",
					"nocase":"no"					
				}
			]
		}
	}`

	log3 := &LogSystem{}
	log3.Field3 = &LogCondition{}

	json.Unmarshal([]byte(jsonvalue), &log3)
	fmt.Println(jsonvalue)
	fmt.Println(log3.Field3)

	filter := log3.Field3.(*LogCondition)
	fmt.Println(filter.Program[0].Name)
	fmt.Println(filter.System[0].Name)
	if filter.System2 != nil {
		fmt.Println(filter.System2[0].Name)
	}
}
