# CPU-Scheduling
Program uses GO to read an input file containing several configuration parameters to simulate First-Come First-Served, preemptive Shortest Job First, and Round-Robin CPU scheduling algorithms. The output will reflect the results of the configured CPU scheduling algorithm to a specified output file.

## Algorithms    

### First-Come First-Served 
The First-Come First-Servedal gorithm is just that:  
processes are allocated the CPU in the order in which they arrive and run until completion or termination

### Preemptive Shortest Job First
The Preemptive Shortest Job First selects the process for execution which has the smallest amount of time remaining until completion.  Therefore, a process with a budget or expected time of 2 time units will be scheduled ahead of a process with 6 time units.

### Round-Robin 
Round-Robin was originally built for time-sharing systems. It works on the premise of providing a rigorous enforcement of an interrupt every configured time interval, then swapping to the next task in the process listat the moment of the interrupt.  This scheduling algorithm treats the process list as a circular list until each process is complete. New processes can be added infinitum. Note that the Round-Robin algorithm requires a configuration of the time quantum for the period of time used for the process to be in the CPU. Once the period expires the next process in the list will have CPU resources until its time quantum expires.

## Input Format
The input file will be passed as the first command line argument. Also the program ignores everything on a line after a '#' mark and ignore additional spaces in the input file.

Note: Blank lines are not considered for processing.

```
processcount 2 # Read 2 processes
runfor 15 # Run for 15 time units
use rr # Can be fcfs, sjf, or rr
quantum 2 # Time quantum, only if using rr
process name P1 arrival 3 burst 5
process name P2 arrival 0 burst 9
end
```

### Input Files
| File Name     | Description                                    |
| ------------- |:----------------------------------------------:| 
| c2-fcfs.in    | 2 processes scheduled *First-Come First-Served*|
| c2-rr.in      | 2 processes scheduled *Round-Robin*            |
| c2-sjf.in     | 2 processes scheduled *preemptive Shortest Job First* |
| c5-fcfs.in    | 5 processes scheduled *First-Come First-Served*|
| c5-rr.in      | 5 processes scheduled *Round-Robin*            |
| c5-sjf.in     | 5 processes scheduled *preemptive Shortest Job First* |
| c10-fcfs.in   | 10 processes scheduled *First-Come First-Served* |
| c10-rr.in     | 10 processes scheduled *Round-Robin* |
| c10-sjf.in    | 10 processes scheduled *preemptive Shortest Job First* |

## Command Line Inputs
The *command line inputs* or *cli* are the input file name as the first parameter and the output file name as the second parameter.
For example:
```
./pa1 c5-fcfs.in c5-fcfs.stu
```
where the files’ extensions correspond to `input` and `student`.

## Output Format
The output of each schedule shall be formatted as shown below. Note that the output below corresponds to the input file shown above. You can also review the files that have the base extension for each test cases’ expected outputs.

```
2 processes
Using Round-Robin
Quantum   2
Time   0 : P2 arrived
Time   0 : P2 selected (burst   9)
Time   2 : P2 selected (burst   7)
Time   3 : P1 arrived
Time   4 : P1 selected (burst   5)
Time   6 : P2 selected (burst   5)
Time   8 : P1 selected (burst   3)
Time  10 : P2 selected (burst   3)
Time  12 : P1 selected (burst   1)
Time  13 : P1 finished
Time  13 : P2 selected (burst   1)
Time  14 : P2 finished
Time  14 : Idle
Finished at time  15

P1 wait   5 turnaround  10
P2 wait   5 turnaround  14
```

#### Notes:
1. All integer numeric data is formatted using %3d.
2.  Process names are specified in the input file.
