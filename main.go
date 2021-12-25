package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	result1, err := doSomething()
	if err != nil {
		handleErrorSomehow(err)
		return
	}

	result2, err := doSomeMore(result1)
	if err != nil {
		handleErrorSomehow(err)
		return
	}

	result3, err := justOneMoreThing(result2)
	if err != nil {
		handleErrorSomehow(err)
		return
	}
	fmt.Println(result3)

	////////////////////////
	// With Result:
	////////////////////////

	resultA := doSomethingInAFunctionalWay()
	resultB := ExecuteOnSuccess(resultA, doSomeMoreInAFunctionalWay)
	resultC := ExecuteOnSuccess(resultB, justOneMoreThingInAFunctionalWay)

	if resultC.err != nil {
		handleErrorSomehow(resultC.err)
		return
	}
	print(resultC.value)
}

func doSomethingInAFunctionalWay() Result[int] {
	fmt.Println("doing something")
	return Ok(42)
}

func doSomeMoreInAFunctionalWay(value int) Result[[]byte] {
	fmt.Println("doing something more")
	if value == 1 {
		return Err([]byte(""), errors.New("blub"))
	}
	return Ok([]byte(strconv.Itoa(value)))
}

func justOneMoreThingInAFunctionalWay(value []byte) Result[string] {
	fmt.Println("doing one last thing")
	return Ok(string(value))
}

func doSomething() (int, error) {
	return 42, nil
}

func doSomeMore(value int) ([]byte, error) {
	return []byte(strconv.Itoa(value)), nil
}

func justOneMoreThing(value []byte) (string, error) {
	return "Result: " + string(value), nil
}

func handleErrorSomehow(err error) {
	fmt.Println("Error: " + err.Error())
}
