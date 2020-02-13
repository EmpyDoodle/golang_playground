package main

import (
	"fmt"
	"time"
  "math"
  "math/rand"
	"runtime"
)

func zero_vals() {  // var declaration - prints zero (empty) values
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q \n", i, f, b, s)  // Printf allows var substitution in order
  fmt.Printf("%T %T %T %T \n", i, f, b, s)  // Lots of sub syntax to play with
}

func time_is() {
  fmt.Println("The time is", time.Now(), "so do some work")  // Println is linear
}

func combine(sx string, sy string) (string) {  // Long var declaration, one str output
	return sx + sy                               // Explicit return var
}

func swap(sx, sy string) (string, string) {  // Shorthand var dec, 2 str output
  return sy, sx
}

var x, y int = rand.Intn(69), rand.Intn(420)  // Global var dec

func sqrt_sum(n1, n2 int) (output float64) {  // Output var named
  n := (n1 + n2)                              // short temp var - takes type from value (int in this case)
	f := float64(f)                             // can convert types using type as method
  output = math.Sqrt(f)                       // assign val to output var
  return                                      // 'naked' return gives named output vars
}

func rand_prob_gen() {
  fmt.Printf("You have %g problems\n", sqrt_sum(x, y))  // Printf does not carriage return
}

const Truth = true  // Constant declared like vars - also get type from value
const (  // vars can be set in blocks too
  Small = 1
  BigInt = Small << 32 - 1      // Shift the 1 left 32 places - biggest int
  BigUInt uint = Small << 64 -1  // uint needed for 64bit integers
  BigNum complex64 = Small << 100  // Too big to be int - needs complex
)

func loop(loops int) {
	sum := 1
	// All loops are "for" loops, same syntax as C / JS: [init statement ; condition expr ; post statement]
	//   [run before 1st iteration ; eval before each itr (kill if) ; run after each iteration]
  for i := 0; i < loops; i++ {
		fmt.Println("Sum is", sum)
		sum += sum
	}
  for sum < 420 {  // init and post statements optional - becomes a 'while' loop
    fmt.Println("Sum is", sum)
		sum += sum
	}
	fmt.Println("Final Sum is", sum)
}

func zibby() {
	for {  // for loops without statements just loop forever (loop do)
		fmt.Println("Zibby is gay")
	}
}

func total_gt10(n1, n2 int) {
	if t := (n1 + n2); t > 10 {  // if statements can also have optional init statement before condition eval expr
		fmt.Println("Your numbers total more than 10!")
	} else {
		fmt.Println("Doesn't quite add up:", t)
	}  // Can use the init statement var in the if/else statement, but not outside
}

func os() {
	fmt.Print("You are running ")  // Print by itself just does part of a line
	switch system := runtime.GOOS; os {  // switch is a contained if/else, similar to case statement
	case "darwin":                       // returns the first case that matches
		fmt.Println("OSX... Really?")  // Printf will print this + whatever has been Print-ed previously in func
	case "linux":
		fmt.Println("Linux, you clever chap or lady.")
	default:                         // if no cases match (else) ...
		fmt.Println("some unknown OS - probably Windows. More like Win-doze.")
	}
}

func is_saturday() {
	today := time.Now().Weekday()
	fmt.Println("Today is", today)
	switch {  // Switches can be left blank - clean if/else/then statements
	case int(today) < 6:  // if so, use case as you would "if"
		fmt.Println("Too early!")
	case int(today) > 6:
		fmt.Println("Too late...")
	case int(today) == 6:
		fmt.Println("HAPPY SATURDAY!!!")
	}
}

func main() {  // 'main' func is run when script executed - call other functions from this
  time_is()
  rand_prob_gen()
  str1, str2 := "poo", "wee"
	fmt.Println(combine(swap(str1, str2)))
  fmt.Println("BigInt is", BigInt, "BigUInt is", BigUInt, "BigNum is", BigNum)
  fmt.Printf("BigInt is %T, BigUInt is %T, BigNum is %T \n", BigInt, BigUInt, BigNum)
}
