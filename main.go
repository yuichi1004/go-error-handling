package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Pattern 1: Creates on initialize

var ErrNeedCake = errors.New("error - i need cake not this")

func Func1(name string) error {
	if name != "cake" {
		return ErrNeedCake
	}
	return nil
}

// Pattern 2: Creates ad-hok

func Func2(name string) error {
	if name != "cake" {
		return fmt.Errorf("error - i need cake not %s", name)
	}
	return nil
}

// Pattern 3: Define custom error object

type ErrNeedElse struct {
	Got  string
	Need string
}

func (e ErrNeedElse) Error() string {
	return fmt.Sprintf("error - i need %s not %s", e.Need, e.Got)
}

func Func3(name string) error {
	if name != "cake" {
		return ErrNeedElse{name, "cake"}
	}
	return nil
}

// Pattern 4: Use Error interface not concrete struct

func Func4(name string) *ErrNeedElse {
	if name != "cake" {
		return &ErrNeedElse{name, "cake"}
	}
	return nil
}

type GiftToYou struct {
	Name string `json:name`
}

func (g *GiftToYou) UnmarshalJSON(data []byte) error {
	type alias *GiftToYou
	var this alias = alias(g)
	json.Unmarshal(data, &this)

	return Func4(g.Name)
}

func main() {
	fmt.Println("\n== Pattern1: Creates on initialize ==")
	if err := Func1("chocolate"); err != nil {
		if err == ErrNeedCake {
			fmt.Printf("Hmm, he seems to need a cake\n")
		} else {
			fmt.Printf("Got unexpected error!\n")
			panic(err)
		}
	}

	fmt.Println("\n== Pattern 2: Creates ad-hok ==")
	if err := Func2("banana"); err != nil {
		fmt.Printf("Func2() got error: %v\n", err)
	}

	fmt.Println("\n== Pattern 3: Define custom error object ==")
	if err := Func3("strawberry"); err != nil {
		switch err := err.(type) {
		case ErrNeedElse:
			fmt.Printf("Hmm, he seems to need %s not %s\n", err.Need, err.Got)
		default:
			fmt.Printf("Got unexpected error!\n")
			panic(err)
		}
	}

	fmt.Println("\n== Pattern 4: Use Error interface not concrete struct ==")
	v := GiftToYou{}
	if err := json.Unmarshal([]byte(`{"name":"cake"}`), &v); err != nil {
		switch err := err.(type) {
		case *ErrNeedElse:
			fmt.Printf("Hmm, he seems to need %s not %s\n", err.Need, err.Got)
		default:
			fmt.Printf("Got unexpected error!\n")
			panic(err)
		}
	}
}
