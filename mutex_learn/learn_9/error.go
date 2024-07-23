package main

func main() {
	m := make(map[int]int, 10)

	go func() {
		m[1] = 1
	}()

	go func() {
		_ = m[1]
	}()
	select {}
}
