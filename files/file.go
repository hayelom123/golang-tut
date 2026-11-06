package main 

import  ("fmt"
"os")


func main(){
  
	file,err:=os.Create("example.txt")
	if err !=nil {
		fmt.Println("Error creating a file ",err)
		return
	}

	defer file.Close()

	fmt.Println("File created:",file.Name())
     
	file.WriteString("Jesus christ is God!!!!\n")
	file.WriteString("In the beginning was the Word, and the Word was with God, and the Word was God.(john 1:1)\n")
	file.WriteString("He was with God in the beginning.Through him all things were made; without him nothing was made that has been made.\n")
	file.WriteString("In him was life, and that life was the light of all mankind.\nThe light shines in the darkness, and the darkness has not overcome[a] it\n")
	file.WriteString("14 The Word became flesh and made his dwelling among us. We have seen his glory, \nthe glory of the one and only Son, who came from the Father, full of grace and truth.\n")

	 
}