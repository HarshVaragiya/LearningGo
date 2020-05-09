package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/tour/tree"
	"sort"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	log.Debug("Current Node value : ", t.Value)
	if t.Left != nil {
		log.Debug("Left is not nil, from ", t.Value, " node -> walking left")
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		log.Debug("Right is not nil, from ", t.Value, " node -> walking right")
		go Walk(t.Right, ch)
	}
}

func analyseTree(tree *tree.Tree, size int) (values []int) { // Calls walk, sorts data in the channel
	treeChannel := make(chan int, size)
	Walk(tree, treeChannel)
	for i := 0; i < 10; i++ {
		x := <-treeChannel
		log.Info("Index ", i, " = ", x)
		values = append(values, x)
	}
	sort.Ints(values)
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree, size int) bool {
	tree1Values := analyseTree(t1, size)
	tree2Values := analyseTree(t2, size)
	validCount := 0
	for index := 0; index < len(tree1Values); index++ {
		if tree1Values[index] != tree2Values[index] {
			return false
		} else {
			validCount += 1
		}
	}
	if validCount == len(tree1Values) {
		log.Info("All ", validCount, " Values match between the two trees!")
		return true
	}
	return false
}

func main() {
	log.SetLevel(log.WarnLevel)
	areEqual := Same(tree.New(1), tree.New(1), 10)
	fmt.Println("Equivalence : ", areEqual)
}
