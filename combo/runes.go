package combo

// runeSlice provides sort.Interface methods
type runeSlice []rune

// Len is the number of elements in the collection.
func (rs runeSlice) Len() int {
    return len(rs)
}
 
// Less reports whether the element with
// index i should sort before the element with index j.
func (rs runeSlice) Less(i, j int) bool {
    return rs[i] < rs[j]
}

// Swap swaps the elements with indexes i and j.
func (rs runeSlice) Swap(i, j int) {
    rs[i], rs[j] = rs[j], rs[i]
}
