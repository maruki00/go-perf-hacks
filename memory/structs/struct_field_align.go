package structs

type BadStruct struct {
	Name    string
	Age     int
	Status  bool
	Sex     byte
	History []string
}

type GoodStruct struct {
	History []string
	Name    string
	Age     int
	Status  bool
	Sex     byte
}
