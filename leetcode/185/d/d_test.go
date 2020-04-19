// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	exampleIns := [][]string{{`2`, `3`, `1`}, {`5`, `2`, `3`}, {`9`, `1`, `1`}, {`50`, `100`, `25`}, {`37`, `17`, `7`}}
	exampleOuts := [][]string{{`6`}, {`0`}, {`1`}, {`34549172`}, {`418930126`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithCase(t, numOfArrays, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-185/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/
