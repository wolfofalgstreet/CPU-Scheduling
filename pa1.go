// Created By Isaias Perez Vega
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
  "strconv"
  "os"
)

// --------------------------- //
// Streamlining error checking //
func checkErr(e error) {
  if e != nil {
    fmt.Println("Error ocurred when trying to open file.\n")
    panic(e)
  }
}


// ------------------------------------ //
// Gets scheduling algorithm setup info //
func setupInfo(words []string) (processCount, runFor, quantum, algorithm int) {
  fmt.Println("Runing setup:\n") // debug

  // Searching, converting, and assigning setup info
  processCount = toInt("processcount", words)
  runFor = toInt("runfor", words)
  quantum = toInt("quantum", words)
  algorithm = algoType(words)

  return processCount, runFor, quantum, algorithm
}

// ----------------------------------------- //
// Represent the type of algorithm as an int //
func algoType(words []string) (int) {
  index := find("use", words)
  algo := 0
  switch words[index + 1] {
  case "fcfs":
    algo = 1
  case "sjf":
    algo = 2
  case "rr":
    algo = 3
  }
  return algo
}

// ------------------------------- //
// Search location of word in list //
func find(item string, list []string) (index int) {
  count := 0
  index = -1
  for _,  word:= range list {
    if word == item {
      index = count
    }
    count = count + 1
  }

  // Before returning check if the word was found
  if index == -1 {
    fmt.Println("Did not find ", item, " in memory..")
  }

  return index
}

// --------------------- //
// Convert string to int //
func toInt(word string, words []string) (int) {
  index := find(word, words)
  num, err := strconv.Atoi(words[index + 1])

  // Type conversion error handling
  if err != nil {
    fmt.Println("Conversion of var ", word, " failed!\n")
  }

  return num
}

// -------------------------------- //
// Execute FCFS Scheduling Algotihm //
func runFCFS()() {

}

// -------------------------------- //
// Execute SJF Scheduling Algotihm //
func runSJF()() {

}

// -------------------------------- //
// Execute RR Scheduling Algotihm //
func runRR()() {

}


func main () {

  // Reading input file
  input, err := os.Open("c10-rr.in") //c10-rr.in
  checkErr(err)
  defer input.Close()

  // Scanning word by Word
  scanner := bufio.NewScanner(input)
  scanner.Split(bufio.ScanWords)

  // Words will be scanned into words[]
  var words []string
  for scanner.Scan() {
    words = append(words, scanner.Text())
  }

  // Checking contents of words[]
  count := 0
  fmt.Println("Words list:\n")
  for _, word := range words {
    count = count + 1
    fmt.Println("word: ",word, " count: ", count)
  }

  // Scheduling Algorithm setup info
  processCount, runFor, quantum, algorithm := setupInfo(words)

  // Execute selected scheduling algorithm
  switch algorithm {
  case 1:           // First-Come First-Served
    fmt.Println("Run: ", algorithm)
  case 2:           // Shortest-Job First
    fmt.Println("Run: ", algorithm)
  case 3:           // Round Robin
    fmt.Println("Run: ", algorithm)
  }




  processCount = processCount
  runFor = runFor
  quantum = quantum
  algorithm = algorithm

  fmt.Println("pcount: ", processCount, " runfor: ", runFor, " quantum: ", quantum, "algorithm: ", algorithm)




}
