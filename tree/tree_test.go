package tree

import (
	"testing"
	"fmt"
)

// go test -v .\tree -test.run TestPreorderTraversal
func TestPreorderTraversal(t *testing.T){

	gNode := TreeNode{
		Val: "g",
		Left:	nil,
		Right: nil,
	}

	fNode := TreeNode{
		Val: "f",
		Left:	&gNode,
		Right: nil,
	}

	eNode := TreeNode{
		Val: "e",
		Left:	nil,
		Right: nil,
	}



	bNode := TreeNode{
		Val: "b",
		Left:	nil,
		Right: nil,
	}

	dNode := TreeNode{
		Val: "d",
		Left:	nil,
		Right: &eNode,
	}

	cNode := TreeNode{
		Val: "c",
		Left:	&dNode,
		Right: &fNode,
	}



	aNode := TreeNode{
		Val: "a",
		Left:	&bNode,
		Right: &cNode,
	}



	p := travPre(&aNode)

	fmt.Println(p)
}
