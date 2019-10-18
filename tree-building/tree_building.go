package tree

import "fmt"
import "sort"

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

var continuos bool

func Build(records []Record) (*Node, error) {
	records = orderRecords(records)
	if isDisContinuous(records) {
		return nil, fmt.Errorf("Not continuous")
	}
	result, error := BuildTree(records, true)
	return result, error
}

func BuildTree(records []Record, ordered bool) (*Node, error) {
	if !ordered {
		records = orderRecords(records)
	}
	if len(records) == 0 {
		return nil, nil
	}
	var parentPresent bool
	var result Node
	tree := subTree(records)
	for _, v := range records {
		err := handleErrors(records, v, parentPresent)
		if err != "" {
			return nil, fmt.Errorf(err)
		}
		if v.ID == v.Parent {
			result.ID = v.ID
			parentPresent = true
		} else {
			if result.ID == v.Parent {
				sub_child, _ := BuildTree(append(tree[v.ID], Record{v.ID, v.ID}), false)
				result.Children = append(result.Children, sub_child)
			}
		}
	}
	if !parentPresent {
		return nil, fmt.Errorf("No Parent")
	}
	orderChildren(result.Children)
	return &result, nil
}

func subTree(records []Record) map[int][]Record {
	tree := make(map[int][]Record)
	for _, v := range records {
		if v.ID != v.Parent {
			if len(tree[v.Parent]) == 0 {
				tree[v.Parent] = []Record{v}
			} else {
				tree[v.Parent] = append(tree[v.Parent], v)
			}

		}
	}
	return tree

}

func contains(s []Record, e Record) bool {
	var counter int
	for _, a := range s {
		if a == e {
			counter = counter + 1
		}
	}
	return counter == 2
}

func isDisContinuous(orderedRecords []Record) (result bool) {
	var prev int
	for _, a := range orderedRecords {
		result = !((prev == a.ID) || ((prev + 1) == a.ID))
		if result {
			return true
		}
		prev = a.ID
	}
	return
}

func handleErrors(records []Record, v Record, parentPresent bool) string {
	if contains(records, v) {
		return "No Parent"
	}
	if v.ID < v.Parent {
		return "Greater than Parent"
	}
	if (v.ID == v.Parent) && parentPresent {
		return "Cyclic"
	}
	return ""
}

func orderRecords(records []Record) []Record {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	return records

}

func orderChildren(children []*Node) []*Node {
	sort.Slice(children, func(i, j int) bool {
		return children[i].ID < children[j].ID
	})
	return children
}
