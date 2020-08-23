package IntroductionAlgorithms

import (
	"bufio"
	"fmt"
	"os"
)

func solution(line string) string  {
	var aws string
	return aws
}

func main()  {
	r := bufio.NewReaderSize(os.Stdin,20488)
	for line,_,err := r.ReadLine(); err == nil; line,_,err = r.ReadLine(){
		fmt.Println(solution(string(line)))
	}
}
