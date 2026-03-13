package main
import "fmt"
func main() {
    var name string 
    var age int
    fmt.Scan(&name, &age)

    var r string
    if age < 12 {
        r = "crianca"
    } else if age < 18 {
        r = "jovem"
    } else if age < 65 {
        r = "adulto"
    } else if age < 1000 {
        r = "idoso"
    } else {
        r = "mumia"
    }

    fmt.Printf("%s eh %s\n", name, r)
}
