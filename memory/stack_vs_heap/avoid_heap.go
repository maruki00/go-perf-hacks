package stackvsheap

func Bad1() *int {
	x := 1235456
	//some logic goes here
	//x goes into the heap
	return &x
}

func Bad2() func() int {
	x := 5
	//escape to the heap
	return func() int { return x }
}

func Bad3() interface{} {
	x := 1000
	//escape the heap when 
	//type is needed on the runtime
	return x
}


//this will bad on large struct 
//cuz its needed to copy it if 
//you pass it as value
func Good() int {
	x := 1235456
	//some logic goes here
	//x stays on the stack
	return x
}
