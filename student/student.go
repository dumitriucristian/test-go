package student

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Classes []Class `json:"classes"`
	Info    Info    `json:"info"`
}

type Class struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

type Info struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Enrolled bool   `json:"enrolled"`
}

func (s *Student) Convert(data []byte) error {

	err := json.Unmarshal(data, &s)

	if err != nil {
		fmt.Println("Error:", err)
	}

	return err
}
