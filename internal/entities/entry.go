package entities

import (
	"encoding/json"
	"strings"
)

type StringList []string

func (list *StringList) UnmarshalCSV(csv string) (err error) {
	*list = strings.Split(csv, ",")
	return nil
}

func (list *StringList) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*list = strings.Split(value, ",")
	return nil
}

type Entry struct {
	Title      string     `json:"title" csv:"title"`
	Inflexions StringList `json:"inflexions" csv:"inflexions"`
	Grammar1   string     `json:"grammar_1" csv:"grammar_1"`
	Meaning1   string     `json:"meaning_1" csv:"meaning_1"`
	Example1   string     `json:"example_1" csv:"example_1"`
	Grammar2   string     `json:"grammar_2" csv:"grammar_2"`
	Meaning2   string     `json:"meaning_2" csv:"meaning_2"`
	Example2   string     `json:"example_2" csv:"example_2"`
	Grammar3   string     `json:"grammar_3" csv:"grammar_3"`
	Meaning3   string     `json:"meaning_3" csv:"meaning_3"`
	Example3   string     `json:"example_3" csv:"example_3"`
	Grammar4   string     `json:"grammar_4" csv:"grammar_4"`
	Meaning4   string     `json:"meaning_4" csv:"meaning_4"`
	Example4   string     `json:"example_4" csv:"example_4"`
	Note1      string     `json:"note_1" csv:"note_1"`
	Note2      string     `json:"note_2" csv:"note_2"`
	Note3      string     `json:"note_3" csv:"note_3"`
	Note4      string     `json:"note_4" csv:"note_4"`
}
