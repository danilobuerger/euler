package fibonacci

// Generator generates fibonacci numbers until closed
func Generator(done chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i, j := 0, 1; ; i, j = j, i+j {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()

	return out
}

// EvenGenerator generates even fibonacci numbers until closed
func EvenGenerator(done chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i, j := 2, 8; ; i, j = j, 4*j+i {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()

	return out
}
