
//Creating a file txt
//Writing inside the file txt
//Reading a file line by line

package main
import(
	"fmt"
	"log"
	"os"
	"bufio"
)
//Creating
func CreateFile() {
	file, er := os.Create("sampletexto.txt")
	if er != nil {
		log.Fatalf("failed to create file")
	}
	fmt.Printf("Creating file %s", file.Name())
	defer file.Close()

//Writing
	len, er := file.WriteString("Lorem ipsum dolor sit amet, consectetur adipiscing elit.\n" + "Phasellus eleifend elementum purus, at efficitur lorem rutrum id.\n" + "Mauris ut eros a augue viverra volutpat.\n" + "Vivamus porttitor urna sit amet diam molestie, sit amet aliquam elit viverra.\n")
	if er != nil {
		log.Fatalf("failed to write file")
	}
	fmt.Printf("\nSize is %d bytes\n", len)	
}
//Reading line by line
func ReadFile() {
	fmt.Printf("Reading the file line by line: \n")
	file, er := os.Open("sampletexto.txt")
	if er != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	for _, each_ln := range text {
		fmt.Println(each_ln)
	}
}

func main() {
	CreateFile()
	ReadFile()
}
