package main

func breachRun(target int, nums []int) int {
    for i, n := range nums {
        if n >= target {
            return i
        }
    }
    return -1
}