package main

import (
	"fmt"
)

type Case struct {
  Order int
  Name  string
}

var cases = []Case{
    {Order: 1, Name: "one"},
    {Order: 2, Name: "two"},
    {Order: 3, Name: "tri"}, 
    {Order: 4, Name: "fou"}, 
    {Order: 5, Name: "fiv"},
  }

func reorder(c *[]Case, len int){
  for i := 0; i < len; i++{
    (*c)[i].Order = i+1;
  }
}

func rearrange(c []Case, a int, b int) []Case {
  a--;
  b--;
  //nothing to rearrange
  if a == b {
    return c;
  }
  if a < b {
    moved := c[a]; 
    for i := a; i <= b; i++ {
      c[i] = c[i+1];
    }
    c[b] = moved;
    reorder(&c, len(c));
  }else if a > b {
    moved := c[a];
    for i := a; i > b; i--{
      c[i] = c[i-1];
    }
    c[b] = moved;
    reorder(&c, len(c));
  }
  return c
}

func main() {
  cases = rearrange(cases, 4, 2);
  cases = rearrange(cases, 1, 2);

  fmt.Println(cases)
}