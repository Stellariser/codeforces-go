// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	exampleIns := [][]string{{`["/a","/a/b","/c/d","/c/d/e","/c/f"]`}, {`["/a","/a/b/c","/a/b/d"]`}, {`["/a/b/c","/a/b/d","/a/b/ca"]`}}
	exampleOuts := [][]string{{`["/a","/c/d","/c/f"]`}, {`["/a"]`}, {`["/a/b/c","/a/b/ca","/a/b/d"]`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	const targetCaseNum = 0
	if err := testutil.RunLeetCodeFuncWithCase(t, removeSubfolders, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}