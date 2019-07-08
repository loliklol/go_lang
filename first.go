package main

import "fmt"

func average(xs []int) int {

	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}

func maxarg(arg ...int) {

	for _, v := range arg {
		println(v)
	}

}

func main() {
	///////работа с функциями

	someOtherName := average([]int{56, 34, 24, 43, 434, 343})
	maxarg(1, 2, 3, 4)

	println(someOtherName)
	test1 := make([]int, 5, 7)
	//append  -- добавление в слайс
	test2 := append(test1, 4, 5, 7)
	test3 := append(test1, test2...)
	fmt.Println(test3)
	println("!!!!!!!!!!!!!!!!!!")

	var bufLen, bufCap int = len(test3), cap(test3)
	println(bufLen, " ", bufCap)

	// fmt.Print("Enter a number: ")
	// var input float64
	// fmt.Scanf("%f", &input)

	// output := input * 2

	// fmt.Println(output)

	// var x [5]int

	// x[4] = 3

	// fmt.Println(x)
	// y := x[3:5]
	// b := x[0:1]
	// fmt.Println(b)
	// y = append(y, 5, 6, 8)
	// fmt.Println(y)

	var mapi = make(map[string]string)
	mapi["x"] = "10"
	mapi["y"] = "11"
	mapi["z"] = "12"
	mapi["k"] = "13"
	// fmt.Println(mapi["k"])
	// name, ok := mapi["k"]
	// fmt.Println(name)

	// Массивы
	for _, value := range mapi {

		fmt.Println(value)
	}

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["tesr"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	var testik = make([]int, 8, 9)
	fmt.Println(testik)

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17, 4, 5, 2, 1,
	}
	var itog int = x[0]

	for _, value := range x {
		//fmt.Println(value)
		//fmt.Println(itog)

		if itog > value {
			println("sadasdas")
			itog = value
		}

	}
	fmt.Println(itog)

	if _, keyExist := mapVal["name"]; keyExist {

		fmt.Println("key exist")
	}

	mapVa1 := map[string]string{"name": "vasiliy"}
	if keyValue, keyExist := mapVa1["name"]; keyExist {

		fmt.Println("name =", keyValue)
	}

	// t := 1
	// for t <= 15 {
	// 	println(t)
	// 	t = t + 1

	// }

	// for i := 1; i <= 16; i++ {
	// 	println(i)
	// 	if i == 12 {
	// 		fmt.Println("НОРМ ВСЕ")
	// 	} else if i == 3 {
	// 		fmt.Println("УРА")

	// 	} else {
	// 		fmt.Println("!!!!!!!!!!!!!!!!")
	// 	}

	// }
	// // i := 4
	// switch i {

	// case 0:
	// 	fmt.Println("Zero")
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// case 5:
	// 	fmt.Println("Five")
	// default:
	// 	fmt.Println("Unknown Number")
	// }

	// //Преобразование типов
	// fmt.Println(float64(len("#$@#$@#$@#$@#DEWQWWE")))

	// var total float64 = 0
	// for _, value := range x {
	// 	total += float64(value)
	// }
	// fmt.Println(total / float64(len(x)))

}
