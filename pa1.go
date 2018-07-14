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
  //"sort"
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
  pnum int
  ID string
  arrival int
  burst int
  wait int
  turnaround int
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
  count := 0

  // Scanning lines
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  // Extract processes info onto struct array
  for i := 0; i < len(lines) - 1; i = i + 1 {

    // Will only extract Round Robin format
    if algorithm == 3 && i > 3 {
      proc = extractInfo(lines[i], count)
      procs = append(procs, *proc)
      count = count + 1

    // Will extract anything else
    } else if algorithm != 3 && i > 2 {
      proc = extractInfo(lines[i], count)
      procs = append(procs, *proc)
      count = count + 1
    }
  }
  return procs
}

// ----------------------------------------------- //
// Returns a process struct from string            //
// No need to search since input is structured as: //
// "process name ID arrival # burst #"             //
func extractInfo(line string, pnum int) (*process) {

  // Allocating memory for process' info extraction
  proc := new(process)
  var words []string

  // Splitting string
  words = strings.Split(line, " ")

  // Parse info to struct for lines with 7 words or more
  if len(words) > 6 {
    proc.pnum = pnum
    proc.ID = words[2]
    proc.arrival = toInteger(words[4])
    proc.burst = toInteger(words[6])
  }
  return proc
}


// ------------------------------------------------- //
// Will sort the structs by arrival and return procs //
func sortArrivals(procs []process)([]process) {

  // Bubble sort the structs by arrival time
  for x := 0; x < len(procs); x = x + 1 {
    for i := 0; i < len(procs) - 1; i = i + 1 {
      if procs[i].arrival > procs[i + 1].arrival {
        procs[i], procs[i + 1] = procs[i + 1], procs[i]
      }
    }
  }

  return procs
}

// ------------------------------------------ //
// Calculate waiting time and turnaround time //
func schedulePerformace(procs []process, selecTime, finished []int)() {

  var wt []int            // Storage for weight times
  var tt []int            // Storage for turnaround times
  proc := new(process)    // Process created
  var times = make([]process, len(procs))    // Array of processes created

  // Calcualte wait times
  for i := 0; i < len(procs); i = i + 1 {
    wt = append(wt, selecTime[i] - procs[i].arrival)
  }

  // Calculate turnaround times
  for i := 0; i < len(procs); i = i + 1 {
    tt = append(tt, finished[i] - procs[i].arrival)
  }

  // Assign performances and sort by increasing process ID
  for i := 0; i < len(procs); i = i + 1 {
    proc.ID = procs[i].ID
    proc.arrival = procs[i].arrival
    proc.wait = wt[i]
    proc.turnaround = tt[i]
    times[procs[i].pnum] = *proc
  }

  // Output schedule performace results
  for i := 0; i < len(times); i = i + 1 {
    fmt.Println(times[i].ID, " wait ", times[i].wait, " turnaround ", times[i].turnaround)
  }
}


// -------------------------------- //
// Execute FCFS Scheduling Algotihm //
func runFCFS(proccessNum, runFor, quantum int, procs []process, file string)() {

  time := 0           // Clock
  run := true         // Starts/Terminates algorithm
  proc := 0           // Index of current process
  busy := false       // True indicates a process is running
  finish := 0         // Index of process that is running
  began := 0          // Flag tracks when in time the process started
  var selecTime []int // Saves order of when a process was selected
  var finTime []int   // Saves order of when a process finished


  for run {

    // Process arrived
    for i := 0; i < len(procs); i = i + 1 {
      if time == procs[i].arrival {
        fmt.Println("Time ", time, " : ", procs[i].ID, " arrived")
        break
      }
    }

    // Process Finished
    if (time == began + procs[finish].burst) && busy {
      fmt.Println("Time ", time, " : ", procs[finish].ID, " finished")
      busy = false
      finTime = append(finTime, time)
      if proc < len(procs) {
        proc = proc + 1
      }
    }

    // Selecting process to run next
    if !busy && proc < len(procs){
      if time >= procs[proc].arrival {
        fmt.Println("Time ", time, " : ", procs[proc].ID, " selected (burst ", procs[proc].burst,")")
        busy = true
        finish = proc
        began = time
        selecTime = append(selecTime, time)
      }
    }

    // Processor in idle state
    if !busy && (procs[finish].arrival < time) {
      fmt.Println("Time ", time, " : ", " Iddle")
    }

    time = time + 1

    // Only run for given time
    if time == runFor {
      fmt.Println("Finished at time ", time, "\n")
      run = false
    }
  }

  // Print performance of schedule
  schedulePerformace(procs, selecTime, finTime)

}


// ------------------------------- //
// Execute SJF Scheduling Algotihm //
func runSJF(proccessNum, runFor, quantum int, procs []process)() {

}

// ------------------------------- //
// Execute RR Scheduling Algotihm //
func runRR(proccessNum, runFor, quantum int, procs []process)() {

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

  fmt.Println("pcount: ", processCount, " runfor: ", runFor, " quantum: ", quantum, "algorithm: ", algorithm, "\n")


  ss := qProcesses(algorithm, file)
  sortedArr := sortArrivals(ss)
  //fmt.Println("ID: ",ss[9].ID, " arrival: ", ss[9].arrival, " burst: ", ss[9].burst)

  //selected, finished := runFCFS(processCount, runFor, quantum, sortedArr)
  //schedulePerformace(sortedArr, selected, finished)
  runFCFS(processCount, runFor, quantum, sortedArr, file)

}
