package main

const (
	Max = 100001
)

var (
	tree [Max << 2]int
	arr  [Max]int
)

func build(l, r, i int) {
	if l == r {
		tree[i] = arr[l]
		return
	}

}

func main() {
}
