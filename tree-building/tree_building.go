package tree

import (
    "fmt"
    "sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

var continuous bool

func Build(records []Record) (*Node, error) {
	records = sortRecords(records)
	if isDiscontinuous(records) {
		return nil, fmt.Errorf("Not continuous")
	}
	result, error := BuildTree(records, true)
	return result, error
}

func BuildTree(records []Record, ordered bool) (*Node, error) {
	//Wrote this as a separate function to avoid isDiscontinuous check for subtrees
	if !ordered {
		records = sortRecords(records)
	}
	if len(records) == 0 {
		return nil, nil
	}
	var parentPresent bool //To check if multiple parents are present
	var result Node
	tree := subTree(records) // To get subtrees of the element
	for _, v := range records {
		err := validate(records, v, parentPresent)
		if err != nil {
			return nil, err
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
	sortChildren(result.Children)
	return &result, nil
}

func subTree(records []Record) map[int][]Record {
	//Function to get all children for a node
	tree := make(map[int][]Record)
	for _, v := range records {
		if v.ID != v.Parent {
		  tree[v.Parent] = append(tree[v.Parent], v)
		}
	}
	return tree

}

func containsDuplicate(rs []Record, rec Record) bool {
	// Function to detect duplicates
	var counter int
	for _, r := range rs {
		if r == rec {
			counter ++
			if counter == 2 {
    			return true
			}
		}
	}
	return false
}

func isDiscontinuous(orderedRecords []Record) (result bool) {
	// Function to check if the Node.ID are continuous
	var prev int
	for _, a := range orderedRecords {
		if a.ID != prev && a.ID != prev + 1 {
			return true
		}
		prev = a.ID
	}
	return
}

func validate(records []Record, v Record, parentPresent bool) error {
	if containsDuplicate(records, v) {
		return fmt.Errorf("Duplicate")
	}
	if v.ID < v.Parent {
		return fmt.Errorf("Greater than Parent")
	}
	if (v.ID == v.Parent) && parentPresent {
		return fmt.Errorf("Cyclic")
	}
	return nil
}

// Is it possible to make use of interface or something so that I could have one function instead of `sortRecords` and `sortChildren`
func sortRecords(records []Record) []Record {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	return records

}

func sortChildren(children []*Node) []*Node {
	sort.Slice(children, func(i, j int) bool {
		return children[i].ID < children[j].ID
	})
	return children
}
