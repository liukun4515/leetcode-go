package tests

import (
    "reflect"
    "testing"
)

/**
 * [389] Find the Difference
 *
 * 
 * Given two strings s and t which consist of only lowercase letters.
 * 
 * String t is generated by random shuffling string s and then add one more letter at a random position.
 * 
 * Find the letter that was added in t.
 * 
 * Example:
 * 
 * Input:
 * s = "abcd"
 * t = "abcde"
 * 
 * Output:
 * e
 * 
 * Explanation:
 * 'e' is the letter that was added.
 * 
 */

func TestFindtheDifference(t *testing.T) {
    var cases = []struct {
        f      string
        s      string
        output byte
    }{
        {
            f:      "abcd",
            s:      "abcde",
            output: 'e',
        },
        {
            f:      "zx",
            s:      "zxa",
            output: 'a',
        },
    }
    for _, c := range cases {
        x := findTheDifference(c.f, c.s)
        if !reflect.DeepEqual(x, c.output) {
            t.Fail()
        }
    }
}

// submission codes start here

func findTheDifference(s string, t string) byte {
   l := 0
   for _, i := range s {
       l = l ^ int(i)
   }
   for _, i := range t {
       l = l ^ int(i)
   }
   return byte(l)
}

// submission codes end
