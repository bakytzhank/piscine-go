package piscine

type NodeL struct {
	Data interface{}                  
	Next *NodeL                                                            
}                                                                       

type List struct {                                              
	Head *NodeL                                                      
	Tail *NodeL                                                           
}                                                              

func ListRemoveIf(l *List, data_ref interface{}) {
	current := l.Head
	var prev *NodeL
	for current != nil {
		if current.Data == data_ref {
			if prev == nil {
				l.Head = current.Next
			} else {
				prev.Next = current.Next
			}
		} else {
			prev = current
		}
		current = current.Next
	}
}
