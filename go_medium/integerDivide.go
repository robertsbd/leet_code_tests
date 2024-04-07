// This is a brute force implementation, and I could implement a long division algorithm, but I can't be bothered
// Integer divide function implemented with out using *, /, % or any Math functions that may make use of these

func checkIfLarge(n int) int {

    if n > 2147483647 {return 2147483647}
    if n < -2147483648 {return -2147483648}
    return n
}

func getSign(x, y int) int {

    if (x > 0 && y < 0) || (x < 0 && y > 0) {
        return - 1
    }
    return 1
}

func setSign(n, sign int) int {
    if sign == - 1 {
        return n - n - n
    }
    return n
}

func divide(dividend int, divisor int) int {
    
    var n int = 0 
    var sign int = 1
    var rem int = dividend
    var divi int = divisor

    if rem < 0 {
        rem = rem - rem - rem
        sign = -1
    }

    if divi < 0 {
        divi = divi - divi - divi
        sign = getSign(-1, sign)
    }

    if divi == 1 {n = rem
    } else if divi > rem {n = 0
    } else if divi == rem {n = 1
    } else {
        for rem >= divi {
            rem -= divi
            n++
        }
    }

    return checkIfLarge(setSign(n, sign))
}
