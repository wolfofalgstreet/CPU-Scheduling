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
