package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func spend(a, b, c, d, e int) int {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	if a > 0 {
		//  众所周知水在低于 0 摄氏度的时候是固态，在高于 0 摄氏度的时候是液态，在等于 0 摄氏度的时候是可能是液态也可能是固态还可能是固液混合
		// 现在有个加热工具，加热 c 秒可以使冰升温 1 摄氏度。
		// 加热 d 秒可以使 0 摄氏度的冰变成 0 摄氏度的水。
		// 加热 e 秒可以使水升温 1 摄氏度。
		// 现在给出两个温度 a , b ，请计算用这个加热工具将水从 a 摄氏度加热到 b 摄氏度所需要的时间。
		// 保证a不为零。
		return (b - a) * e
	} else {
		return (-a * c) + d + (b * e)
	}
}

func readLineFromStdin(lineCount int) (result []string) {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 5; i++ {
		text, _ := reader.ReadString('\n')
		fmt.Printf("%d", len(text))
		result = append(result, text[:len(text)-1])
	}
	return
}

func main() {
	input := readLineFromStdin(5)
	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])
	d, _ := strconv.Atoi(input[3])
	e, _ := strconv.Atoi(input[4])
	fmt.Println(spend(a, b, c, d, e))
}
