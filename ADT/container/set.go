package container

import(
	"sync"
)


//利用字典key 不能重复原理
type Set struct{
	m  map[int]struct{}
	len int
	sync.RWMutex

}

//初始化
func NewSet(c int64) *Set{
	temp := make(map[int]struct{}, c)
	return &Set{
		m : temp,
	}

}
//add
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()

	s.m[item] = struct{}{}
	s.len = len(s.m)

}

func (s *Set) Remove(item int){
	s.Lock()
	defer s.Unlock()

	if s.len ==0{
		return
	}
	delete(s.m, item)

	s.len = len(s.m)



}

