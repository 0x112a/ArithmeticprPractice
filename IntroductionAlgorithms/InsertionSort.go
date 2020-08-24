package IntroductionAlgorithms

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(line string) string {
	var aws string
	key := strings.Split(line, " ")
	n, _ := strconv.Atoi(key[1])
	fmt.Println(n)
	return aws
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20488)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}
