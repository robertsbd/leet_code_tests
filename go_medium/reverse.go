func reverse(x int) int {

    t := strconv.Itoa(x)

    var out string = ""
    var sign int = 1

    if x < 0 {
        t = t[1:]
        sign = -1
    } 
    
    for _,c:= range t{
        out = string(c) + out
    }

    out_int, _ := strconv.Atoi(out) 

    if out_int > 2147483647 {return 0}

    return out_int * sign
}
