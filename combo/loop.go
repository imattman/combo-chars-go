package combo

import "sort"

// FromRunes generates combinations of letters represented as strings.
// The empty string is omitted.
func FromRunes(rs []rune) []string {
    // process runes in sorted order
    sort.Sort(runeSlice(rs))

    combos := []string{""}
    for _, c := range rs {
        cAdded := []string{}
        for _, combo := range combos {
            cAdded = append(cAdded, combo + string(c))
        }
        combos = append(combos, cAdded...)
    }
    
    // we don't optimize the loop to remote dupes
    // handle that here
    unique := map[string]int{}
    
    for _, cmb := range combos {
        if len(cmb) == 0 {
            continue // omit empty string
        }
        
        unique[cmb] = 1
    }
    
    // rebuild with unique vals
    combos = []string{}
    for key := range unique {
        combos = append(combos, key)
    }
    
    sort.Strings(combos)
    
    return combos
}


