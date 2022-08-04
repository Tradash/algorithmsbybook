package types

import "sync"

type NodeQueue struct {
	Items []Vertex
	Lock  sync.RWMutex
}

type PriorityQueue []*Vertex

type Node struct {
	Value string
}

type Vertex struct {
	Node     *Node
	Distance int
}

type Edge struct {
	Node   *Node
	Weight int
}

func (s *NodeQueue) NewQ() *NodeQueue {
	s.Lock.Lock()
	s.Items = []Vertex{}
	s.Lock.Unlock()
	return s
}

func (s *NodeQueue) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items) == 0
}

func (s *NodeQueue) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items)
}

// Удаление веришны с начала очереди

func (s *NodeQueue) Dequeue() *Vertex {
	s.Lock.Lock()
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	s.Lock.Unlock()
	return &item
}

// Добавление узла в конец очереди
func (s *NodeQueue) Enqueue(t Vertex) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		return
	}
	insertFlag := false

	for k, v := range s.Items {
		if t.Distance < v.Distance {
			// Если не первый элемент
			if k > 0 {
				s.Items = append(s.Items[:k+1], s.Items[k:]...)
				s.Items[k] = t
				insertFlag = true
			} else {
				s.Items = append([]Vertex{t}, s.Items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
		if !insertFlag {
			s.Items = append(s.Items, t)
		}
	}

}
