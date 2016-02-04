// Copyright (c) 2016 Bob Ziuchkovski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package haven

import (
	"reflect"
	"testing"
)

/*
 * Functions we implement -- these should be tested thoroughly
 */

func TestAbs(t *testing.T) {
	var tests = []struct {
		Operand  int
		Expected int
	}{
		{Operand: -10, Expected: 10},
		{Operand: -1, Expected: 1},
		{Operand: 0, Expected: 0},
		{Operand: -1, Expected: 1},
		{Operand: 10, Expected: 10},
	}

	for _, test := range tests {
		result := Abs(test.Operand)
		if result != test.Expected {
			t.Errorf("Abs result incorrect.  Operand: %d, Expected: %d, Received: %d", test.Operand, test.Expected, result)
		}
	}
}

func TestGrep(t *testing.T) {
	var tests = []struct {
		Operand  []string
		Pattern  string
		Expected []string
		Valid    bool
	}{
		{Operand: []string{"dog", "cat"}, Pattern: "dog", Expected: []string{"dog"}, Valid: true},
		{Operand: []string{"dog", "cat", "horse"}, Pattern: "^?o", Expected: []string{"dog", "horse"}, Valid: true},
		{Operand: []string{"dog", "cat", "horse"}, Pattern: "zzz", Expected: nil, Valid: true},
		{Operand: []string{"dog", "cat", "horse"}, Pattern: "[[:bogus:]]", Valid: false},
	}

	for _, test := range tests {
		result, err := Grep(test.Pattern, test.Operand)
		if test.Valid && err != nil {
			t.Errorf("Grep encountered unexpected error: %s.  Operand: %#v, Pattern: %s", err, test.Operand, test.Pattern)
		}
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Grep result incorrect.  Operand: %#v, Pattern: %s, Expected: %#v, Received: %#v", test.Operand, test.Pattern, test.Expected, result)
		}
	}
}

func TestHead(t *testing.T) {
	var tests = []struct {
		Operand  []string
		N        int
		Expected []string
	}{
		{Operand: []string{"dog", "cat", "horse"}, N: 2, Expected: []string{"dog", "cat"}},
		{Operand: []string{"dog", "cat", "horse"}, N: 30, Expected: []string{"dog", "cat", "horse"}},
	}

	for _, test := range tests {
		result := Head(test.N, test.Operand)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Head result incorrect.  Operand: %#v, N: %d, Expected: %#v, Received: %#v", test.Operand, test.N, test.Expected, result)
		}
	}
}

func TestIntersect(t *testing.T) {
	var tests = []struct {
		Operand  []string
		A        []string
		Expected []string
	}{
		{Operand: []string{"dog", "cat", "horse"}, A: []string{"horse", "dog"}, Expected: []string{"dog", "horse"}},
		{Operand: []string{"dog", "dog", "cat", "horse"}, A: []string{"horse", "dog"}, Expected: []string{"dog", "horse"}},
		{Operand: []string{"dog"}, A: []string{"horse"}, Expected: []string{}},
		{Operand: []string{"dog"}, A: []string(nil), Expected: []string{}},
		{Operand: []string(nil), A: []string{"dog"}, Expected: []string{}},
	}

	for _, test := range tests {
		result := Intersect(test.A, test.Operand)
		if len(result) != len(test.Expected) {
			t.Errorf("Intersect result incorrect.  Operand: %#v, A: %#v, Expected: %#v, Received: %#v", test.Operand, test.A, test.Expected, result)
		}
		for _, expectedElem := range test.Expected {
			foundElem := false
			for _, resultElem := range result {
				if resultElem == expectedElem {
					foundElem = true
					break
				}
			}
			if !foundElem {
				t.Errorf("Intersect result incorrect.  Operand: %#v, A: %#v, Expected: %#v, Received: %#v", test.Operand, test.A, test.Expected, result)
			}
		}
	}
}

func TestLines(t *testing.T) {
	var tests = []struct {
		Operand  string
		Expected []string
	}{
		{Operand: "a\nb\nc\n", Expected: []string{"a", "b", "c"}},
		{Operand: "a few\nwords  \n  on each line\n", Expected: []string{"a few", "words  ", "  on each line"}},
		{Operand: "single line", Expected: []string{"single line"}},
		{Operand: "empty\n\nlines", Expected: []string{"empty", "", "lines"}},
		{Operand: "trailing\n", Expected: []string{"trailing"}},
		{Operand: "\nleading", Expected: []string{"", "leading"}},
		{Operand: "windows\r\nlines", Expected: []string{"windows", "lines"}},
	}

	for _, test := range tests {
		result, _ := Lines(test.Operand)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Lines result incorrect.  Operand: %#v, Expected: %#v, Received: %#v", test.Operand, test.Expected, result)
		}
	}
}

func TestMax(t *testing.T) {
	var tests = []struct {
		Operand  int
		A        int
		Expected int
	}{
		{Operand: -10, A: -10, Expected: -10},
		{Operand: -10, A: -1, Expected: -1},
		{Operand: -10, A: 0, Expected: 0},
		{Operand: -10, A: 1, Expected: 1},
		{Operand: -10, A: 10, Expected: 10},
		{Operand: 0, A: -10, Expected: 0},
		{Operand: 0, A: -1, Expected: 0},
		{Operand: 0, A: 0, Expected: 0},
		{Operand: 0, A: 1, Expected: 1},
		{Operand: 0, A: 10, Expected: 10},
		{Operand: 10, A: -10, Expected: 10},
		{Operand: 10, A: -1, Expected: 10},
		{Operand: 10, A: 0, Expected: 10},
		{Operand: 10, A: 1, Expected: 10},
		{Operand: 10, A: 10, Expected: 10},
	}

	for _, test := range tests {
		result := Max(test.A, test.Operand)
		if result != test.Expected {
			t.Errorf("Max result incorrect.  Operand: %d, A: %d, Expected: %d, Received: %d", test.Operand, test.A, test.Expected, result)
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		Operand  int
		A        int
		Expected int
	}{
		{Operand: -10, A: -10, Expected: -10},
		{Operand: -10, A: -1, Expected: -10},
		{Operand: -10, A: 0, Expected: -10},
		{Operand: -10, A: 1, Expected: -10},
		{Operand: -10, A: 10, Expected: -10},
		{Operand: 0, A: -10, Expected: -10},
		{Operand: 0, A: -1, Expected: -1},
		{Operand: 0, A: 0, Expected: 0},
		{Operand: 0, A: 1, Expected: 0},
		{Operand: 0, A: 10, Expected: 0},
		{Operand: 10, A: -10, Expected: -10},
		{Operand: 10, A: -1, Expected: -1},
		{Operand: 10, A: 0, Expected: 0},
		{Operand: 10, A: 1, Expected: 1},
		{Operand: 10, A: 10, Expected: 10},
	}

	for _, test := range tests {
		result := Min(test.A, test.Operand)
		if result != test.Expected {
			t.Errorf("Min result incorrect.  Operand: %d, A: %d, Expected: %d, Received: %d", test.Operand, test.A, test.Expected, result)
		}
	}

}

func TestReverse(t *testing.T) {
	var tests = []struct {
		Operand  []string
		Expected []string
	}{
		{Operand: []string{"dog", "cat", "horse"}, Expected: []string{"horse", "cat", "dog"}},
		{Operand: []string{"dog"}, Expected: []string{"dog"}},
		{Operand: []string{}, Expected: []string{}},
	}

	for _, test := range tests {
		result := Reverse(test.Operand)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Reverse result incorrect.  Operand: %#v, Expected: %#v, Received: %#v", test.Operand, test.Expected, result)
		}
	}
}

func TestSeq(t *testing.T) {
	var tests = []struct {
		First    int
		Last     int
		Incr     []int
		Expected []int
	}{
		{First: 0, Last: 5, Incr: nil, Expected: []int{0, 1, 2, 3, 4, 5}},
		{First: 3, Last: 12, Incr: []int{3}, Expected: []int{3, 6, 9, 12}},
		{First: 3, Last: 13, Incr: []int{3}, Expected: []int{3, 6, 9, 12}},
		{First: -1, Last: 3, Incr: []int{2}, Expected: []int{-1, 1, 3}},
		{First: -1, Last: 3, Expected: []int{-1, 0, 1, 2, 3}},
		{First: 4, Last: 0, Expected: nil},
		{First: 4, Last: 0, Incr: []int{-1}, Expected: []int{4, 3, 2, 1, 0}},
	}

	for _, test := range tests {
		result := Seq(test.First, test.Last, test.Incr...)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Reverse result incorrect.  First: %d, Last: %d, Incr: %#v, Expected: %#v, Received: %#v", test.First, test.Last, test.Incr, test.Expected, result)
		}
	}
}

func TestShuffle(t *testing.T) {
	var tests = []struct {
		Operand []string
	}{
		{Operand: []string{"dog", "cat", "horse"}},
	}
	for _, test := range tests {
		result := Shuffle(test.Operand)
		if len(result) != len(test.Operand) {
			t.Errorf("Shuffle produced slice with bad length.  Operand: %#v, Expected Length: %d, Reveived: %#v", test.Operand, len(test.Operand), result)
		}
		for _, expectedElem := range test.Operand {
			foundElem := false
			for _, resultElem := range result {
				if resultElem == expectedElem {
					foundElem = true
					break
				}
			}
			if !foundElem {
				t.Errorf("Shuffle produced result not in original slice.  Operand: %#v, Received: %#v", test.Operand, result)
			}
		}
	}
}

func TestSort(t *testing.T) {
	var tests = []struct {
		Operand  []string
		Expected []string
	}{
		{Operand: []string{"dog", "cat", "horse"}, Expected: []string{"cat", "dog", "horse"}},
		{Operand: []string{"cat"}, Expected: []string{"cat"}},
		{Operand: []string{}, Expected: []string{}},
	}
	for _, test := range tests {
		result := Sort(test.Operand)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Sort result incorrect.  Operand: %#v, Expected: %#v, Received: %#v", test.Operand, test.Expected, result)
		}
	}
}

func TestTail(t *testing.T) {
	var tests = []struct {
		Operand  []string
		N        int
		Expected []string
	}{
		{Operand: []string{"dog", "cat", "horse"}, N: 2, Expected: []string{"cat", "horse"}},
		{Operand: []string{"dog", "cat", "horse"}, N: 30, Expected: []string{"dog", "cat", "horse"}},
	}

	for _, test := range tests {
		result := Tail(test.N, test.Operand)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Tail result incorrect.  Operand: %#v, N: %d, Expected: %#v, Received: %#v", test.Operand, test.N, test.Expected, result)
		}
	}
}

func TestUnion(t *testing.T) {
	var tests = []struct {
		Operand  []string
		A        []string
		Expected []string
	}{
		{Operand: []string{"dog", "cat"}, A: []string{"horse", "dog"}, Expected: []string{"cat", "dog", "horse"}},
		{Operand: []string{"dog"}, A: []string{"horse"}, Expected: []string{"dog", "horse"}},
		{Operand: []string{"dog", "dog", "cat"}, A: []string{"horse"}, Expected: []string{"cat", "dog", "horse"}},
		{Operand: []string{"dog"}, A: []string{}, Expected: []string{"dog"}},
		{Operand: []string{"dog"}, A: []string(nil), Expected: []string{"dog"}},
		{Operand: []string{}, A: []string{"dog"}, Expected: []string{"dog"}},
		{Operand: []string(nil), A: []string{"dog"}, Expected: []string{"dog"}},
		{Operand: []string{}, A: []string{}, Expected: []string{}},
	}

	for _, test := range tests {
		result := Union(test.A, test.Operand)
		if len(result) != len(test.Expected) {
			t.Errorf("Union result incorrect.  Operand: %#v, A: %#v, Expected: %#v, Received: %#v", test.Operand, test.A, test.Expected, result)
		}
		for _, expectedElem := range test.Expected {
			foundElem := false
			for _, resultElem := range result {
				if resultElem == expectedElem {
					foundElem = true
					break
				}
			}
			if !foundElem {
				t.Errorf("Union result incorrect.  Operand: %#v, A: %#v, Expected: %#v, Received: %#v", test.Operand, test.A, test.Expected, result)
			}
		}
	}
}

/*
 * Functions we wrap -- these should have a basic test to ensure they run, but
 * their logic is already tested upstream, so there's no sense in beating on them.
 */

func TestAdd(t *testing.T) {
	operand, a, expected := 42, 3, 45
	result := Add(a, operand)
	if result != expected {
		t.Errorf("Add result incorrect.  Operand: %d, A: %d, Expected: %d, Received: %d", operand, a, expected, result)
	}
}

func TestBase64Decode(t *testing.T) {
	operand, expected := "dGVzdCB0ZXh0", "test text"
	result, err := Base64Decode(operand)
	if err != nil {
		t.Errorf("Base64Decode encountered unexpected error: %s.  Operand: %s", err, operand)
	}
	if result != expected {
		t.Errorf("Base64Decode result incorrect.  Operand: %s, Expected: %s, Received: %s", operand, expected, result)
	}
}

func TestBase64Encode(t *testing.T) {
	operand, expected := "test text", "dGVzdCB0ZXh0"
	result := Base64Encode(operand)
	if result != expected {
		t.Errorf("Base64Encode result incorrect.  Operand: %s, Expected: %s, Received: %s", operand, expected, result)
	}
}

func TestCompileERE(t *testing.T) {
	_, err := CompileERE(".*test.*")
	if err != nil {
		t.Errorf("CompileERE encountered unexpected error: %s", err)
	}
}

func TestCompileRegex(t *testing.T) {
	_, err := CompileRegex(".*test.*")
	if err != nil {
		t.Errorf("CompileRegex encountered unexpected error: %s", err)
	}
}

func TestContains(t *testing.T) {
	operand, substr, expected := "cat dog horse", "dog", true
	result := Contains(substr, operand)
	if result != expected {
		t.Errorf("Contains result incorrect.  Operand: %s, Substr: %s, Expected: %t, Received: %t", operand, substr, expected, result)
	}
}

func TestContainsAny(t *testing.T) {
	operand, chars, expected := "dog=horse", ".=", true
	result := ContainsAny(chars, operand)
	if result != expected {
		t.Errorf("ContainsAny result incorrect.  Operand: %s, Chars: %s, Expected: %t, Received: %t", operand, chars, expected, result)
	}
}

func TestCount(t *testing.T) {
	operand, str, expected := "cat=dog=horse", "=", 2
	result := Count(str, operand)
	if result != expected {
		t.Errorf("Count result incorrect.  Operand: %s, Str: %s, Expected: %d, Received: %d", operand, str, expected, result)
	}
}

func TestDivide(t *testing.T) {
	operand, a, expected := 42, 2, 21
	result := Divide(a, operand)
	if result != expected {
		t.Errorf("Divide result incorrect.  Operand: %d, A: %d, Expected: %d, Received: %d", operand, a, expected, result)
	}
}

func TestFields(t *testing.T) {
	operand, expected := "cat dog horse", []string{"cat", "dog", "horse"}
	result := Fields(operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Fields result incorrect.  Operand: %s, Expected: %#v, Received: %#v", operand, expected, result)
	}
}

func TestHasPrefix(t *testing.T) {
	operand, prefix, expected := "cat dog horse", "cat", true
	result := HasPrefix(prefix, operand)
	if result != expected {
		t.Errorf("HasPrefix result incorrect.  Operand: %s, Prefix: %s, Expected: %t, Received: %t", operand, prefix, expected, result)
	}
}

func TestHasSuffix(t *testing.T) {
	operand, suffix, expected := "cat dog horse", "horse", true
	result := HasSuffix(suffix, operand)
	if result != expected {
		t.Errorf("HasSuffix result incorrect.  Operand: %s, Suffix: %s, Expected: %t, Received: %t", operand, suffix, expected, result)
	}
}

func TestIndex(t *testing.T) {
	operand, substr, expected := "cat dog horse", "dog", 4
	result := Index(substr, operand)
	if result != expected {
		t.Errorf("Index result incorrect.  Operand: %s, substr: %s, Expected: %d, Received: %d", operand, substr, expected, result)
	}
}

func TestIndexAny(t *testing.T) {
	operand, chars, expected := "cat dog horse", ".=d", 4
	result := IndexAny(chars, operand)
	if result != expected {
		t.Errorf("IndexAny result incorrect.  Operand: %s, Chars: %s, Expected: %d, Received: %d", operand, chars, expected, result)
	}
}

func TestJoin(t *testing.T) {
	operand, sep, expected := []string{"cat", "dog", "horse"}, ",", "cat,dog,horse"
	result := Join(sep, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Join result incorrect.  Operand: %#v, Sep: %s, Expected: %s, Received: %s", operand, sep, expected, result)
	}
}

func TestLastIndex(t *testing.T) {
	operand, substr, expected := "cat dog horse", "o", 9
	result := LastIndex(substr, operand)
	if result != expected {
		t.Errorf("LastIndex result incorrect.  Operand: %s, Substr: %s, Expected: %d, Received: %d", operand, substr, expected, result)
	}
}

func TestLastIndexAny(t *testing.T) {
	operand, chars, expected := "cat dog horse", "oe", 12
	result := LastIndexAny(chars, operand)
	if result != expected {
		t.Errorf("LastIndexAny result incorrect.  Operand: %s, Chars: %s, Expected: %d, Received: %d", operand, chars, expected, result)
	}
}

func TestMatches(t *testing.T) {
	operand, pattern, expected := "cat dog horse", ".*dog.*", true
	result, err := Matches(pattern, operand)
	if err != nil {
		t.Errorf("Matches encountered unexpected error: %s.  Operand: %s", err, operand)
	}
	if result != expected {
		t.Errorf("Matches result incorrect.  Operand: %s, Pattern: %s, Expected: %t, Received: %t", operand, pattern, expected, result)
	}
}

func TestModulo(t *testing.T) {
	operand, a, expected := 42, 10, 2
	result := Modulo(a, operand)
	if result != expected {
		t.Errorf("Modulo result incorrect.  Operand: %d, A: %d, Expected: %d, Received: %d", operand, a, expected, result)
	}
}

func TestMultiply(t *testing.T) {
	operand, a, expected := 42, 10, 420
	result := Multiply(a, operand)
	if result != expected {
		t.Errorf("Multiply result incorrect.  Operand: %d, A: %d, Expected: %d, Received: %d", operand, a, expected, result)
	}
}

func TestNow(t *testing.T) {
	Now()
}

func TestParseBool(t *testing.T) {
	operand, expected := "false", false
	result, err := ParseBool(operand)
	if err != nil {
		t.Errorf("ParseBool encountered unexpected error: %s.  Operand: %s", err, operand)
	}
	if result != expected {
		t.Errorf("ParseBool result incorrect.  Operand: %s, Expected: %t, Received: %t", operand, expected, result)
	}
}

func TestParseFloat(t *testing.T) {
	operand, expected := "2.0", float64(2.0)
	result, err := ParseFloat(operand)
	if err != nil {
		t.Errorf("ParseFloat encountered unexpected error: %s.  Operand: %s", err, operand)
	}
	if result != expected {
		t.Errorf("ParseFloat result incorrect.  Operand: %s, Expected: %f, Received: %f", operand, expected, result)
	}
}

func TestParseInt(t *testing.T) {
	operand, expected := "42", 42
	result, err := ParseInt(operand)
	if err != nil {
		t.Errorf("ParseInt encountered unexpected error: %s.  Operand: %s", err, operand)
	}
	if result != expected {
		t.Errorf("ParseInt result incorrect.  Operand: %s, Expected: %d, Received: %d", operand, expected, result)
	}
}

func TestParseTime(t *testing.T) {
	operand, format, day, year := "2013-Feb-03", "2006-Jan-02", 3, 2013
	result, err := ParseTime(format, operand)
	if err != nil {
		t.Errorf("ParseTime encountered unexpected error: %s.  Operand: %s", err, operand)
	}
	if result.Year() != year {
		t.Errorf("ParseTime year incorrect.  Operand: %s, Expected: %d, Received: %d", operand, year, result.Year())
	}
	if result.Day() != day {
		t.Errorf("ParseTime day incorrect.  Operand: %s, Expected: %d, Received: %d", operand, day, result.Day())
	}
}

func TestParseURL(t *testing.T) {
	operand, scheme, host, path := "https://github.com/bobziuchkovski/haven", "https", "github.com", "/bobziuchkovski/haven"
	result, err := ParseURL(operand)
	if err != nil {
		t.Errorf("ParseURL encountered unexpected error: %s.  Operand: %s", err, operand)
	}
	if result.Scheme != scheme {
		t.Errorf("ParseURL scheme incorrect.  Operand: %s, Expected: %s, Received: %s", operand, scheme, result.Scheme)
	}
	if result.Host != host {
		t.Errorf("ParseURL host incorrect.  Operand: %s, Expected: %s, Received: %s", operand, host, result.Host)
	}
	if result.Path != path {
		t.Errorf("ParseURL path incorrect.  Operand: %s, Expected: %s, Received: %s", operand, path, result.Path)
	}
}

func TestQuote(t *testing.T) {
	operand, expected := "cat", "\"cat\""
	result := Quote(operand)
	if result != expected {
		t.Errorf("Quote result incorrect.  Operand: %s, Expected: %s, Received: %s", operand, expected, result)
	}
}

func TestQuoteRegex(t *testing.T) {
	operand, expected := "[cat].*", "\\[cat\\]\\.\\*"
	result := QuoteRegex(operand)
	if result != expected {
		t.Errorf("QuoteRegex result incorrect.  Operand: %s, Expected: %s, Received: %s", operand, expected, result)
	}
}

func TestRepeat(t *testing.T) {
	operand, count, expected := "duck", 2, "duckduck"
	result := Repeat(count, operand)
	if result != expected {
		t.Errorf("Repeat result incorrect.  Operand: %s, Count: %d, Expected: %s, Received: %s", operand, count, expected, result)
	}
}

func TestReplace(t *testing.T) {
	operand, old, newstr, n, expected := "horse", "rse", "g", 1, "hog"
	result := Replace(old, newstr, n, operand)
	if result != expected {
		t.Errorf("Replace result incorrect.  Operand: %s, Old: %s, New: %s, N: %d, Expected: %s, Received: %s", operand, old, newstr, n, expected, result)
	}
}

func TestSlice(t *testing.T) {
	operand, first, last, expected := []string{"cat", "dog", "mouse"}, 1, 3, []string{"dog", "mouse"}
	result := Slice(first, last, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Slice result incorrect.  Operand: %s, First: %d, Last: %d, Expected: %s, Received: %s", operand, first, last, expected, result)
	}
}

func TestSplit(t *testing.T) {
	operand, sep, expected := "cat,dog,horse", ",", []string{"cat", "dog", "horse"}
	result := Split(sep, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Split result incorrect.  Operand: %#v, Sep: %s, Expected: %s, Received: %s", operand, sep, expected, result)
	}
}

func TestSplitAfter(t *testing.T) {
	operand, sep, expected := "cat,dog,horse", ",", []string{"cat,", "dog,", "horse"}
	result := SplitAfter(sep, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SplitAfter result incorrect.  Operand: %#v, Sep: %s, Expected: %s, Received: %s", operand, sep, expected, result)
	}
}

func TestSplitAfterN(t *testing.T) {
	operand, sep, n, expected := "cat,dog,horse", ",", 2, []string{"cat,", "dog,horse"}
	result := SplitAfterN(sep, n, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SplitAfterN result incorrect.  Operand: %#v, Sep: %s, N: %d, Expected: %s, Received: %s", operand, sep, n, expected, result)
	}
}

func TestSplitN(t *testing.T) {
	operand, sep, n, expected := "cat,dog,horse", ",", 2, []string{"cat", "dog,horse"}
	result := SplitN(sep, n, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SplitN result incorrect.  Operand: %#v, Sep: %s, N: %d, Expected: %s, Received: %s", operand, sep, n, expected, result)
	}
}

func TestSubtract(t *testing.T) {
	operand, a, expected := 42, 10, 32
	result := Subtract(a, operand)
	if result != expected {
		t.Errorf("Subtract result incorrect.  Operand: %d, A: %d, Expected: %d, Received: %d", operand, a, expected, result)
	}
}

func TestTitle(t *testing.T) {
	operand, expected := "cat", "Cat"
	result := Title(operand)
	if result != expected {
		t.Errorf("Title result incorrect.  Operand: %s, Expected: %s, Received: %s", operand, expected, result)
	}
}

func TestToLower(t *testing.T) {
	operand, expected := "CAT", "cat"
	result := ToLower(operand)
	if result != expected {
		t.Errorf("ToLower result incorrect.  Operand: %s, Expected: %s, Received: %s", operand, expected, result)
	}
}

func TestToUpper(t *testing.T) {
	operand, expected := "cat", "CAT"
	result := ToUpper(operand)
	if result != expected {
		t.Errorf("ToUpper result incorrect.  Operand: %s, Expected: %s, Received: %s", operand, expected, result)
	}
}

func TestTrim(t *testing.T) {
	operand, chars, expected := "cat,dog,horse", ",ctae", "dog,hors"
	result := Trim(chars, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Trim result incorrect.  Operand: %s, Chars: %s, Expected: %s, Received: %s", operand, chars, expected, result)
	}
}

func TestTrimLeft(t *testing.T) {
	operand, chars, expected := "cat,dog,horse", ",ctae", "dog,horse"
	result := TrimLeft(chars, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TrimLeft result incorrect.  Operand: %s, Chars: %s, Expected: %s, Received: %s", operand, chars, expected, result)
	}
}

func TestTrimPrefix(t *testing.T) {
	operand, prefix, expected := "cat,dog,horse", "cat", ",dog,horse"
	result := TrimPrefix(prefix, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TrimPrefix result incorrect.  Operand: %s, Prefix: %s, Expected: %s, Received: %s", operand, prefix, expected, result)
	}
}

func TestTrimRight(t *testing.T) {
	operand, chars, expected := "cat,dog,horse", ",ctae", "cat,dog,hors"
	result := TrimRight(chars, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TrimRight result incorrect.  Operand: %s, Chars: %s, Expected: %s, Received: %s", operand, chars, expected, result)
	}
}

func TestTrimSpace(t *testing.T) {
	operand, expected := "   cat,dog,horse   ", "cat,dog,horse"
	result := TrimSpace(operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TrimSpace result incorrect.  Operand: %s, Expected: %s, Received: %s", operand, expected, result)
	}
}

func TestTrimSuffix(t *testing.T) {
	operand, suffix, expected := "cat,dog,horse", "horse", "cat,dog,"
	result := TrimSuffix(suffix, operand)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TrimSuffix result incorrect.  Operand: %s, Suffix: %s, Expected: %s, Received: %s", operand, suffix, expected, result)
	}
}

func TestUnquote(t *testing.T) {
	operand, expected := "\"cat\"", "cat"
	result, err := Unquote(operand)
	if err != nil {
		t.Errorf("Unquote encountered unexpected error: %s.  Operand: %s", err, operand)
	}
	if result != expected {
		t.Errorf("Unquote result incorrect.  Operand: %s, Expected: %s, Received: %s", operand, expected, result)
	}
}
