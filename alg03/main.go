package main

import (
	"fmt"
)

// Пути и ДНК, стр 44

type Routes struct {
	isUsed *bool
	vertex string
}

type Branch struct {
	name   string
	routes []Routes
}

type Tree []Branch

func findRoot(t Tree, root string) (int, bool) {

	for i, v := range t {
		if v.name == root {
			return i, true
		}
	}

	return 0, false
}

func createTree(s []string) Tree {
	tree := make([]Branch, 0, 0)

	for i, v := range s {
		f := v[:2]
		s := v[1:]
		fmt.Println(i, v, f, s)
		ind, ok := findRoot(tree, f)
		if !ok {
			tree = append(tree, Branch{
				name:   f,
				routes: nil,
			})
			ind = len(tree) - 1
		}
		flag := false
		tree[ind].routes = append(tree[ind].routes, Routes{
			isUsed: &flag,
			vertex: s,
		})
		ind, ok = findRoot(tree, s)
		if !ok {
			tree = append(tree, Branch{
				name:   s,
				routes: nil,
			})
			ind = len(tree) - 1
		}
		tree[ind].routes = append(tree[ind].routes, Routes{
			isUsed: &flag,
			vertex: f,
		})

	}
	return tree
}

func findEuler(t *Tree, mainStart, start string, r string) string {

	for _, v := range *t {
		if v.name == start {
			for _, w := range v.routes {
				if !*w.isUsed {
					arrow := ","
					if r == "" {
						arrow = ""
					}
					r = r + arrow + w.vertex
					*w.isUsed = true
					return findEuler(t, mainStart, w.vertex, r)
				}
			}
		}
	}
	return r

}

func findFreeRoute(t *Tree) int {
	for i, v := range *t {
		for _, w := range v.routes {
			if !*w.isUsed {
				return i
			}
		}
	}
	return -1
}

func main() {
	dnk := []string{"GTG", "TGG", "ATG", "GGC", "GCG", "CGT", "GCA", "TGC", "CAA", "AAT", "ATC", "TCA"}
	tree := createTree(dnk)
	r := make([][]string, 0)
	for i := findFreeRoute(&tree); i != -1; i = findFreeRoute(&tree) {
		r = append(r, []string{findEuler(&tree, tree[i].name, tree[i].name, tree[i].name)})
	}
	fmt.Println(r)

}
