package utils

import (
	"reflect"
	"testing"
)

func TestChunksWithSizeParamLessItemsLength(t *testing.T) {
	t.Parallel()

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}

	if actual := Chunks(items, 3); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("\nActual:%v\nExpect:%v", actual, expected)
	}
}

func TestChunksWithItemsLengthMultipleSizeParam(t *testing.T) {
	t.Parallel()

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	if actual := Chunks(items, 3); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("\nActual:%v\nExpect:%v", actual, expected)
	}
}

func TestChunksWithSizeParamGreaterItemsLength(t *testing.T) {
	t.Parallel()

	items := []int{1, 2, 3, 4, 5}
	expected := [][]int{{1, 2, 3, 4, 5}}

	if actual := Chunks(items, 10); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("\nActual:%v\nExpect:%v", actual, expected)
	}
}

func TestChunksWithSizeParamEqualOne(t *testing.T) {
	t.Parallel()

	items := []int{1, 2, 3}
	expected := [][]int{{1}, {2}, {3}}

	if actual := Chunks(items, 1); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("\nActual:%v\nExpect:%v", actual, expected)
	}
}

func TestChunksWithEmptyItems(t *testing.T) {
	t.Parallel()

	items := []int{}
	expected := [][]int{{}}

	if actual := Chunks(items, 3); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("\nActual:%v\nExpect:%v", actual, expected)
	}
}

func TestChunksWithNilItems(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Chunks did not panic")
		}
	}()

	var items []int

	Chunks(items, 3)
}

func TestChunksWithNotPositiveSizeParam(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Chunks did not panic")
		}
	}()

	items := []int{1, 2, 3, 4, 5}

	Chunks(items, 0)
}

func TestInvertMapWithUniqueValues(t *testing.T) {
	t.Parallel()

	data := map[string]int{
		"go":     1,
		"ozon":   2021,
		"school": 3,
	}
	expected := map[int]string{
		1:    "go",
		2021: "ozon",
		3:    "school",
	}

	if actual := InvertMap(data); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("\nActual:%v\nExpect:%v", actual, expected)
	}
}

func TestInvertMapEmptyData(t *testing.T) {
	t.Parallel()

	data := map[string]int{}
	expected := map[int]string{}

	if actual := InvertMap(data); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("\nActual:%v\nExpect:%v", actual, expected)
	}
}

func TestInvertMapWithNotUniqueValues(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("InvertMap did not panic")
		}
	}()

	data := map[string]int{
		"go":     1,
		"ozon":   2021,
		"school": 1,
	}

	InvertMap(data)
}

func TestFilterStopWords(t *testing.T) {
	t.Parallel()

	words := []string{"Keanu", "Timati"}
	expected := []string{"Keanu"}

	if actual := FilterStopWords(words); !reflect.DeepEqual(actual, expected) {
		t.Fatalf("\nActual:%v\nExpect:%v", actual, expected)
	}
}
