package main

import (
	"fmt"
	"os" 
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("error: Argument Eingeben")
		os.Exit(0)
	}

	text := os.Args[1]

	for v := 0; v < len(text); v++ {
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println(text)

	for v := 0; v < len(text); v++ {
		fmt.Printf("-")
	}

	printWaldhaus() // Aufrufen von Methode

}
func printWaldhaus() {

	waldhaus := `  
	                    \
	                     \ 
                              \
                               \ 
	                        \
	                         \                           [*]
					                   /*****\
					                 /**********\ 
	                  /.\ | |                       / ********** \
	                 /   \| |                      /**************\             
                        /     \ |           __        /****************\
                       /  (O)  \           /**\      / **************** \
                      /_________\        /******\    /___  ********   ___\
                      |         |        \******/      /   ********    \
                      |     [O] |         \****/      /___********** ___\
                      |         |          \**/          /___________\       
                      | []      |           ||                ||    
                      | []      |           ||                ||
||||||||||||||||||||| -----------|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||		           
			       
			   `

	fmt.Println(waldhaus)
}
