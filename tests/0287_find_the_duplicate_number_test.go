package tests

import (
    "testing"
)

/**
 * [287] Find the Duplicate Number
 *
 * Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,3,4,2,2]
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,1,3,4,2]
 * Output: 3
 * 
 * Note:
 * 
 * 
 * 	You must not modify the array (assume the array is read only).
 * 	You must use only constant, O(1) extra space.
 * 	Your runtime complexity should be less than O(n^2).
 * 	There is only one duplicate number in the array, but it could be repeated more than once.
 * 
 * 
 */

func TestFindtheDuplicateNumber(t *testing.T) {
    var cases = []struct {
        input  []int
        output int
    }{
        {
            input:  []int{1,3,4,2,2},
            output: 2,
        },
        {
            input:  []int{3,1,3,4,2},
            output: 3,
        },
    }
    for _, c := range cases {
        x := findDuplicate(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func findDuplicate(nums []int) int {
    var fast int
    slow := nums[0]
    fast = nums[nums[0]]

    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }

    fast = 0
    for slow != fast {
        fast = nums[fast]
        slow = nums[slow]
    }
    return slow
}

// submission codes end
