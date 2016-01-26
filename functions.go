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
	"bufio"
	"encoding/base64"
	"math/rand"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var pseudo = rand.New(rand.NewSource(time.Now().UnixNano()))

// FuncMap is a map of all functions exported by haven.  It is meant for use with
// ext/template.Template.Funcs()
var FuncMap = map[string]interface{}{
	"Abs":          Abs,
	"Add":          Add,
	"Base64Decode": Base64Decode,
	"Base64Encode": Base64Encode,
	"CompileERE":   CompileERE,
	"CompileRegex": CompileRegex,
	"Contains":     Contains,
	"ContainsAny":  ContainsAny,
	"Count":        Count,
	"Divide":       Divide,
	"Fields":       Fields,
	"Grep":         Grep,
	"HasPrefix":    HasPrefix,
	"HasSuffix":    HasSuffix,
	"Head":         Head,
	"Index":        Index,
	"IndexAny":     IndexAny,
	"Intersect":    Intersect,
	"Join":         Join,
	"LastIndex":    LastIndex,
	"LastIndexAny": LastIndexAny,
	"Lines":        Lines,
	"Matches":      Matches,
	"Max":          Max,
	"Min":          Min,
	"Modulo":       Modulo,
	"Multiply":     Multiply,
	"Now":          Now,
	"ParseBool":    ParseBool,
	"ParseFloat":   ParseFloat,
	"ParseInt":     ParseInt,
	"ParseTime":    ParseTime,
	"ParseURL":     ParseURL,
	"Quote":        Quote,
	"QuoteRegex":   QuoteRegex,
	"Repeat":       Repeat,
	"Replace":      Replace,
	"Reverse":      Reverse,
	"Seq":          Seq,
	"Shuffle":      Shuffle,
	"Slice":        Slice,
	"Sort":         Sort,
	"Split":        Split,
	"SplitAfter":   SplitAfter,
	"SplitAfterN":  SplitAfterN,
	"SplitN":       SplitN,
	"Subtract":     Subtract,
	"Tail":         Tail,
	"Title":        Title,
	"ToLower":      ToLower,
	"ToUpper":      ToUpper,
	"Trim":         Trim,
	"TrimLeft":     TrimLeft,
	"TrimPrefix":   TrimPrefix,
	"TrimRight":    TrimRight,
	"TrimSpace":    TrimSpace,
	"TrimSuffix":   TrimSuffix,
	"Union":        Union,
	"Unquote":      Unquote,
}

/*
 * String manipulation
 */

// Contains uses strings.Contains to check if substr is part of operand.
func Contains(substr, operand string) bool { return strings.Contains(operand, substr) }

// ContainsAny uses strings.Contains to check if any of chars are part of operand.
func ContainsAny(chars, operand string) bool { return strings.ContainsAny(operand, chars) }

// Count uses strings.Count to return the number of occurrences of str in operand.
func Count(str, operand string) int { return strings.Count(operand, str) }

// Fields uses strings.Fields to split operand on whitespace.
func Fields(operand string) []string { return strings.Fields(operand) }

// HasPrefix uses strings.HasPrefix to check if operand begins with prefix.
func HasPrefix(prefix, operand string) bool { return strings.HasPrefix(operand, prefix) }

// HasSuffix uses strings.HasSuffix to check if operand ends with suffix.
func HasSuffix(suffix, operand string) bool { return strings.HasSuffix(operand, suffix) }

// Index uses strings.Index to return the first index of substr in operand, or -1 if missing.
func Index(substr, operand string) int { return strings.Index(operand, substr) }

// IndexAny uses strings.IndexAny to return the first index of any of chars in operand, or -1 if missing.
func IndexAny(chars, operand string) int { return strings.IndexAny(operand, chars) }

// Join uses strings.Join to return the strings of operand joined by sep.
func Join(sep string, operand []string) string { return strings.Join(operand, sep) }

// LastIndex uses strings.LastIndex to return the last index of substr in operand, or -1 if missing.
func LastIndex(substr, operand string) int { return strings.LastIndex(operand, substr) }

// LastIndexAny uses strings.LastIndexAny to return the last index of any of chars in operand, or -1 if missing.
func LastIndexAny(chars, operand string) int { return strings.LastIndexAny(operand, chars) }

// Lines splits operand into a slice of lines with the end-of-line terminators removed.
func Lines(operand string) (lines []string, err error) {
	// TODO: Give thought to eliminating error return
	reader := strings.NewReader(operand)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

// Quote uses strconv.Quote to return operand in quoted form.
func Quote(operand string) string { return strconv.Quote(operand) }

// Repeat uses strings.Repeat to return operand repeated count times.
func Repeat(count int, operand string) string { return strings.Repeat(operand, count) }

// Replace uses strings.Replace to replace the first n occurrences of old with new.
func Replace(old, new string, n int, operand string) string {
	return strings.Replace(operand, old, new, n)
}

// Split uses strings.Split to split operand on sep, dropping sep from the resulting slice.
func Split(sep, operand string) []string { return strings.Split(operand, sep) }

// SplitAfter uses strings.SplitAfter to split operand after sep, leaving sep in the resulting slice.
func SplitAfter(sep, operand string) []string { return strings.SplitAfter(operand, sep) }

// SplitAfterN uses strings.SplitAfterN to split operand after sep, leaving sep in the resulting
// slice.  Splits at most n times.
func SplitAfterN(sep string, n int, operand string) []string {
	return strings.SplitAfterN(operand, sep, n)
}

// SplitN uses strings.SplitN to split operand on sep, dropping sep from the resulting slice.
// Splits at most n times.
func SplitN(sep string, n int, operand string) []string { return strings.SplitN(operand, sep, n) }

// Title uses strings.Title to return operand with the first unicode codepoint of each word
// converted to uppercase.
func Title(operand string) string { return strings.Title(operand) }

// ToLower uses strings.ToLower to return operand with all unicode codepoints converted to lowercase.
func ToLower(operand string) string { return strings.ToLower(operand) }

// ToUpper uses strings.ToUpper to return operand with all unicode codepoints converted to uppercase.
func ToUpper(operand string) string { return strings.ToUpper(operand) }

// Trim uses strings.Trim to remove any occurrences of chars from the beginning and end of operand.
func Trim(chars string, operand string) string { return strings.Trim(operand, chars) }

// TrimLeft uses strings.TrimLeft to remove any occurrences of chars from the beginning of operand.
func TrimLeft(chars string, operand string) string { return strings.TrimLeft(operand, chars) }

// TrimPrefix uses strings.TrimPrefix to remove prefix from the beginning of operand.
func TrimPrefix(prefix, operand string) string { return strings.TrimPrefix(operand, prefix) }

// TrimRight uses strings.TrimRight to remove any occurrences of chars from the end of operand.
func TrimRight(cutset string, operand string) string { return strings.TrimRight(operand, cutset) }

// TrimSpace uses strings.TrimSpace to remove all unicode whitespace codepoints from the beginning
// and end of operand.
func TrimSpace(operand string) string { return strings.TrimSpace(operand) }

// TrimSuffix uses strings.TrimSuffix to remove suffix from the end of operand.
func TrimSuffix(suffix, operand string) string { return strings.TrimSuffix(operand, suffix) }

// Unquote uses strconv.Unquote to the underlying, unquoted string value of operand.
func Unquote(operand string) (unquoted string, err error) { return strconv.Unquote(operand) }

/*
 * Slice Manipulation
 */

// Grep filters operand according to pattern, returning a slice of matching elements.
// Pattern is treated as a regexp.
func Grep(pattern string, operand []string) ([]string, error) {
	rex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	var matching []string
	for _, elem := range operand {
		if rex.MatchString(elem) {
			matching = append(matching, elem)
		}
	}
	return matching, nil
}

// Head returns the first n elements of operand.  If less than n elements are in operand,
// it returns all of operand.
func Head(n int, operand []string) []string {
	if len(operand) < n {
		return operand
	}
	return operand[:n]
}

// Intersect returns the intersection of a and operand.  Duplicate elements are removed.
// Ordering is not preserved.
func Intersect(a, operand []string) []string {
	intersection := make([]string, Max(len(a), len(operand)))
	uniqA := make(map[string]bool, len(a))
	for _, elem := range a {
		uniqA[elem] = true
	}
	i := 0
	uniqOp := make(map[string]bool, len(operand))
	for _, elem := range operand {
		if uniqA[elem] && !uniqOp[elem] {
			intersection[i] = elem
			i++
		}
		uniqOp[elem] = true
	}
	return intersection[:i]
}

// Reverse returns a copy of operand with the elements in reverse order.
func Reverse(operand []string) []string {
	reversed := make([]string, len(operand))
	for i := range operand {
		reversed[len(operand)-i-1] = operand[i]
	}
	return reversed
}

// Seq generates a sequence of ints from first to last.  If incr is specified
// (an optional third argument), then the sequence will increment by incr.
// Otherwise, incr defaults to 1.  Incr may be negative to generate a sequence
// of descending ints.
func Seq(first, last int, incr ...int) []int {
	j := 1
	if len(incr) > 1 {
		panic("Seq incr cannot be more than one value in length")
	}
	if len(incr) == 1 {
		if (incr[0]) == 0 {
			panic("Seq incr value cannot be zero")
		}
		j = incr[0]
	}

	var values []int
	current := first
	if j > 0 {
		for current <= last {
			values = append(values, current)
			current += j
		}
	} else {
		for current >= last {
			values = append(values, current)
			current += j
		}
	}
	return values
}

// Shuffle returns a copy of operand with the elements shuffled pseudo-randomly.
func Shuffle(operand []string) []string {
	shuffled := make([]string, len(operand))
	for i, p := range pseudo.Perm(len(operand)) {
		shuffled[i] = operand[p]
	}
	return shuffled
}

// Slice returns operand[first:last].
func Slice(first, last int, operand []string) []string { return operand[first:last] }

// Sort returns a copy of operand sorted by sort.Strings.
func Sort(operand []string) []string {
	sorted := make([]string, len(operand))
	for i := range operand {
		sorted[i] = operand[i]
	}
	sort.Strings(sorted)
	return sorted
}

// Tail returns the last n elements of operand.  If less than n elements are in operand,
// it returns all of operand.
func Tail(n int, operand []string) []string {
	if len(operand) < n {
		return operand
	}
	return operand[len(operand)-n:]
}

// Union returns the union of a and operand.  Duplicate elements are removed.
// Ordering is not preserved.
func Union(a, operand []string) []string {
	uniq := make(map[string]bool, len(a)+len(operand))
	for _, elem := range a {
		uniq[elem] = true
	}
	for _, elem := range operand {
		uniq[elem] = true
	}
	union := make([]string, len(uniq))
	i := 0
	for k := range uniq {
		union[i] = k
		i++
	}
	return union[:i]
}

/*
 * Time
 */

// Now uses time.Now to return the current time as a time.Time instance.
func Now() time.Time { return time.Now() }

// ParseTime uses time.Parse to return a time.Time instance of operand parsed according to format.
func ParseTime(format, operand string) (time.Time, error) { return time.Parse(format, operand) }

/*
 * Regular Expressions
 */

// Matches use regexp.MatchString to check if operand matches pattern.
func Matches(pattern string, operand string) (bool, error) {
	return regexp.MatchString(pattern, operand)
}

// CompileRegex uses regexp.Compile to compile a new *regexp.Regexp according to pattern.
func CompileRegex(pattern string) (*regexp.Regexp, error) { return regexp.Compile(pattern) }

// CompileERE uses regexp.CompilePOSIX to compile a new *regexp.Regexp according to pattern.
// Uses POSIX extended regexp behavior.
func CompileERE(pattern string) (*regexp.Regexp, error) { return regexp.CompilePOSIX(pattern) }

// QuoteRegex uses regexp.QuoteMeta to returnsoperand with all regex metachars escaped.
func QuoteRegex(operand string) string { return regexp.QuoteMeta(operand) }

/*
 * Encoding
 */

// Base64Encode uses base64.StdEncoding to encode operand.
func Base64Encode(operand string) string { return base64.StdEncoding.EncodeToString([]byte(operand)) }

// Base64Decode uses base64.StdEncoding to decode operand.
func Base64Decode(operand string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(operand)
	return string(bytes), err
}

// ParseBool uses strconv.ParseBool to parse operand as a bool.
func ParseBool(operand string) (value bool, err error) { return strconv.ParseBool(operand) }

// ParseInt uses strconv.ParseInt with base == 0 (auto) and bitsize = 32 to parse operand as an int.
func ParseInt(operand string) (value int, err error) {
	i64, err := strconv.ParseInt(operand, 0, 32)
	return int(i64), err
}

// ParseFloat uses strconv.ParseFloat with bitsize == 64 to parse operand as a float.
func ParseFloat(operand string) (f float64, err error) { return strconv.ParseFloat(operand, 64) }

// ParseURL uses url.Parse to parse operand as a url.
func ParseURL(operand string) (url *url.URL, err error) { return url.Parse(operand) }

/*
 * Math
 */

// Abs returns the absolute value of operand.
func Abs(operand int) int {
	if operand < 0 {
		return operand * -1
	}
	return operand
}

// Add a to operand.
func Add(a, operand int) int { return operand + a }

// Subtract a from operand.
func Subtract(a, operand int) int { return operand - a }

// Divide operand by a.
func Divide(a, operand int) int { return operand / a }

// Modulo returns operand modulo a.
func Modulo(a, operand int) int { return operand % a }

// Multiply operand and a.
func Multiply(a, operand int) int { return operand * a }

// Min returns the minimum of a and operand.
func Min(a, operand int) int {
	if a < operand {
		return a
	}
	return operand
}

// Max returns the maximum of a and operand.
func Max(a, operand int) int {
	if a > operand {
		return a
	}
	return operand
}
