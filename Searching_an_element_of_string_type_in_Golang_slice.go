
package main 
  
import ( 
    "fmt"
    "sort"
) 
  
func main() { 
   
    slice_1 := []string{"C", "Go", "Java", "C#", "Ruby"} 
    slice_2 := []string{"GEEKs", "123geeks", "gfg", "GeeksforGeeks"} 
  
    var f1, f2, f3 string 
    f1 = "GEEKs"
    f2 = "C"
    f3 = "gfg"
  
    sort.Strings(slice_1) 
    sort.Strings(slice_2) 
  
    fmt.Println("Slice 1: ", slice_1) 
    fmt.Println("Slice 2: ", slice_2) 
 
 
    res1 := sort.SearchStrings(slice_1, f1) 
    res2 := sort.SearchStrings(slice_2, f2) 
    res3 := sort.SearchStrings(slice_2, f3) 
  
   
    fmt.Println("Result 1: ", res1) 
    fmt.Println("Result 2: ", res2) 
    fmt.Println("Result 3: ", res3) 
  
} 
