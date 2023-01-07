package main

import (
	"fmt"
	"os"
)

func interactive_mode() {
	fmt.Println("Welcome to snnt_notes! Still under development!")
}

func program_was_invoked_without_args() bool {
	return len(os.Args) == 1
}

func first_arg_is_meta_arg() bool {
	if program_was_invoked_without_args() {
		panic(1)
	}
	return os.Args[1] == "--arg"
}

func handle_args() {
	if !first_arg_is_meta_arg() {
		panic(2)
	}

	fmt.Println("Arg handling: snnt_notes was invoked with these args:")

	for index, arg := range os.Args {
		fmt.Printf("Index %v: %v\n", index, arg)

	}
}

func direct_mode() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
	var input string
	fmt.Scanln(&input)

}

func main() {

	if program_was_invoked_without_args() {
		interactive_mode()
	} else if first_arg_is_meta_arg() {
		handle_args()
	} else {
		direct_mode()
	}

}
