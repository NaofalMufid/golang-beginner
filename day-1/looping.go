package topic

import "fmt"

func Hape(hape []string) {
	for i, val := range hape {
		fmt.Println("index", i, "isi =", val)
	}
}

func Horang(orang map[string]string) {
	for key, v := range orang {
		fmt.Println(key, "=", v)
	}
}
