package main

import()

func main() {}

type NestedInteger struct {}
func (this NestedInteger) IsInteger() bool {return false}
func (this NestedInteger) GetInteger() int {return 0}
func (n *NestedInteger) SetInteger(value int) {} 
func (this *NestedInteger) Add(elem NestedInteger) {}

func (this NestedInteger) GetList() []*NestedInteger {return nil}

type NestedIterator struct {
    flattened []int
    index     int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{
        flattened: flatten(nestedList),
        index:     0,
    }
}

func flatten(nested []*NestedInteger) []int {
    var result []int
    for _, ni := range nested {
        if ni.IsInteger() {
            result = append(result, ni.GetInteger())
        } else {
            result = append(result, flatten(ni.GetList())...)
        }
    }
    return result
}

func (this *NestedIterator) Next() int {
    val := this.flattened[this.index]
    this.index++
    return val
}

func (this *NestedIterator) HasNext() bool {
    return this.index < len(this.flattened)
}


