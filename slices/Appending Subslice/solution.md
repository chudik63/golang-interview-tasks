# Solution
1) Initially the slice has length of 3 and capacity of 4 => there are 3 zeros inside: `[0 0 0]`
2) We take `slice[:1]` that has length 1 and still has capacity of 4 => `[0]`.
3) We append 1 to the end of the copy of the subslice => `[0 1]`. Considering the fact that the initial array stored by pointer inside all slices, we are replaced second element in it with 1.
4) When we print our first slice we get `[0 1 0]`