package main

import "fmt"

func main(){


for i := 1; i <= 5; i++ {
    fmt.Println(i)
}


count := 1

for count <= 5 {
    fmt.Println(count)
    count++
}


for i := 1; i <= 5; i++ {
    if i == 3 {
        continue
    }

    fmt.Println(i)
}



}