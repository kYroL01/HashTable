package HashTable

import (
	"testing"
)

func TestNewHashMap(t *testing.T) {
	// Testing NewHashMap function
	newHT := NewHashMap()
	if newHT == nil {
		t.Error("HT creation failed!")
	}
}

func TestInsert(t *testing.T) {
	// Testing Insert function
	testCases := []struct {
		description string
		input       [2]int
		expected    [2]int
	}{
		{"input sequence #1", [2]int{1, 10}, [2]int{1, 10}},
		{"input sequence #2", [2]int{100, 33}, [2]int{100, 33}},
		{"input sequence #3", [2]int{5, 999}, [2]int{5, 999}},
	}

	HT := NewHashMap()

	for _, testC := range testCases {
		t.Run(testC.description, func(t *testing.T) {
			HT.Insert(testC.input[0], testC.input[1])
			val, b := HT.Get(testC.expected[0])
			if b == false {
				t.Errorf("Err 1 - Wrong value inserted compared to the input!")
			} else if val != testC.expected[1] {
				t.Errorf("Err 2 - Wrong value inserted compared to the input!")
			} else {
				HT.PrintHM()
			}
		})
	}
}

func TestGet(t *testing.T) {
	//TODO
}

func TestRemove(t *testing.T) {
	//TODO
}
