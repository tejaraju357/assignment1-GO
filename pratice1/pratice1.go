package main

import "fmt" 


type product struct{
		Name string
		Quantity int
}

func (p *product) increase(){
		p.Quantity++
}

func (p *product) decrease(){
		p.Quantity--
}

func (p *product) changeName(newName string){
		p.Name = newName
}

func main(){
	
Product := product{Name : "Mouse",Quantity: 10}
fmt.Println(Product)
Product.increase()
fmt.Println(Product)
Product.decrease()
fmt.Println(Product)
Product.changeName("keyboard")
fmt.Println(Product)

}
