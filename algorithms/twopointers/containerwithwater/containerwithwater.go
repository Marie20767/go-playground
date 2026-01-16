package containerwithwater

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Examples
// ________________________________
// Example 1:
// Input: height = [1,5,8,2,5,4,8,3,7]
// Output: 49
// Explanation: See image. The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
// Height of the smallest line at the boundary of the cointainer is max water so in this example from 8 (index 1) -> 6 add 7 all th way until 3 -> 7 so 7 * 7 = 49

// Example 2:
// Input: height = [1,1]
// Output: 1

// given that l, r are indexes of heights:
// get smallest boundary * distance(l, r)
// e.g. [1,5,8,2] l, r := 0, 3 then you'd have 1 * 1 * 1 -> 1 * 3 (distance(r, l) = 3)

func MaxArea(heights []int) int {
    l, r := 0, len(heights) - 1
    res := 0

    for l < r {
        area := min(heights[l], heights[r]) * (r - l)
        if area > res {
            res = area
        }
        if heights[l] <= heights[r] {
            l++
        } else {
            r--
        }
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}