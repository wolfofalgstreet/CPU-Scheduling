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


// ------------------------------------------------- //
// Get processes information needed for scheduling   //
// Return all precesses as array of process structs  //
func qProcesses(algorithm int, input string)([]process) {

  // Reading file and close when done
  file, err := os.Open(input)
  checkErr(err)
  defer file.Close()

  // Scanner buffer
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  // Storage
  var lines []string
  var procs []process
  var proc *process

  // Scanning lines
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  // Extract processes info onto struct array
  for i := 0; i < len(lines) - 1; i = i + 1 {

    // Will only extract Round Robin format
    if algorithm == 3 && i > 3 {
      proc = extractInfo(lines[i])
      procs = append(procs, *proc)

    // Will extract anything else
    } else if algorithm != 3 && i > 2 {
      proc = extractInfo(lines[i])
      procs = append(procs, *proc)
    }
  }
  return procs
}

// ----------------------------------------------- //
// Returns a process struct from string            //
// No need to search since input is structured as: //
// "process name ID arrival # burst #"             //
func extractInfo(line string) (*process) {

  // Allocating memory for process' info extraction
  proc := new(process)
  var words []string

  // Splitting string
  words = strings.Split(line, " ")

  // Parse info to struct for lines with > 7 words
  if len(words) > 6 {
    fmt.Println("word: ", words, "len: ", len(words))
    proc.ID = words[2]
    proc.arrival = toInteger(words[4])
    proc.burst = toInteger(words[6])
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

  // Read file and close it when done
  file := "c10-fcfs.in"
  input, err := os.Open(file)
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


  ss := qProcesses(algorithm, file)
  ss = ss
  fmt.Println("ID: ",ss[9].ID)



}
