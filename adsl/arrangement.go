// Arrangement
package adsl

import (
	"fmt"
)

func swap(str []string, a int, b int) {
	c := str[a]
	str[a] = str[b]
	str[b] = c
}

//k 表示当前选取到第几个数，m表示共有多少个数
func allRange(str []string, k int, m int) {
	if k == m {
		//i := 0
		//a := i + 1
		fmt.Println(str)
	} else {
		for i := k; i <= m; i++ {
			swap(str, k, i)
			allRange(str, k+1, m)
			swap(str, k, i)
		}
	}

}

/*func Readfile() {
	userfile := "passwd.txt"
	fl, err := os.Open(userfile)
	if err != nil {
		fmt.Println(userfile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		fmt.Println(string(buf[:n]))
	}

}*/

func Foo(str []string) {
	allRange(str, 0, len(str)-1)
}

func main6() {

	//list := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
	//"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	list := []string{"1", "2", "3"}

	Foo(list)

	//for i := 0; i < len(list); i++ {
	//fmt.Println(list[i])
	//}
}
