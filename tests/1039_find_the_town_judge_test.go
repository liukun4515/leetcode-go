package tests

import (
	"testing"
)

/**
 * [1039] Find the Town Judge
 *
 * In a town, there are N people labelled from 1 to N.  There is a rumor that one of these people is secretly the town judge.
 *
 * If the town judge exists, then:
 *
 *
 * 	The town judge trusts nobody.
 * 	Everybody (except for the town judge) trusts the town judge.
 * 	There is exactly one person that satisfies properties 1 and 2.
 *
 *
 * You are given trust, an array of pairs trust[i] = [a, b] representing that the person labelled a trusts the person labelled b.
 *
 * If the town judge exists and can be identified, return the label of the town judge.  Otherwise, return -1.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: N = 2, trust = [[1,2]]
 * Output: 2
 *
 *
 *
 * Example 2:
 *
 *
 * Input: N = 3, trust = [[1,3],[2,3]]
 * Output: 3
 *
 *
 *
 * Example 3:
 *
 *
 * Input: N = 3, trust = [[1,3],[2,3],[3,1]]
 * Output: -1
 *
 *
 *
 * Example 4:
 *
 *
 * Input: N = 3, trust = [[1,2],[2,3]]
 * Output: -1
 *
 *
 *
 * Example 5:
 *
 *
 * Input: N = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
 * Output: 3
 *
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= N <= 1000
 * 	trust.length <= 10000
 * 	trust[i] are all different
 * 	trust[i][0] != trust[i][1]
 * 	1 <= trust[i][0], trust[i][1] <= N
 *
 *
 */

func TestFindtheTownJudge(t *testing.T) {
	var cases = []struct {
		input  int
		trust  [][]int
		output int
	}{
		{
			input:  2,
			trust:  [][]int{{1, 2}},
			output: 2,
		},
		{
			input:  3,
			trust:  [][]int{{1, 3}, {2, 3}},
			output: 3,
		},
		{
			input:  3,
			trust:  [][]int{{1, 3}, {2, 3}, {3, 1}},
			output: -1,
		},
		{
			input:  3,
			trust:  [][]int{{1, 2}, {2, 3}},
			output: -1,
		},
	}
	for _, c := range cases {
		x := findJudge(c.input, c.trust)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func findJudge(N int, trust [][]int) int {
	l := len(trust)
	if l == 0 {
		return 1
	}
	notJudge := make([]bool, N+1)
	trusts := make([]int, N+1)
	for i := 0; i < l; i++ {
		notJudge[trust[i][0]] = true
		trusts[trust[i][1]] += 1
	}
	judge := 0
	for i := 1; i <= N; i++ {
		if !notJudge[i] {
			if judge != 0 {
				// there cannot be more than one judge
				return -1
			}
			judge = i
		}
	}
	if judge != 0 {
		if trusts[judge] == N-1 {
			return judge
		}
	}
	return -1
}

// submission codes end
