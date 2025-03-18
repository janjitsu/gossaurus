package entities

import "strings"

type StringList []string

func (list *StringList) UnmarshalCSV(csv string) (err error) {
	*list = strings.Split(csv, ",")
	return nil
}

type Entry struct {
	Title      string     `csv:"title"`
	Inflexions StringList `csv:"inflexions"`
	Grammar1   string     `csv:"grammar_1"`
	Meaning1   string     `csv:"meaning_1"`
	Example1   string     `csv:"example_1"`
	Grammar2   string     `csv:"grammar_2"`
	Meaning2   string     `csv:"meaning_2"`
	Example2   string     `csv:"example_2"`
	Grammar3   string     `csv:"grammar_3"`
	Meaning3   string     `csv:"meaning_3"`
	Example3   string     `csv:"example_3"`
	Grammar4   string     `csv:"grammar_4"`
	Meaning4   string     `csv:"meaning_4"`
	Example4   string     `csv:"example_4"`
	Note1      string     `csv:"note_1"`
	Note2      string     `csv:"note_2"`
	Note3      string     `csv:"note_3"`
	Note4      string     `csv:"note_4"`
}
