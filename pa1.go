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
    //fmt.Println("Conversion of var ", word, " failed!\n")
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
    //fmt.Println("Did not find ", item, " in memory..")
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
    //fmt.Println("Conversion of var ", word, " failed!\n")
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
  run int
}


// ------------------------------------------------- //
// Get processes information needed for scheduling   //
// Return all precesses as array of process structs  //
func qProcesses(algorithm int, input string)([]process) {

  // Read file and close when done
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

  var wt []int                                // Storage for weight times
  var tt []int                                // Storage for turnaround times
  proc := new(process)                        // Process created
  var times = make([]process, len(procs))     // Array of processes created

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
func runFCFS(proccessNum, runFor, quantum int, procs []process)() {

  time := 0              // Clock
  run := true            // Starts/Terminates algorithm
  proc := 0              // Index of current process
  busy := false          // True indicates a process is running
  finish := 0            // Index of process that is running
  began := 0             // Flag tracks when in time the process started
  var selecTime []int    // Saves order of when a process was selected
  var finTime []int      // Saves order of when a process finished

  fmt.Println(" ", proccessNum, " processes")
  fmt.Println("Using First-Come First-Served")

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
func runSJFbk(proccessNum, runFor, quantum int, procs []process)() {


  time := 0                    // Clock
  run := true                  // Flag terminates algorithm
  busy := false                // True indicates a process is running
  finish := 0                  // Tracks the process tthat was selected
  began := 0                   // Tracks when in time the process was selected
  proc := 0                    // Tracks the index of the current process
  inHold := false              // Flag indicates when a process preemptived
  partialProcSize := 0         //
  var currentProc, holdProc, minProc process
  //var mins []process

  fmt.Println(" ", proccessNum, " processes")
  fmt.Println("Using preemptive Shortest Job First")
  for run {

    // process arrived
    for i := 0; i < len(procs); i = i + 1 {
      if time == procs[i].arrival {
        fmt.Println("Time ", time, " : ", procs[i].ID, " arrived")

        // Check if new arrival has shorter burst than current
        if procs[i].burst < currentProc.burst {
          partialProcSize = time - currentProc.arrival
          //fmt.Println("Detected shorter burst")
          fmt.Println("Time ", time, " : ", procs[i].ID, " selected (burst ", procs[i].burst,")")
          holdProc = currentProc
          inHold = true
        }
        break
      }
    }

    holdProc = holdProc //
    partialProcSize = partialProcSize //
    minProc = minProc //

    // Process Finished
    if (time == began + procs[finish].burst) && busy {
      fmt.Println("Time ", time, " : ", procs[finish].ID, " finished")
      busy = false
      //finTime = append(finTime, time)
      if proc < len(procs) {
        proc = proc + 1
      }

      // Check if there is a partial process that needs to finish
      if inHold {

      }


    }


    // Selecting process to run next
    if !busy && proc < len(procs){
      if time >= procs[proc].arrival {
        //fmt.Println("Time ", time, " : ", procs[proc].ID, " selected (burst ", procs[proc].burst,")")
        busy = true
        finish = proc
        began = time
        currentProc = procs[proc]
        //selecTime = append(selecTime, time)
      }
    }


    // process idles state


    time = time + 1

    if time == runFor {
      fmt.Println("Finished at time ", time, "\n")
      run = false
    }
    finish = finish
    began = began
  }
}

// ------------------------------- //
// Execute SJF Scheduling Algotihm //
func runSJF(proccessNum, runFor, quantum int, procs []process)() {


  time := 0                    // Clock
  run := true                  // Flag terminates algorithm
  busy := false                // True indicates a process is running
  finish := 0                  // Tracks the process tthat was selected
  began := 0                   // Tracks when in time the process was selected
  proc := 0                    // Tracks the index of the current process
  //inHold := false              // Flag indicates when a process preemptived
  partialProcSize := 0         //
  minBurstIndex := 0           //
  var holdProc, minProc process
  burst := 0
  currentProc := new(process)
  //currIndex := 0

  fmt.Println(" ", proccessNum, " processes")
  fmt.Println("Using preemptive Shortest Job First")
  for run {

    // process arrived
    for i := 0; i < len(procs); i = i + 1 {
      if time == procs[i].arrival {
        fmt.Println("Time ", time, " : ", procs[i].ID, " arrived")

        // Find min burst of all arrivals
        for j := 0; j <= i; j = j + 1 {
          if procs[j].burst < procs[minBurstIndex].burst {
            minBurstIndex = j
          }
        }

        break
      }
    }

    holdProc = holdProc //
    partialProcSize = partialProcSize //
    minProc = minProc //
    burst = burst

    // Current process is always the min of arrivals
    // currentProc.ID = procs[minBurstIndex].ID
    // currentProc.arrival = procs[minBurstIndex].arrival
    // currentProc.burst = procs[minBurstIndex].burst
    //fmt.Println("current: ", currentProc.ID)



    // Process finished
    if currentProc.burst == 0 && busy {
      fmt.Println("Time ", time, " : ", currentProc.ID, " finished")

      //procs[minBurstIndex] = *currentProc
      //procs[minBurstIndex].ID = currentProc.ID
      //procs[minBurstIndex].arrival = currentProc.arrival
      //procs[minBurstIndex].burst = currentProc.burst

      if proc < len(procs) {
        proc = proc + 1
      }
      busy = false
    }

    // Selecting process to run Next
    if !busy && proc < len(procs) {
      if procs[minBurstIndex].burst > 0 {
        fmt.Println("Time ", time, " : ", currentProc.ID, " selected (burst ", currentProc.burst,")")
        busy = true

      }
      //fmt.Println("Time ", time, " : ", procs[minBurstIndex].ID, " selected (burst ", procs[minBurstIndex].burst,")")

    }

    // Running selected process
    if currentProc.burst > 0 {
      currentProc.burst = currentProc.burst - 1
      fmt.Println("Time ", time, " : ", currentProc.ID, " arrival: ", currentProc.arrival, "burst: ", currentProc.burst)
    }


    //Selecting process to run next
    // if !busy && proc < len(procs){
    //   if time >= procs[proc].arrival {
    //     //fmt.Println("Time ", time, " : ", procs[proc].ID, " selected (burst ", procs[proc].burst,")")
    //     busy = true
    //     finish = proc
    //     began = time
    //     currentProc = procs[proc]
    //     //selecTime = append(selecTime, time)
    //   }
    // }


    // process idles state


    time = time + 1

    if time == runFor {
      fmt.Println("Finished at time ", time, "\n")
      run = false
    }
    began, finish = began, finish
  }
}

func remove(slice []*process, s int) []*process {
    return append(slice[:s], slice[s+1:]...)
}


// ----------------------------------------- //
// Create a process instance and fill fields //
func create(pnum, arrival, burst, wt, ta, run int, ID string )(*process) {
  var node *process
  node = new(process)

  // node.ID = proc.ID
  // node.arrival = proc.arrival
  // node.burst = proc.burst
  // node.wait = proc.wait
  // node.turnaround = proc.turnaround
  // node.run = proc.run
  node.pnum = pnum
  node.ID = ID
  node.arrival = arrival
  node.burst = burst
  node.wait = wt
  node.turnaround = ta
  node.run = run

  return node
}

// --------------------------------------- //
// Execute Round Robin Scheduling Algotihm //
func runRRbk(proccessNum, runFor, quantum int, procs []process)() {
  run := true                     // Initiates/Terminates algorithm
  busy := false                   // Inidicates when a process is running
  time := 0                       // Clock
  proc := 0                       // Number of process completed
  ranFor := 0                     // Tracks runtime of a process
  //var requestQuee []process       // Circular queue
  var currProc *process            // Current process at work
  currProc = new(process)
  var requestQuee = make([]process, len(procs))
  size := 0


  passedby := 0

  // Clock
  for run {

    // Process Arrival
    for i := 0; i < len(procs); i = i + 1 {
      if time == procs[i].arrival {
        fmt.Println("Time ", time, " : ", procs[i].ID, " arrived")

        //requestQuee = append(requestQuee, procs[i])
        requestQuee[size] = procs[i]
         //curr process cant be set everytime a proc arrives!
        size = size + 1

        for i := 0; i < len(requestQuee); i = i + 1 { //
          //fmt.Println("reqQ[", i, "]: ", requestQuee[i].ID, "burst: ", requestQuee[i].burst)
        }

        //fmt.Println("appended: ", requestQuee[size].ID, "s: ", size, "curr: ", currProc.ID)
        //fmt.Println("curr: ", currProc.ID, " run: ", currProc.run, " burst: ", currProc.burst)
        //fmt.Println("Added by arrival: ", procs[i].ID, "rqSize: ", len(requestQuee)) //db
        break
      }
    }

    *currProc = requestQuee[size - 1]

    // Get the first element of queue
    // currProc.ID = requestQuee[0].ID
    // currProc.arrival = requestQuee[0].arrival
    // currProc.burst = requestQuee[0].burst
    // currProc.run = requestQuee[0].run
    //fmt.Println("curr: ", currProc.ID, " run: ", currProc.run, " burst: ", currProc.burst)

    //*currProc = requestQuee[0]
    //fmt.Println("curr: ", currProc.ID, " run: ", currProc.run)
    //currProc.run = currProc.run + 1
    //fmt.Println("curr: ", currProc.ID, " run: ", currProc.run)
    //currProc.run = time
    //fmt.Println("curr: ", currProc.ID, " run: ", currProc.run)


    //fmt.Println("about to hit finish")
    //fmt.Println("t: ", time," curr: ", currProc.ID, " run: ", currProc.run) //db

    // Check if Process Finished
    if busy && (currProc.run == currProc.burst) && (currProc.run != 0) {
      fmt.Println("Time ", time, " : ", currProc.ID, " finished: ", currProc.burst - currProc.run)

      // Remove it from request queue
      if size > 0 {
        size = size - 1
      }
      //remove(requestQuee,  0)

      // Processor availabe
      busy = false

      // Tracking number of completed processes
      if proc < len(procs) {
        proc = proc + 1
      }
    }

    // Process Selected
    if  proc < len(procs) {

         // Processor is free
         if !busy {

           // Select next process in queue
           *currProc = requestQuee[0];
           fmt.Println("Time ", time, " : ", currProc.ID, " selected (burst ", currProc.burst,")")
           busy = true
         }


         // Time Quantum
         // Processor can be free or working
         if time % quantum == 0 {

           // Getting next process in queue
           *currProc = requestQuee[0]

           // If we get here, it means that currProc has not finished
           // save progress and add to queue
           requestQuee[size] = *currProc
           size = size + 1

           // Checking next process is a valid process to run
           if requestQuee[0].arrival <= time {

             // select it as the new process to run
             *currProc = requestQuee[0]
             fmt.Println("Time ", time, " : ", currProc.ID, " selected (burst ", currProc.burst,")")
           }


           // Check if it has run for completion
           // If so it was removed from quee already just select new one
           if currProc.run == currProc.burst && currProc.run != 0 {
             //if size < len(procs) {//

             //}
             //remove(requestQuee, size) //
             //currProc = &requestQuee[0]
             // currProc.ID = requestQuee[0].ID
             // currProc.arrival = requestQuee[0].arrival
             // currProc.burst = requestQuee[0].burst
             // currProc.run = requestQuee[0].run
             //fmt.Println("Time s", time, " : ", currProc.ID, " selected (burst ", currProc.burst,")")

           // If not save progress and re-add to quee
           } else {
             //remove(requestQuee, 0)
             //requestQuee = append(requestQuee, *currProc) //

           }

    }


      //fmt.Println("quantum time: ", time, "currProc: ", currProc.ID)
    }



    // Process Idle




    currProc = currProc



    // Clock and runtime updates
    time = time + 1
    ranFor = ranFor + 1
    currProc.run = currProc.run + 1
    passedby = passedby + 1 //
    //fmt.Println(passedby, "currRun: ", currProc.run, " from: ", currProc.ID)

    // Terminate Algorithm
    if time == runFor {
      fmt.Println("Finished at time ", time, "\n")
      run = false
    }

  }
}

// --------------------------------------- //
// Execute Round Robin Scheduling Algotihm //
func runRR(proccessNum, runFor, quantum int, procs []process)() {
  run := true                     // Initiates/Terminates algorithm
  busy := false                   // Inidicates when a process is running
  time := 0                       // Clock
  proc := 0                       // Number of process completed
  ranFor := 0                     // Tracks runtime of a process
  var requestQuee []*process       // Circular queue
  var currProc *process            // Current process at work


  // Clock
  for run {

    // Process Arrival
    for i := 0; i < len(procs); i = i + 1 {
      if time == procs[i].arrival {
        fmt.Println("Time ", time, " : ", procs[i].ID, " arrived")


        // currProc = create(requestQuee[0].pnum, requestQuee[0].arrival,
        // requestQuee[0].burst, requestQuee[0].wait, requestQuee[0].turnaround,
        // ranFor, requestQuee[0].ID)

        requestQuee = append(requestQuee, &procs[i]) //
        //requestQuee = append(requestQueem, currProc)

        //requestQuee[size] = procs[i] //
        //size = size + 1//

        for i := 0; i < len(requestQuee); i = i + 1 { //
          //fmt.Println("reqQ[", i, "]: ", requestQuee[i].ID, "burst: ", requestQuee[i].burst)
        }

        break
      }
    }

    // currProc = create(requestQuee[0].pnum, requestQuee[0].arrival,
    // requestQuee[0].burst, requestQuee[0].wait, requestQuee[0].turnaround,
    // ranFor, requestQuee[0].ID)

    //currProc = create(requestQuee[0])

    // Check if Process Finished
    if  (currProc.run == currProc.burst) {
      //fmt.Println("Time ", time, " : ", currProc.ID, " finished: ", currProc.burst - currProc.run)

      // Dequeue first process
      requestQuee = requestQuee[1:]

      // Create instance of new process
      currProc = create(requestQuee[0].pnum, requestQuee[0].arrival,
      requestQuee[0].burst, requestQuee[0].wait, requestQuee[0].turnaround,
      requestQuee[0].run, requestQuee[0].ID)

      // Processor availabe
      busy = false
      ranFor = currProc.run

      // Tracking number of completed processes
      if proc < len(procs) {
        proc = proc + 1
      }
    }

    // Select Process
    if  proc < len(procs) {

      // Time Quantum
      // Processor can be free or working
      if time % quantum == 0 {

        // If we get here, it means that currProc has not finished
        // save progress and re-add to queue
        if currProc.run != currProc.burst {
          currProc = create(currProc.pnum, currProc.arrival,
          currProc.burst, currProc.wait, currProc.turnaround,
          ranFor + 1, currProc.ID)

          // Enqueue process
          requestQuee = append(requestQuee, currProc)

          // Dequeue process
          requestQuee = requestQuee[1:]

          // Update currProc
          currProc = create(requestQuee[0].pnum, requestQuee[0].arrival,
          requestQuee[0].burst, requestQuee[0].wait, requestQuee[0].turnaround,
          requestQuee[0].run, requestQuee[0].ID)



        }
        //fmt.Fprintf("this is)%3d")
        // Processor is free
        if  currProc.arrival <= time {

          // Select next process in queue
          //*currProc = requestQuee[0];
          //currProc = create(requestQuee[0])
          //fmt.Println("Time ", time, " : ", currProc.ID, " selected (burst ", currProc.burst,")")
          //fmt.Println("Time ", time, " : ", currProc.ID, " selected (burst ", currProc.burst,")", "run: ", currProc.run)
          busy = true
        }


        for i:=0;i < len(requestQuee); i = i + 1 {
          //fmt.Println("i: ", i, " ",requestQuee[i].ID, "run: ", requestQuee[i].run )
        }


        //requestQuee = append(requestQuee, currProc)

        //requestQuee = requestQuee[1:] //

        // // Checking next process is a valid process to run
        // if requestQuee[0].arrival <= time {
        //
        //   // select it as the new process to run
        //   //currProc = create(requestQuee[0])
        //   currProc = create(requestQuee[0].pnum, requestQuee[0].arrival,
        //   requestQuee[0].burst, requestQuee[0].wait, requestQuee[0].turnaround,
        //   requestQuee[0].run, requestQuee[0].ID)
        //   fmt.Println("selencted from quantum")
        //   fmt.Println("Time ", time, " : ", currProc.ID, " selected (burst ", currProc.burst,")")
        //   busy = true
        // }
      }


    }



    // Process Idle


    // Runtime Updates
    time = time + 1
    ranFor = ranFor + 1
    fmt.Println("Ran: ", currProc.run, "ranFor: ",ranFor, " ID: ", currProc.ID)
    currProc = create(currProc.pnum, currProc.arrival,
    currProc.burst, currProc.wait, currProc.turnaround,
    ranFor, currProc.ID)
    fmt.Println("curr.run: ", currProc.run)
    currProc = currProc
    busy = busy

    // Terminate Algorithm
    if time == runFor {
      fmt.Println("Finished at time ", time, "\n")
      run = false
    }

  }
}



func main () {

  // Read file and close it when done
  file := "c5-rr.in"
  //file := os.Args[1]
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

  ss := qProcesses(algorithm, file)
  sortedArr := sortArrivals(ss)


  // Execute selected scheduling algorithm
  switch algorithm {
  case 1:                               // First-Come First-Served
    runFCFS(processCount, runFor, quantum, sortedArr)

  case 2:                               // Shortest-Job First
    runSJF(processCount, runFor, quantum, sortedArr)

  case 3:                               // Round Robin
    runRR(processCount, runFor, quantum, sortedArr)
  }


  // DEBUGGING //
  processCount, runFor, quantum, algorithm = processCount, runFor, quantum, algorithm
  //fmt.Println("pcount: ", processCount, " runfor: ", runFor, " quantum: ", quantum, "algorithm: ", algorithm, "\n")



}
