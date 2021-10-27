import (
    "fmt"
)
 
// declare function type
type FoodNameGetter func(string) string
 
type Food struct {
    name string
    getter FoodNameGetter // declare function
}
 
 
func main() {
    pizza := Food{
        name: "Pizza",
        getter: func(name string) string {         // declare function body
            return "This is " + name + "."
        },
    }
     
    fmt.Println(pizza.getter(pizza.name))     // prints "This is Pizza."
}
