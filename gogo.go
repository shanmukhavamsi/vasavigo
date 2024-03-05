package main

import "fmt"

func main() {
	n:=6
	if n%2 == 0 || n == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	if num := 9;num > 0{
		fmt.Println(num," is negative")
	}
	i:=1
	switch i {
	case 1:
		fmt.Println("iam")
		default :
		fmt.Println("you")
	}
	/*whatami(1.0)
	whatami := func( i any){
		switch t := i.(type){
		case bool:
			fmt.Println("boolean value")
		}
	
	}*/
var arr[5]int
	fmt.Println(len(arr))
	brr := [5]int{1,2,3,4}
	fmt.Println(brr)
    var crr[3][2]int
	fmt.Println(crr)
	drr := [3][2]int{{1,2},{3,4},{5,6}}
	fmt.Println(drr)
	var b [3][2]int
	var x int =0
    for i:=0;i<3;i++{
		for j:=0;j<2;j++{
			b[i][j]=x
				x++
			}
		}
		fmt.Println(b)
	/*var arrt []int
	arrt =append(arrt,1)
	fmt.Println(arrt)
	var a = make([]int,3)
	fmt.Println(cap(a),len(a))
	arrt = append(arrt,2)
	fmt.Println(arrt)
	a[3]=10
	fmt.Println(cap(a),len(a))*/
	var str1 ="abcd"
	/*var str2 ="efgh"*/
	fmt.Println(str1[1:3])
	m :=make(map[string]int)
    m["hello"]=1
	m["world"]=2
	m["vamsi"]=3
	fmt.Println(m["vamsi"])
	delete(m,"world")

	fmt.Println(m)
   for i,0=range(array)
}

	
