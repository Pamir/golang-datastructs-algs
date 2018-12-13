///https://flaviocopes.com/golang-data-structure-linked-list/

package main
import(
	"fmt"
	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Node struct {
	content Item
	next *Node
}

type LinkedList struct{
	head *Node
	size int

}
func(ll *LinkedList) Append(t Item){

	node := Node{t,nil}
	if ll.head == nil {
		ll.head = &node
	}else{
		last := ll.head
		for{
			if last.next == nil{
				break
			}
			last = last.next
		}
		last.next = &node
	}
	ll.size++
}

func Dump(ll *LinkedList){
	if ll == nil{
		return
	}
	node := ll.head
	for{
		if node == nil{
			break
		}
		fmt.Println(node.content)
		node = node.next
	}
}

func main()  {
	ll := LinkedList{nil,0}
	ll.Append(10)
	ll.Append(12)
	ll.Append(15)
	Dump(&ll)
}


