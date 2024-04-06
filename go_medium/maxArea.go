func maxArea(height []int) int {
    
    var length = len(height) - 1
    var max int = 0
    var a int
   
    // set an initial value to max
    if height[0] > height[length] {
        max = height[length] * length
    } else {
        max = height[0] * length
    }
 
    for i, h := range height {
        for j := length; j > i; j-- {
            if max > (j-i) * h {break}
            if h > height[j] {
                a = height[j] * (j-i)
            } else {
                a = h * (j-i)
            }
            if a > max {max = a}
        } 
    } 

    return max
}

