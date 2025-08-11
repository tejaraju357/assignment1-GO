package main


import "fmt"


type user struct {
	Name string
	Age int
	Email string
}

func main(){
	userData := make(map[string]user)

	for{
		fmt.Println("User appilication")
		fmt.Println("1. Add USer")
		fmt.Println("2. View user")
		fmt.Println("3. List All Users")
		fmt.Println("4. exit")

		var choice int
		fmt.Print("enter the number : ")
		fmt.Scan(&choice)
		fmt.Println()

		switch choice{
		case 1 : 
			var username,name,email string
			var age int

				fmt.Print("Enter the username: ")
				fmt.Scan(&username)
				fmt.Println()

				fmt.Print("Enter the name: ")
				fmt.Scan(&name)
				fmt.Println()

				fmt.Print("Enter the age: ")
				fmt.Scan(&age)
				fmt.Println()

				fmt.Print("Enter the email: ")
				fmt.Scan(&email)
				fmt.Println()

				userData[username] = user{Name : name,Age : age,Email : email}
		case 2 :
			  var searchItem string
			  fmt.Println("enter the username to view details of the user")
			  fmt.Scan(&searchItem)
			  exist, ok := userData[searchItem]
			  if ok {
				fmt.Println(exist)
			  }else{
				fmt.Println("user data not found...")
			  }

		case 3 :
			fmt.Println(userData)

		case 4 :
			fmt.Println("exit")
			return
			
		default:
			fmt.Println("invalid choice")
		}


	}
}
