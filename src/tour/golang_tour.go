package main

import (
	"fmt"
	"time"
  "math"
  "math/rand"
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
  f := float64(n1 + n2)                       // short temp var - takes type from value
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

func main() {  // 'main' func is run when script executed - call other functions from this
  time_is()
  rand_prob_gen()
  str1, str2 := "poo", "wee"
	fmt.Println(combine(swap(str1, str2)))
  fmt.Println("BigInt is", BigInt, "BigUInt is", BigUInt, "BigNum is", BigNum)
  fmt.Printf("BigInt is %T, BigUInt is %T, BigNum is %T \n", BigInt, BigUInt, BigNum)
}
