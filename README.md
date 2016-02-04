[![Build Status](https://travis-ci.org/ziuchkovski/haven.svg?branch=master)](https://travis-ci.org/ziuchkovski/haven)
[![Coverage](http://gocover.io/_badge/github.com/ziuchkovski/haven?2)](http://gocover.io/github.com/ziuchkovski/haven)
[![Report Card](http://goreportcard.com/badge/ziuchkovski/haven)](http://goreportcard.com/report/ziuchkovski/haven)
[![GoDoc](https://godoc.org/github.com/ziuchkovski/haven?status.svg)](https://godoc.org/github.com/ziuchkovski/haven)

# Haven

## Overview

Haven is a collection of safe template functions for use with Go's text/template package. "Safe" refers to safe execution
on the host system.  These functions provide no access to the host filesystem, resources, environment variables, etc.

Many of the functions are thin wrappers around existing standard library functions.  However, the primary operand is
always the last parameter, which enables chained pipelining.  Example: haven.Split and haven.Join reverse the parameters
of strings.Split and strings.Join, making the following possible:
{{"dogs,cats,horses" | Split "," | Join "\n" }}.

Haven is a work-in-progress.  Contributions and suggestions are welcome!

## API Promise

Minor breaking changes may occur prior to the 1.0 release.  After the 1.0 release, the API is guaranteed to remain backwards compatible.

## Function By Category

Please see the [godocs](https://godoc.org/github.com/ziuchkovski/haven) for usage details.

### String and String Slice Manipulation

Contains, ContainsAny, Count, Fields, HasPrefix, HasSuffix, Index, IndexAny, Join, LastIndex, LastIndexAny, Lines, Quote, Repeat, Replace, Split, SplitAfter, SplitAfterN, SplitN, Title, ToLower, ToUpper, Trim, TrimLeft, TrimPrefix, TrimRight, TrimSpace, TrimSuffix, Unquote, Grep, Head, Intersect, Reverse, Seq, Shuffle, Slice, Sort, Tail, Union

### Map Manipulation

_Absent.  I would love help in compiling a simple, flexible set of map functions that don't require too much magic.  The main issue
comes down to type handling.  Should functions operate only on map[string]string, or should map[string]interface{} be supported for
nested maps and/or non-string types.  What's the right balance of simplicity and reflection?_

### Date and Time

Now, ParseTime

### Regular Expressions

Matches, CompileRegex, CompileERE, QuoteRegex

### Encoding and Parsing

Base64Encode, Base64Decode, ParseBool, ParseInt, ParseFloat, ParseURL

### Math

Abs, Add, Subtract, Divide, Modulo, Multiply, Min, Max

## Authors

Bob Ziuchkovski (@ziuchkovski)

## License (MIT)

Copyright (c) 2016 Bob Ziuchkovski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

