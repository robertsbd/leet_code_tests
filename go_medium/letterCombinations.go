func letterCombinations(digits string) []string {

    var out_string[]string

    my_map := map[byte][]byte{
        '2': {'a', 'b', 'c'},
        '3': {'d', 'e', 'f'},
        '4': {'g', 'h', 'i'},   
        '5': {'j', 'k', 'l'},
        '6': {'m', 'n', 'o'},
        '7': {'p', 'q', 'r', 's'},    
        '8': {'t', 'u', 'v'},
        '9': {'w', 'x', 'y', 'z'},}

    if len(digits) == 1 {
        for _, first_num := range my_map[digits[0]] {
            out_string = append(out_string, string(first_num))
        }
    }

    if len(digits) == 2 {
        for _, first_num := range my_map[digits[0]] {
            for _, second_num := range my_map[digits[1]] {
                out_string = append(out_string, string(first_num) + string(second_num))
            }
        }
    }

    if len(digits) == 3 {
        for _, first_num := range my_map[digits[0]] {
            for _, second_num := range my_map[digits[1]] {
                for _, third_num := range my_map[digits[2]] {
                    out_string = append(out_string, string(first_num) + string(second_num) + string(third_num))
                }
            }
        }
    }

    if len(digits) == 4 {
        for _, first_num := range my_map[digits[0]] {
            for _, second_num := range my_map[digits[1]] {
                for _, third_num := range my_map[digits[2]] {
                    for _, fourth_num := range my_map[digits[3]] {
                        out_string = append(out_string, string(first_num) + string(second_num) + string(third_num) + string(fourth_num))
                    }
                }
            }
        }
    }

    return out_string

}
