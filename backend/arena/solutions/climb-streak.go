package main

func climbStreak(heights []int) int {
    maxStreak := 0
    current := 0
    for _, h := range heights {
        if h > 0 {
            current++
            if current > maxStreak {
                maxStreak = current
            }
        } else {
            current = 0
        }
    }
    return maxStreak
}