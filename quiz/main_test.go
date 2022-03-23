package main

import (
	"reflect"
	"strings"
	"testing"
)

var nilSlice [][]string

var parseCSVTestdata = []struct {
	testName string
	filepath string
	expected [][]string
	err      string
}{
	{
		testName: "Positive Case",
		filepath: "test_problems.csv",
		expected: [][]string{{"5+5", "10"}, {"7+3", "10"}},
		err:      "",
	},
	{
		testName: "Non Existent File",
		filepath: "test.csv",
		expected: nilSlice,
		err:      "",
	},
}

var createQuizListTestData = []struct {
	testName string
	data     [][]string
	expected []Quiz
}{
	{
		testName: "Positive Case",
		data:     [][]string{{"5+5", "10"}, {"what2+2, sir?", "4"}, {"7+3", "10"}},
		expected: []Quiz{
			{
				question: "5+5",
				answer:   "10",
			},
			{
				question: "what2+2, sir?",
				answer:   "4",
			},
			{
				question: "7+3",
				answer:   "10",
			},
		},
	},
	{
		testName: "Space in CSV",
		data:     [][]string{{"5+5", "  10 "}, {"what2+2, sir?", " 4"}, {"7+3", "10"}},
		expected: []Quiz{
			{
				question: "5+5",
				answer:   "10",
			},
			{
				question: "what2+2, sir?",
				answer:   "4",
			},
			{
				question: "7+3",
				answer:   "10",
			},
		},
	},
}

func TestParseCSV(t *testing.T) {
	for _, test := range parseCSVTestdata {
		t.Run(test.testName, func(t *testing.T) {
			got, err := ParseCSV(test.filepath)
			if err != nil {
				if !strings.Contains(err.Error(), test.err) {
					t.Errorf("Expected %v Got %v", test.expected, got)
				}
			}
			if !reflect.DeepEqual(test.expected, got) {
				t.Errorf("Expected %v Got %v", test.expected, got)
			}
		})
	}
}

func TestCreateQuizList(t *testing.T) {

	for _, test := range createQuizListTestData {
		t.Run(test.testName, func(t *testing.T) {
			got := CreateQuizList(test.data)
			if !reflect.DeepEqual(test.expected, got) {
				t.Errorf("Expected %v Got %v", test.expected, got)
			}
		})
	}
}
