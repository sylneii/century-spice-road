package main

func main() {

	ch := make(chan int)

	go func() {
		<-ch
	}()

	// x := <-ch
	// fmt.Println(x)
}
