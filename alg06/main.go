package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// Эффект Матфея , стр 64, Самоорганизующийся список

type ListHeader struct {
	start *List
	last  *List
	lock  sync.RWMutex
}

type List struct {
	id   int
	next *List
}

func newList() *ListHeader {
	list := new(ListHeader)
	return list
}

func (l *ListHeader) add(x int) {
	l.lock.Lock()
	n := &List{id: x, next: nil}
	if l.start == nil {
		l.start = n
		l.last = n
	} else {
		l.last.next = n
		l.last = n
	}
	l.lock.Unlock()
}

func (l *ListHeader) next(n *List) *List {
	return n.next
}

func (l *ListHeader) findBefore(n *List) (*List, *List) {

	if l.start == n {
		return nil, nil
	}
	var before00, before01 *List
	before00 = nil
	before01 = l.start
	for x := l.start.next; x != nil; x = x.next {
		if x == n {
			return before00, before01
		}
		before00 = before01
		before01 = x

	}
	return nil, nil
}

func (l *ListHeader) moveUp(n *List) {
	//l.print()

	before00, before01 := l.findBefore(n)

	if before00 == nil && before01 == nil {
		return
	}

	if before00 == nil {
		nNext := n.next
		sNext := l.start
		l.start = n
		l.start.next = sNext
		sNext.next = nNext
	} else {
		nTmp := n.next
		n.next = before01
		before01.next = nTmp
		before00.next = n
	}
	//l.print()
}

func (l *ListHeader) find(x int) *List {
	if l.start == nil {
		return nil
	}
	for i := l.start; i.next != nil; i = i.next {
		if i.id == x {
			return i
		}
	}
	return nil
}

func (l *ListHeader) print() {
	for i := l.start; i.next != nil; i = i.next {
		fmt.Printf("%d, ", i.id)
	}

}

func (l *ListHeader) findSwap(x int) *List {
	if l.start == nil {
		return nil
	}

	for i := l.start; i.next != nil; i = i.next {
		if i.id == x {
			l.lock.Lock()
			l.moveUp(i)
			l.lock.Unlock()
			return i
		}

	}

	return nil
}

var searchData []int

func createList(m int) *ListHeader {

	list := newList()
	for i := 0; i < m; i++ {
		x := rand.Intn(10000000)
		//fmt.Println(x)
		if list.find(x) == nil {
			list.add(x)
			if i%10 == 0 {
				searchData = append(searchData, x)
			}
		}

	}
	return list
}

type megaSearch []search

type search struct {
	id    int
	value int
}

func (s megaSearch) Len() int {
	return len(s)
}

func (s megaSearch) Less(i, j int) bool {
	return s[i].id < s[j].id
}

func (s megaSearch) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	fmt.Println("Подготовка данных")
	searchData = make([]int, 0, 0)
	maxElement := 5000
	l := createList(maxElement)
	sData := make(megaSearch, 0, 0)

	for i := 0; i < len(searchData); i++ {
		for j := 0; j < i; j++ {
			sData = append(sData, search{
				id:    rand.Intn(maxElement * 10),
				value: searchData[i],
			})
		}

	}

	sort.Sort(sData)

	//fmt.Println(searchData, l)

	tries := 20
	fmt.Println("Поиск без оптимизации")
	for i := 0; i < tries; i++ {
		startTime := time.Now()
		for j := 0; j < len(sData); j++ {
			d := l.find(sData[j].value)
			if d == nil {
				fmt.Println("Не найдено", sData[j].value)
			}
		}
		timeDuration := time.Now().Sub(startTime)
		fmt.Printf("Тест №: %d, Размер данных: %d, Поисковых запросов %d, Продолжительность %s\n", i+1, maxElement, len(sData), timeDuration)
	}
	//l.print()
	fmt.Println("Поиск c оптимизацей")
	for i := 0; i < tries; i++ {
		//l.print()
		startTime := time.Now()
		for j := 0; j < len(sData); j++ {
			d := l.findSwap(sData[j].value)
			if d == nil {
				fmt.Println("Не найдено", sData[j].value)
			}
		}
		timeDuration := time.Now().Sub(startTime)
		fmt.Printf("Тест №: %d, Размер данных: %d, Поисковых запросов %d, Продолжительность %s\n", i+tries+1, maxElement, len(sData), timeDuration)
	}
}
