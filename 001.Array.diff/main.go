package main

// https://www.codewars.com/kata/523f5d21c841566fde000009/train/go
// func ArrayDiff(a, b []int) (r []int) {
// 	for _, aval := range a {
// 		isExists := false
// 		for _, bval := range b {
// 			if aval == bval {
// 				isExists = true
// 				break
// 			}
// 		}
// 		if !isExists {
// 			r = append(r, aval)
// 		}
// 	}
// 	return r
// }

// Clever solution
func ArrayDiff(a, b []int) (r []int) {
	m := map[int]bool{}
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if !m[v] {
			r = append(r, v)
		}
	}
	return
}

func main() {
	res := ArrayDiff([]int{1, 2}, []int{1})
	println("Result:")
	for _, aval := range res {
		print("[", aval, "]\n")
	}
}
