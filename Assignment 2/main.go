package main
import "fmt"

func contain(arr []string, word string)(string){
  var res = "Not Contain";
  if len(arr) == 0 {
    res = "Not Contain";
  } else if arr[0] == word {
    res = "Contain";
  }else {
    res = contain(arr[1:], word);
  }
  return res;
}

func main(){
	// your code goes here
	arr := []string{"Banna", "Apple", "Orange", "Durian", "Pear"};
	var word = "Apple"
  
  var res = contain(arr, word);
  fmt.Printf("%s", res);
}