// This is a brute force implementation, and I could implement a long division algorithm, but I can't be bothered

func checkIfLarge(n int) int {

    if n > 2147483647 {return 2147483647}
    if n < -2147483648 {return -2147483648}
    return n
}

func divide(dividend int, divisor int) int {
    
    var n int = 0 
    var sign int = 1
    var rem int = dividend
    var divi int = divisor

    if rem < 0 {
        rem = rem * -1
        sign = sign * -1
    }

    if divi < 0 {
        divi = divi * -1
        sign = sign * -1
    }

    if divi > rem {return 0}

    if divi == rem {return 1 * sign}

    if divi == 1 {return checkIfLarge(rem * sign)}

    for rem >= divi {
        rem = rem - divi
        n++
    }

    return checkIfLarge(n * sign)
}
