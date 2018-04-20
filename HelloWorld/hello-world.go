package main

import ("fmt")

const abc string = "xxxxxx"

func main() {

	fmt.println("hello ramesh");



	// m := make(map[string]string)
	// m["first"] = "matha"
	// m["second"] = "pitha"
	// m["third"] = "guru"
	// m["fourth"] = "dhaivam"
	// fmt.Println(m)

	// _, prs := m["fiasdfrst"]
	// fmt.Println(prs)

	// fmt.Println(len(m))

	/*
		slicee := make([]string, 3)

		fmt.Println(slicee)

		slicee[0] = "ramesh"
		slicee[1] = "veeramani"
		slicee[2] = "bangalore"

		fmt.Println(slicee)

		A := slicee[2:3]

		B := append(A, "success")

		fmt.Println(B[:2])
		fmt.Println(A)
	*/ //fmt.Println(B)

	/* //b := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	const ii, jj int = 50, 50

	var zzz [ii][jj]int

	for i := 0; i < ii; i++ {
		for j := 0; j < jj; j++ {
			zzz[i][j] = i + j
		}
	}

	//	fmt.Println(zzz)

	//	fmt.Println(b)

	var a [2]int
	a[0] = 1212
	a[1] = 323

	//	fmt.Println(a)

	var str string = "hello world"
	var aa int
	aa = 1000

	//	fmt.Println(str)
	//	fmt.Println(aa)

	fmt.Println(abc)

	fmt.Println("This is going to make me a billionaire...")

	fmt.Println("This is " + "it")

	fmt.Println("1+1", 1+1)

	fmt.Println(true && false)
	fmt.Println(true && true)
	//fmt.Println(true & false)
	fmt.Println(true || false)

	i := 0

	if num := 9; num < 9 {
		println("less than 9")
	} else if num == 9 {
		println("well its 9")
	}
	for {
		i = i + 1
		println(i)

		if i > 10 {
			break
		}
	}

	whatami := func(i interface{}) {
		switch i.(type) {
		case int:
			println("its a boy")
		case (string):
			println("its a girl")
		}

	}

	whatami(1)
	whatami("ramesh")
	*/
}
