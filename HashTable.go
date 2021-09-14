package HashTable

import (
	"bytes"
	"fmt"
)

const MAP_SIZE = 1000

type Node struct {
	key   int
	value int
	next  *Node
}

type HashMap struct {
	Data []*Node
}

/**
	Create a new HashMap structure
**/
func NewHashMap() *HashMap {
	return &HashMap{Data: make([]*Node, MAP_SIZE)}
}

/**
	Insert a new element in the HashMap
	@param key
	@param value
	NOTE: collision are handled with linked list
**/
func (h *HashMap) Insert(key int, value int) {

	index := getIndex(key)
	if h.Data[index] == nil {
		h.Data[index] = &Node{key: key, value: value} // new element
	} else {
		// Index collision, add element to linked-list
		starting_node := h.Data[index]
		curr_node := h.Data[index]
		for ; curr_node != nil; curr_node = curr_node.next {
			if curr_node.key == key {
				// the key exists, modify value
				curr_node.value = value
				return
			}
		}
		// new element
		curr_node = &Node{key: key, value: value}
		starting_node.next = curr_node
	}
}

/**
	Get an element from the HashMap
	@param key
	@param value
	@return (value, true) if element exists
			(-1, false) otherwise
**/
func (h *HashMap) Get(key int) (int, bool) {
	index := getIndex(key)
	if h.Data[index] != nil {
		starting_node := h.Data[index]
		for ; starting_node != nil; starting_node = starting_node.next {
			if starting_node.key == key {
				// key matched
				return starting_node.value, true
			}
		}
	}

	// key does not exists
	return -1, false
}

/**
Remove an element from the HashTable
@param key
**/
func (h *HashMap) Remove(key int) {

	index := getIndex(key)
	if h.Data[index] != nil {
		head := h.Data[index]
		if h.Data[index].key == key && h.Data[index].next == nil {
			h.Data[index] = nil
		} else {
			to_del := head
			if to_del.key == key {
				h.Data[index] = to_del.next
				to_del = nil
			} else {
				curr := to_del.next
				for ; curr != nil; curr = curr.next {
					if curr.key == key {
						to_del.next = curr.next
						curr = nil
						break
					} else {
						to_del = curr
					}
				}
			}
		}
	}
}

/* Hash function (see https://stackoverflow.com/a/12996028) */
func hash(x uint32) (hash uint32) {
	x = ((x >> 16) ^ x) * 0x45d9f3b
	x = ((x >> 16) ^ x) * 0x45d9f3b
	x = (x >> 16) ^ x
	return x
}

func getIndex(key int) (index int) {
	return int(hash(uint32(key))) % MAP_SIZE
}

func (n *Node) PrintNode() string {
	return fmt.Sprintf("<Key: %d, Value: %d>\n", n.key, n.value)
}

func (h *HashMap) PrintHMbuff() {
	var output bytes.Buffer
	fmt.Fprint(&output, "{ ")
	for _, n := range h.Data {
		if n != nil {
			fmt.Fprintf(&output, "[%d: %d] ", n.key, n.value)
			for node := n.next; node != nil; node = node.next {
				fmt.Fprintf(&output, "[%d: %d]", node.key, node.value)
			}
		}
	}
	fmt.Fprint(&output, "}")
	fmt.Println(output.String())
}

func (h *HashMap) PrintHM() {
	fmt.Println("HT { ")
	for _, n := range h.Data {
		if n != nil {
			node := n
			for ; node != nil; node = node.next {
				fmt.Printf("%v", node.PrintNode())
			}
		} //else {
			//fmt.Println("index empty =>", i)
		//}
	}
	fmt.Println("}")
}
