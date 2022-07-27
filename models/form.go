package models

import (
	"encoding/json"
	"log"
	"strings"
)

type Form struct {
	Name  string `form:"name", binding:"required"`
	Email string `form:"email", binding:"required"`
	Phone string `form:"phone", binding:"required"`
}

func (f Form) GetData() [][]string {
	b, err := json.Marshal(f)
	if err != nil {
		log.Fatalln(err)
	}
	s := strings.NewReplacer("{", "", "}", "", "\"", "").Replace(string(b))
	l := strings.Split(s, ",")
	var listString [][]string
	for _, v := range l {
		splitList := strings.Split(v, ":")
		listString = append(listString, splitList)
	}
	return listString
}
