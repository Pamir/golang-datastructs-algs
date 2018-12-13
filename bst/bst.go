package main

import (
	"errors"
	"fmt"
	"github.com/cheekybits/genny/generic"
	"math/rand"
	"time"
)

type Item generic.Type

type Node struct{
	LeftNode *Node
	RightNode *Node
	Data int
}

type BST struct{
	RootNode *Node
}

func(bst *BST) Add(item int) error{
	if bst == nil{
		return errors.New("BST can not be null")
	}
	node := Node{nil,nil,item}
	if bst.RootNode == nil {
		bst.RootNode = &node
		return nil
	}
	temp := bst.RootNode
	for {
		if temp.Data > item{
			if temp.RightNode == nil{
				temp.RightNode = &node
				return nil
			}else{
				temp = temp.RightNode
			}
		}else{
			if temp.LeftNode == nil{
				temp.LeftNode = &node
				return nil
			}else{
				temp = temp.LeftNode
			}
		}

	}
	return nil
}


func main(){
	bst := BST{}
	rand.Seed(time.Now().Unix())
	for i :=0; i < 10; i++{
		data :=  rand.Int() % 100

		fmt.Printf("Inserting %d \n",data)
		bst.Add(data)
	}
	fmt.Println("Finished")

}
