package mr



import (
	"os"
	"strconv"
)

type Job struct {
	Type              int
	WorkerNumber      int
	InputFile         string
	IntermediateFiles []string
	ShouldQuit        bool
}

type MapJob struct {
	InputFile    string
	MapJobNumber int
	ReducerCount int
}

type ReduceJob struct {
	IntermediateFiles []string
	ReduceNumber      int
}

type RequestTaskArgs struct {
	Pid int
}

type RequestTaskReply struct {
	MapJob    *MapJob
	ReduceJob *ReduceJob
	Done      bool
}

type ReportMapTaskArgs struct {
	InputFile        string
	IntermediateFile []string
	Pid              int
}

type ReportMapTaskReply struct {
}

type ReportReduceTaskArgs struct {
	Pid          int
	ReduceNumber int
}

type ReportReduceTaskReply struct {
}


func coordinatorSock() string {
	s := "/var/tmp/cs612-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
