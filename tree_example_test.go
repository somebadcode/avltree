package avltree_test

import (
	"fmt"

	"github.com/somebadcode/avltree"
)

func Example() {
	tree := avltree.New[int, string](1)

	tree.Insert(45, "rabbit")
	tree.Insert(10, "sparrow")
	tree.Insert(87, "capybara")
	tree.Insert(90, "pig")
	tree.Insert(89, "dog")
	tree.Insert(33, "cat")
	tree.Insert(29, "tiger")
	tree.Insert(55, "lion")

	tree.InorderTraversal(func(k int, s string) bool {
		fmt.Println(k, s)

		return true
	})

	fmt.Println("Tree size:", tree.Size())

	// Output:
	// 10 sparrow
	// 29 tiger
	// 33 cat
	// 45 rabbit
	// 55 lion
	// 87 capybara
	// 89 dog
	// 90 pig
	// Tree size: 8
}
