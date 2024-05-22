package main

import (
	"my_module/Task1"
	"my_module/Task2"
	channels "my_module/Task3"

	"encoding/json"
	"fmt"
	"os"
)

type Input struct {
	Slc1 []int
	Text2 string
	Num int
}

type Answer struct {
	Sum int
	Titled string
	Unum int
}

func ReadFile() ([]Input,error) {
	data, err := os.ReadFile("input.json")
	if err != nil {
		panic(err)
	}

	var inputFile []Input
	err = json.Unmarshal(data, &inputFile)
	if err != nil {
		panic(err)
	}

	return inputFile, nil
}

func WriteFile(kirit Answer) {
	f,err := os.OpenFile("data.json", os.O_CREATE | os.O_WRONLY,0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := json.Marshal(kirit)
	if err != nil {
		panic(err)
	}
	f.Write(data)

}


func main()  {
	n,err:=ReadFile()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	javob:= Answer{}
	javob.Sum=summa.Add(n[0].Slc1)
	javob.Titled=letter.Letters(n[0].Text2)
	javob.Unum=channels.Nums(n[0].Num)
	WriteFile(javob)
	fmt.Println(javob)

	
}