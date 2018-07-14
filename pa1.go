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
  "strings"
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
  processCount = lookConvert("processcount", words)
  runFor = lookConvert("runfor", words)
  quantum = lookConvert("quantum", words)
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


// ------------------------- //
// Convert string to integer //
func toInteger(word string) (int) {
  num, err := strconv.Atoi(word)

  // Type conversion error handling
  if err != nil {
    fmt.Println("Conversion of var ", word, " failed!\n")
  }
  return num
}


// ---------------------------- //
// Search index of word in list //
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


// ------------------------------------------------ //
// Looks for a word in array and converts it to int //
func lookConvert(word string, words []string) (int) {
  index := find(word, words)
  num, err := strconv.Atoi(words[index + 1])

  // Type conversion error handling
  if err != nil {
    fmt.Println("Conversion of var ", word, " failed!\n")
  }

  return num
}


// ------------------------------------------- //
// Struct holds information about each process //
type process struct {
  ID string
  arrival int
  burst int
}


// ----------------------------------------------- //
// Get processes information needed for scheduling //
func qProcesses(algorithm int)([]process) {
  // Reading file
  file, err := os.Open("c10-rr.in")
  checkErr(err)
  defer file.Close()

  // Scanner buffer
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  // Storage
  var lines []string
  var procs []process

  //fmt.Println("About to read line by line:\n")

  counter := 0
  procLine := 0
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
    if algorithm == 3 {
      
    }
    fmt.Println(lines[counter], "\n")
    counter = counter + 1
  }
  fmt.Println("lINES[4]: ", lines[4]) //
  procLine = procLine
  return procs

}

// ----------------------------------------------- //
// Returns a process struct from string            //
// No need to search since input is structured as: //
// "process name ID arrival # burst #"             //
func extractInfo(line string) (process) {

  // Allocating resources
  var proc process
  var words []string

  // Splitting string
  words = strings.Split(line, " ")

  // Defining struct
  proc = process{ID: words[2], arrival: toInteger(words[4]), burst: toInteger(words[6])}

  for i := 0; i < len(words); i = i + 1 {
    fmt.Println("Index", i, ": ", words[i], "\n")
  }
  return proc
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
  input, err := os.Open("c2-fcfs.in") //c10-rr.in
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


  // DEBUGGING //
  processCount = processCount
  runFor = runFor
  quantum = quantum
  algorithm = algorithm

  fmt.Println("pcount: ", processCount, " runfor: ", runFor, " quantum: ", quantum, "algorithm: ", algorithm)

  var p = "hello this is not your average hello world"
  fmt.Println("About to print: ", p)
  str := strings.Split(p, " ")

  // Testing splitting text and accessing words
  fmt.Println("Splitting string:\n", strings.Split(p, " "))
  fmt.Println("Printing by index: \n", "Index 0: ", str[0], "\nIndex 6: ", str[5])
  if strings.Contains(p, "world") {
    fmt.Println("World was found at index: ", strings.Index(p, "world"))
  }

  ss := qProcesses(algorithm)
  ss = ss

  // Testing struct info
  var newproc process
  newproc = extractInfo("process name P2 arrival 0 burst 9")
  newproc = newproc
  fmt.Println("Reading process struct:\n")
  fmt.Println("ID: ", newproc.ID, " arrival: ", newproc.arrival, " burst: ", newproc.burst)

  // find use
    // check if rr
  // extract processes




}
