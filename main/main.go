package main 


/*import (

	"fmt"
)
   */


//func main () {



   /*
   // this is how we can take input from the user in Go
   // fmt.Scanln() is used to take input from the user
   var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

   */

	/*
	 // these are the types of declarations in Go
	var i int = 10
	j := 20 // short declaration operator
	        // we cannot assign a value to a variable with the short declaration operator if the variable is already declared
			// we cannot use the short declaration operator outside of a function
	var k int //after declaration, we can assign value later
	k = 30 //we can assign a value to a variable after declaration
	//   constant declaration
	const pi float64 = 3.14 // constant declaration
							// we cannot change the value of a constant
							// we can use the short declaration operator to declare a constant
	*/
	

	/*
	// these are the types of print statements in Go
	i := "Hello"
	j := "world"
	fmt.Print(i)
	fmt.Print(j ,"\n")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("%%",i)
	*/

   /* 
   // these are the genaral formatting verbs in Go which we can use for all data types
   var i float32 = 3.33
   fmt.Printf("%v\n", i)// %v prints the value of the variable
   fmt.Printf("%#v\n", int(i)) // %#v prints the value of the variable in Go syntax
   fmt.Printf("%T\n", i) // %T prints the type of the variable
   fmt.Printf("%v%%\n",i)// %% prints the percent sign

   var s string = "Hello, world!"
   fmt.Printf("%v\n", s) // %v prints the string value
   fmt.Printf("%#v\n", s) // %#v prints the string value in Go syntax
   fmt.Printf("%T\n", s) // %T prints the type of the variable
   */
   
   /*
   // these are the formatting verbs for integers
   var i int = 10
   fmt.Printf("%d\n", i) // %d prints the integer value in decimal
   fmt.Printf("%+d\n", i) // %+d prints the integer value in decimal with a plus sign
   fmt.Printf("%b\n", i) // %b prints the integer value in binary
   fmt.Printf("%o\n", i) // %o prints the integer value in octal
   fmt.Printf("%O\n", i) // %O prints the integer value in octal (uppercase)
   fmt.Printf("%x\n", i) // %x prints the integer value in hexadecimal
   fmt.Printf("%X\n", i) // %X prints the integer value in hexadecimal (uppercase)
   fmt.Printf("%#x\n", i) // %#x prints the integer value in hexadecimal with 0x prefix
   fmt.Printf("%4d\n", i) // %4d prints the integer value in decimal with a width of 4
   fmt.Printf("%04d\n", i) // %04d prints the integer value in decimal with a width of 4 and leading zeros
   fmt.Printf("%+4d\n", i) // %+4d prints the integer value in decimal with a width of 4 and leading zeros with a plus sign
   */
    
  /*  // these are the formatting verbs for strings
   var txt = "Hello"
 
  fmt.Printf("%s\n", txt) // %s prints the string value
  fmt.Printf("%q\n", txt) // %q prints the string value in double quotes
  fmt.Printf("%8s\n", txt) // %8s prints the string value with a width of 8
  fmt.Printf("%-8s\n", txt) // %-8s prints the string value with a width of 8 and left-justified
  fmt.Printf("%x\n", txt) // %x prints the string value in hexadecimal
  fmt.Printf("% x\n", txt) // % x prints the string value in hexadecimal with spaces
  */

  /*
   // these are the formatting verbs for booleans
  var i = true
  var j = false

  fmt.Printf("%t\n", i)// %t prints the boolean value
  fmt.Printf("%v\n", j)// %v prints the boolean value
  */


  /*
   // these are the formatting verbs for floats
  var i = 3.141
  fmt.Printf("%e\n", i)// %e prints the float value in scientific notation
  fmt.Printf("%f\n", i) // %f prints the float value in decimal
  fmt.Printf("%.2f\n", i) // %.2f prints the float value with 2 decimal places
  fmt.Printf("%6.2f\n", i) // %6.2f prints the float value with 2 decimal places and a width of 6
  fmt.Printf("%g\n", i) // %g prints the float value in a compact form
 */
  
  /*
  // intializing arrays in Go
  // arrays are fixed size, so we need to specify the size of the array at the
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}
  arr3 := [...]int{9,10,11,12,13} // the size of the array is determined by the number of elements in the array
  // we can also use the short declaration operator to declare an array

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)
   // we can also declare an array with a specific size and then assign values to it
   var arr4 [5]int
   arr4[0] = 1
   arr4[1] = 2
   arr4[2] = 3
   arr4[3] = 4
   arr4[4] = 5
   fmt.Println(arr4)

  // updating array elements
  prices := [3]int{10,20,30}

  prices[2] = 50
  fmt.Println(prices)

  // array intialization 
  arr12 := [5]int{} //not initialized
  arr23 := [5]int{1,2} //partially initialized
  arr34 := [5]int{1,2,3,4,5} //fully initialized

  fmt.Println(arr12)
  fmt.Println(arr23)
  fmt.Println(arr34)

  // Initialize Only Specific Elements
   arr5 := [5]int{1: 10, 3: 30} // only elements at index 1 and 3 are initialized
   fmt.Println(arr5)

   */


/*
   // slices in Go
   // slices are dynamic arrays, so we don't need to specify the size of the slice
   myslice1 := []int{}
   fmt.Println(len(myslice1))// length of the slice
   fmt.Println(cap(myslice1))// capacity of the slice
   fmt.Println(myslice1)

   myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
   fmt.Println(len(myslice2))
   fmt.Println(cap(myslice2))
   fmt.Println(myslice2)

   // create a slice from an array
   arr1 := [6]int{10, 11, 12, 13, 14, 15}
   myslice := arr1[2:4] // create a slice from an array, this will create a slice with elements at index 2 and 3 of the array

   fmt.Println(myslice)
   fmt.Println(len(myslice))// length of the slice
   fmt.Println(cap(myslice))// capacity of the slice
 
   // creating a slice using the make function
   // the make function is used to create a slice with a specific length and capacity
  myslice12 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice12)
  fmt.Printf("length = %d\n", len(myslice12))
  fmt.Printf("capacity = %d\n", cap(myslice12))

  // with omitted capacity means if not specified, the capacity will be the same as the length
  // so the capacity will be 5 in this case
  myslice22 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice22)
  fmt.Printf("length = %d\n", len(myslice22))
  fmt.Printf("capacity = %d\n", cap(myslice22))
  
  // accessing and slicing and changing elements in a slice is same as in as array

  // we can also append elements to a slice using the append function
  myslice3 := []int{1, 2, 3}
  myslice3 = append(myslice3, 4, 5, 6)
  fmt.Printf("myslice3 = %v\n", myslice3)


  // appending a slice to another slice
   myslice4 := []int{7, 8, 9}  
   myslice3 = append(myslice3, myslice4...) // the ... operator is used to append a slice to another slice
   fmt.Println(myslice3)
   
   // we can also append a slice to another slice using the copy function
   numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  // Original slice
  fmt.Printf("numbers = %v\n", numbers)
  fmt.Printf("length = %d\n", len(numbers))
  fmt.Printf("capacity = %d\n", cap(numbers))

  // Create copy with only needed numbers
  neededNumbers := numbers[:len(numbers)-10]
  numbersCopy := make([]int, len(neededNumbers))
  copy(numbersCopy, neededNumbers)

  fmt.Printf("numbersCopy = %v\n", numbersCopy)
  fmt.Printf("length = %d\n", len(numbersCopy))
  fmt.Printf("capacity = %d\n", cap(numbersCopy))
*/


/*
// Arithmetic Operators in GO
Operator   Name	               Description	                              Example	
+	        Addition	            Adds together two values	               x + y	
-	        Subtraction	         Subtracts one value from another	         x - y	
*	        Multiplication	      Multiplies two values	                  x * y	
/	        Division	            Divides one value by another	            x / y	
%	        Modulus	            Returns the division remainder	         x % y	
++	        Increment	            Increases the value of a variable by 1	   x++	
--	        Decrement	            Decreases the value of a variable by 1	   x--

*/

/*
// Assignment Operators in GO
Operator	      Example	         Same As	
=	            x = 5	             x = 5	
+=	            x += 3	          x = x + 3	
-=	            x -= 3	          x = x - 3	
*=	            x *= 3	          x = x * 3	
/=	            x /= 3	          x = x / 3	
%=	            x %= 3	          x = x % 3	
&=	            x &= 3	          x = x & 3	
|=	            x |= 3	          x = x | 3	
^=	            x ^= 3	          x = x ^ 3	
>>=	         x >>= 3	          x = x >> 3	
<<=	         x <<= 3	          x = x << 3

*/


/*
//logical operators in GO
Operator	    Name	           Description	                                              Example	
&& 	       Logical and	  Returns true if both statements are true	                x < 5 &&  x < 10	
|| 	       Logical or	     Returns true if one of the statements is true	             x < 5 || x < 4	
!	          Logical not	  Reverse the result, returns false if the result is true	 !(x < 5 && x < 10)
*/

/* 
// syntax for conditional statements in Go
if condition {
  // code to be executed if condition is true
}else if anotherCondition {
  // code to be executed if anotherCondition is true
} else {
  // code to be executed if all conditions are false
}

//syntax for nested conditional statements in Go
if condition {
  if anotherCondition {
    // code to be executed if both conditions are true
  }
   else {
    // code to be executed if the first condition is true and the second condition is false
   }
}else {
  // code to be executed if the first condition is false
  }


*/

/*
/// syntax for switch statements in Go
// for single expression switch
switch expression {
case x:
   // code block
case y:
   // code block
case z:
...
default:
   // code block
}

// for multiple expression switch
switch expression {
case x, y, z:
   // code block
default:
   // code block
}
// for no expression switch
switch {
case x:
   // code block
case y:
   // code block
case z:
   // code block
default:
   // code block

// if in case we specify the fallthrough in end of code block in any case if it was false also it will be excuting the code of block
switch {
case x:
   // code block
   fallthrough //it will excute the function if the condition is false also
case y:
   // code block
case z:
   // code block
default:
   // code block


}
*/


/*
// syntax for for loop in Go

for initialization; condition; updation {
   // code to be executed  
   //nested for loop
   for innerInitialization; innerCondition; innerUpdation {
      // code to be executed
      break // is used to end the loop where u keep it 
   }
}
// syntax for for range loop in Go
collection := []int{1, 2, 3, 4, 5} // can be an array, slice, string, map, or channel
// for range loop iterates over the collection and returns the index and value of each element
for index, value := range collection {
   // range keyword is used to iterate over the collection
   // index and value are the variables that hold the index and value of the element in the collection
   // code to be executed
   //index is the index of the element in the collection
   //value is the value of the element in the collection
   // -- is used to ignore the index or value if not needed we can specift what we want to ignore

}

*/

/*
// functions in Go
// syntax to create a basic function in GO
// functions should be declared outside of the main function and called inside the main function
func functionName() {
   //code to be executed

}
// syntax to create a function with parameters in GO
func functionName(parameter1 type, parameter2 type) {
   // code to be executed
   // parameters are the variables that are passed to the function
   // type is the data type of the parameter
   // we can have multiple parameters separated by commas
   // we need to use those parameters in the function and they will send while in the calling function as parameter like this
   // functionName(value1, value2) value2 and value1 aare the parameters and type is known as data type we call value1 and value2 as arguments
}

// syntax to create a function with return value in GO
func functionName(parameter1 type, parameter2 type) returnType {
   // code to be executed
   // returnType is the data type of the return value
   // we need to store the functioncall in a variable and print the value int main function
   // we can have multiple return values separated by commas 
   // we can destructure the return values in the calling function like this
   // value1, value2 := functionName(value1, value2) // value1
   return value1 // return statement is used to return the value from the function
}

// syntax to create a recursive function in GO
func recursiveFunction(parameter1 type) returnType {
   // code to be executed
   // recursive function is a function that calls itself
   // we need to have a base case to stop the recursion
   if parameter1 == 0 {
      return 0 // base case
   }
   return parameter1 + recursiveFunction(parameter1 - 1) // recursive call
}
func main() {
   // calling the function
   functionName() // calling the function without parameters
   functionName(value1, value2) // calling the function with parameters
   value := functionName(value1, value2) // calling the function with return value
   fmt.Println(value) // printing the return value of the function

   // calling the recursive function
   result := recursiveFunction(5) // calling the recursive function with parameter 5
   fmt.Println(result) // printing the result of the recursive function
}

*/


/*
// syntax for defining a struct in Go
type Person struct { // defining a struct named Person
  name string 
  age int
  job string
  salary int
}

func main() {
  var pers1 Person // declaring a variable of type Person
  var pers2 Person

  // Pers1 specification
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 6000

  // Pers2 specification
  pers2.name = "Cecilie"
  pers2.age = 24
  pers2.job = "Marketing"
  pers2.salary = 4500

  // Access and print Pers1 info
  fmt.Println("Name: ", pers1.name)
  fmt.Println("Age: ", pers1.age)
  fmt.Println("Job: ", pers1.job)
  fmt.Println("Salary: ", pers1.salary)


*/

/*
// creating a map in GO and examples of how to use it
var a = make(map[KeyDataType]ValueDataType)
b := make(map[KeyDataType]ValueDataType)

var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"} // one method we can directly intialize a map with values 
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}


var a = map[string]string{} // we can also declare a map without initializing it with values
a["brand"] = "Ford" // we can also add values to a map after declaration also know as inserting
a["brand"] = "Toyota" // we can also update the value of a key in a map
a["model"] = "Mustang" // we the key is not present in the map, it will be added to the map


// to create a empty map using make function
b := make(map[string]int) // we can also create a map using the make function

// to access a value in a map
fmt.Println(a["brand"]) // we can access a value in a map using the key

// to check if a key is present in a map
 value, ok := a["brand"] // ok is a boolean value that will be true if the key is present in the map, otherwise false

// to delete a key-value pair from a map
delete(a, "brand") // we can delete a key-value pair from a map using the delete function

//maps are referance types
 var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := a // b is a reference to a
  fmt.Println(a)
  fmt.Println(b)
  b["year"] = "1970"// we are changing the values in the key and the original array is also changing that why it was referance type
  fmt.Println("After change to b:")
  fmt.Println(a)
  fmt.Println(b)

// we can also use a map as a struct field
type Person struct {
   name   string
   age    int
   job    string
   salary int
   skills map[string]int
}  

// we can iterate over a map using the for range loop
for key, value := range a {
   fmt.Println("Key:", key, "Value:", value)
}

// we can also use the for range loop to iterate over a map
for key := range a {
   fmt.Println("Key:", key)
}

// we can iterate over a map based on the order of the keys 
// basically they will coming in random order so we need to sort the keys first then iterate over the map becuase were accessing the map using keys.
var keys []string
for key := range a {
   keys = append(keys, key)
}
sort.Strings(keys) // sorting the keys
for _, key := range keys {
   fmt.Println("Key:", key, "Value:", a[key])
}


*/

//fmt.Println("Hello World!");

//}
