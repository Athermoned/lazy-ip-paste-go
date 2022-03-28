package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input() ([]int, string) {
	var ip string
	var command string
	fmt.Print("Введите IP: ")
	fmt.Scanln(&ip)
	fmt.Print("Введите желаемую команду с %s там, где нужен IP: ")
	lr := bufio.NewScanner(os.Stdin)
	if !lr.Scan() {
		panic("Ошибка в чтении комманды")
	}
	command = lr.Text()
	ipOctRaw := strings.Split(ip, ".")
	if len(ipOctRaw) != 4 {
		panic("Это не IP адрес")
	}
	ipOct := make([]int, len(ipOctRaw))
	for i, s := range ipOctRaw {
		ipOct[i], _ = strconv.Atoi(s)
	}
	return ipOct, command
}
func checkOc(ipOct []int) {
	for _, b := range ipOct {
		if b < 0 || b > 255 {
			panic("Октет не может быть больше 255")
		}
	}
}

func cidr24(ipOct []int, command string) {
	for cidr24oct := 0; cidr24oct < 255; cidr24oct++ {
		tmp := make([]string, 4)
		for k, v := range append(ipOct[0:2+1], cidr24oct) {
			tmp[k] = strconv.Itoa(v)
		}
		fmt.Printf(command+"\n", strings.Join(tmp, "."))
	}
}

func main() {
	m, c := input()
	checkOc(m)
	cidr24(m, c)

}
