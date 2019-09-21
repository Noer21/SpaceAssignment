package main

import (
	"fmt"
  "strconv"
)

type person struct {
	name   string
	gender string
	age    int
}

func (p *person)Name(name string) *person{
  p.name = name;
  return p
}

func (p *person)Gender(gender string) *person{
  p.gender = gender;
  return p;
}

func (p *person)Age(age int) string{
  p.age = age;
  var res string;
  res = join(res, p.name, ", ");
  res = join(res, p.gender, ", ");
  res = join(res, strconv.Itoa(p.age), ".\n");
  return res;
}

func join(s1 string, s2 string, s3 string) string {
  s1 += s2;
  s1 += s3;
  return s1;
}

func NewPerson () *person{
  p := person{}
  return &p
}

func main() {
	jon := NewPerson().Name("Jon Snow").Gender("Male").Age(24);
	fmt.Println(jon);
}