package main

func bracketOrder(s string) bool {
    count := 0
    for _, ch := range s {
        if ch == '(' {
            count++
        } else if ch == ')' {
            count--
        }
        if count < 0 {
            return false
        }
    }
    return count == 0
}