// Created By Isaias Perez Vega
// NID: is352549
// CPU - Scheduling Algorithms
// Program implements the first-come firs-served, shortest-job first,
// and round robin algorithms.
//

package main


import (
  "bufio"
  "fmt"
  //"io"
  //"io/ioutil"
  "os"
)

// Streamlining error checking
func checkErr(e error) {
  if e != nil {
    fmt.Println("Error ocurred when trying to open file.\n")
    panic(e)
  }
}

//
func main () {

  // Reading in file
  //input , err := ioutil.ReadFile("c2-fcfs.base") //c2-fcfs.in
  input, err := os.Open("fcfs.in")
  checkErr(err)
  //defer input.Close()

  // Testing input read
  fmt.Println("The input is:\n", input)
  // Converting input from bytes to string
  str := string(input)
  fmt.Println("As string:\n", str)

  // Reading Line per Line
  fmt.Println("About to print line per line:\n")
  scanner := bufio.NewScanner(input)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }








}
