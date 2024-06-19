package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	alex := Human{}
	alex.Name = "Alex"
	alex.Age = 25
	alex.Phone = "123-456-7890"
	alex.DateOfBirth = "1995-01-01"

	b, err := json.MarshalIndent(alex, "", "  ")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("alex.json", b, 0644)

}
