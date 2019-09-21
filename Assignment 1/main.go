package main
import "fmt"

func sums(arr []int)(int){

  // Base Condition, 
	if len(arr) == 0 {
		return 0;

  //Iteration
	} else{
		var res = arr[0];
		res += sums(arr[1:]);
		return res;
	}
}


func main(){
	// your code goes here
	arr := []int{10, 20, 30, 40, 50};
	var result = sums(arr);
	fmt.Println(result);
	
}