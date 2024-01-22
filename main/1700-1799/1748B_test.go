// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1748/B
// https://codeforces.com/problemset/status/1748/problem/B
func Test_cf1748B(t *testing.T) {
	testCases := [][2]string{
		{
			`7
1
7
2
77
4
1010
5
01100
6
399996
5
23456
18
789987887987998798`,
			`1
2
10
12
10
15
106`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1748B)
}