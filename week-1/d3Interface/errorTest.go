package main

import (
	"errors"
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return a, errors.New("b 不能为 0")
	}
	return b / a, nil
}

func TestDivide(a, b float64) {
	a, err := Divide(a, b)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("0 !=", a)
}

var users = map[int]string{
	1: "hu",
	2: "li",
}

var ErrUserNotFound = errors.New("user not found")

func Find(id int) (string, error) {
	v, ok := users[id]
	if !ok {
		return "", ErrUserNotFound
	}
	return "hello " + v, nil
}

func Greet(id int) (string, error) {
	s, err := Find(id)
	if err != nil {
		return "", fmt.Errorf("err %w", err)
	}
	return s, nil
}

func TestGreet(id int) {
	s, err := Greet(id)
	if err != nil {
		fmt.Println("is not found: ", errors.Is(err, ErrUserNotFound), err)
		return
	}
	fmt.Println(s)
}

type ValidationError struct {
	Field string
	Msg   string
}

func (v *ValidationError) Error() string {
	return v.Field + ": " + v.Msg
}

func Register(name string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Msg: "不能为空"}
	}
	if age <= 0 || age > 150 {
		return &ValidationError{Field: "age", Msg: "age 范围错误"}
	}
	return nil
}

func TestValudation(name string, age int) {
	err := Register(name, age)
	if err != nil {
		fmt.Println("错误信息", err)
		if ve, ok := errors.AsType[*ValidationError](err); ok {
			fmt.Println("请检查字段", ve.Field, ve.Msg)
		}
	}
	fmt.Println("success")
}

func errorTest() {
	TestDivide(10, 0)
	TestDivide(12, 1)

	TestGreet(1)
	TestGreet(3)

	TestValudation("", 10)
	TestValudation("Tom", 0)
	TestValudation("Tom", 12)
}
