package main
/*
import "fmt"


//Arrays Questions
func main(){


	//create a array length of 5 and print all 5 elements 
    
	arr := [5]int{1,2,3,4,5}

	for i := 0; i < len(arr); i++ {
        fmt.Print( arr[i]," ")
    }
    fmt.Println()
	//maximum number in an array of integers.

	max := maxElement(arr[:])
    fmt.Println(max)

	reverseArray(arr[:])

	for i := 0; i < len(arr); i++ {
        fmt.Print( arr[i]," ")
    }
    fmt.Println()
	

}

func maxElement(nums[] int)int{
	if (len(nums) == 0){
		return 0
	}
	maxEle := nums[0]
	for index := range nums{
		if nums[index] > maxEle {
			maxEle = nums[index]
		}

	}
	return maxEle

}

func reverseArray(Arr[]int){
   i, j := 0, len(Arr)-1
	for i < j {
		swap(&Arr[i], &Arr[j])
		i++
		j--
	}
}


func swap(a, b *int) {
	*a, *b = *b, *a
}
	*/

/*
//Slice Questions	
func main(){
	//create a slice and return a new slice with only odd numbers
	slice := []int{1,2,3,4,5,6,7,8,9}
    
	//appends a value to a slice and prints the new length and capacity of the slice.
    fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 10,11)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
    
    newSlice := oddNumbers(slice)
	for index := range newSlice{
		fmt.Print(newSlice[index]," ")
	}
	fmt.Println()

	afterRemovingEle := removeEle(slice)
	for index := range afterRemovingEle{
		fmt.Print(afterRemovingEle[index]," ")
	}
	fmt.Println()

}

func oddNumbers(nums[] int)[] int{
	newSlice := []int{}
	for index := range nums{
		if(nums[index] % 2 != 0){
			newSlice = append(newSlice, nums[index])
		}
	}
	return newSlice
}

func removeEle(nums[] int)[] int{
	elementToRemove := 2;
	nums = append(nums[:elementToRemove], nums[elementToRemove+1:]...)

	return nums
}

*/

/*// Function to check if a key exists in the map
func ifExist(people map[string]int, key string) {
	if age, ok := people[key]; ok {
		fmt.Println( key, age)
	} else {
		fmt.Printf("%s does not exist in the map\n", key)
	}
}

func main() {
	// Create a map with names as keys and ages as values
	people := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 22,
	}

	// Print all names and ages
	for name, age := range people {
		fmt.Println(name, age)
	}

	// Check if a specific key exists
	ifExist(people, "Bob")
	ifExist(people, "David")
}*/

/*
//Create a program to count how many times each word appears in a slice of strings using a map.
func main(){
	words := []string{"apple","banana","apple","orange","watermelon","watermelon","apple"}

	wordCount := make(map[string]int)

	for index := range words{
		wordCount[words[index]]++
	}

	for word,count := range wordCount {
		fmt.Println(word," ",count)
	}
}
*/
/*
//Functions Questions
func function(a,b int)(int,int,int){
	add := a+b
	diff := a-b
	multiply := a*b
	return add,diff,multiply
}*/

/*
//Checking the prime number or not
func PrimtNUmber(a int) string{
   count := 0
   for i := 2; i < a; i++{
	if(a%i == 0){
       count++
	}
   }
    if (count >= 1){
		return "not prime"
	}else{
		return "prime"
	}
}

func main(){
  add, diff, multiply := function(1,2)
  fmt.Println(add)
  fmt.Println(diff)
  fmt.Println(multiply)

  primtNumber := PrimtNUmber(13)
  fmt.Println(primtNumber)

}
*/

/*
//function that returns the largest number from a slice of integers.
func Largest(Nums []int) int{
	maxVal := Nums[0]
	for index := range Nums{
		if Nums[index] > maxVal{
			maxVal = Nums[index]
		}
	}
	return maxVal
}

func main(){
	slice := []int{1,2,3,12,4,5,9}

	largestNumber := Largest(slice)

	fmt.Println(largestNumber)
}
*/

/*
//Method Questions
// a struct called Rectangle with width and height. Write a method Area() to calculate its area.
type Rectangle struct{
	width float64
	height float64
}

func (r Rectangle) Area() float64{
	 return r.width * r.height 
}

func main(){

   rect := Rectangle{width: 10, height: 5}

	fmt.Println(rect.Area())
}

*/

/*
//a struct Person with a name. Write a method SayHello() that prints “Hi, my name is [name]”
type Person struct{
	Name string
}

func (p Person) sayhello(){
	 fmt.Println("hi hello",p.Name)
}

func main(){

   per := Person{Name : "raju"}

   per.sayhello()
}
*/

/*
//Add a method ChangeName(newName string) to Person that updates the name field
type Person struct {
	Name string
}

func (p Person) SayHello() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

func main() {
	person := Person{Name: "Alice"}

	person.SayHello() 

	person.ChangeName("Bob")

	person.SayHello() 
}
	*/