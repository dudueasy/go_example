package main

func main() {
	var i = 100
	var ip *int
	ip = &i

	println(i)
	println(*ip)

	*ip = 200
	println(*ip)

	add(&i)
	println(i)
}

func add(ip *int) {
	*ip += 1
}
