package main

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"unicode/utf8"
)

// 错误
var ErrTodoNotFound = errors.New("todo not found")

// 类型
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (t Todo) String() string {
	mark := " "
	if t.Done {
		mark = "x"
	}
	return fmt.Sprintf("[%s] %d %s", mark, t.ID, t.Title)
}

// 接口
type Store interface {
	Add(title string) (Todo, error)
	List() ([]Todo, error)
	Done(id int) (Todo, error)
}

// 内存版，为什么要用 map，而不是直接 List
// 因为查找方便，增删不要考虑顺序
type MemoryStore struct {
	todos  map[int]Todo
	nextId int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{todos: map[int]Todo{}, nextId: 1}
}

func (s *MemoryStore) Add(t string) (Todo, error) {
	if utf8.RuneCountInString(t) < 3 {
		return Todo{}, errors.New("title is too short")
	}
	newTodo := Todo{ID: s.nextId, Title: t, Done: false}
	s.todos[newTodo.ID] = newTodo
	s.nextId++
	return newTodo, nil
}

func (s *MemoryStore) List() ([]Todo, error) {
	list := make([]Todo, 0, len(s.todos))
	for _, v := range s.todos {
		list = append(list, v)
	}
	slices.SortFunc(list, func(a, b Todo) int {
		return cmp.Compare(a.ID, b.ID)
	})
	return list, nil
}

func (s *MemoryStore) Done(id int) (Todo, error) {
	t, ok := s.todos[id] // 拷贝一份，而不是直接指向
	if !ok {
		return Todo{}, ErrTodoNotFound
	}
	t.Done = true
	s.todos[id] = t // 为什么要写回
	return t, nil
}

func main() {
	var s1 Store = NewMemoryStore()
	t, err := s1.Add("hello")
	if err != nil {
		fmt.Println("err not found", errors.Is(err, ErrTodoNotFound))
	}
	fmt.Println(t)
	t2, e2 := s1.Done(1)
	if e2 != nil {
		fmt.Println("e2", e2, errors.Is(e2, ErrTodoNotFound))
	}
	fmt.Println(t2)
	fmt.Println(s1.List())
}
