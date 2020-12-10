package main
import (
	"fmt"
	"os"
	"os/exec"
	"golang.org/x/sys/unix"
)

func must(err error) {
	if err!= nil {
		panic(err)
	}
}
func run(){
	fmt.Println("run")
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...
)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &unix.SysProcAttr{
		Cloneflags: unix.CLONE_NEWUTS,
	}
	must(cmd.Run())
}
func main (){
	if os.Geteuid() != 0{
		fmt.Println("you need to be root")
		return
	}
	args := os.Args
	fmt.Println("Hello World")
	fmt.Println(args)
	if len(args) >= 2 {
		if args[1] == "run"{
			run()
		} else if args[1] == "child"{
			child()
		} else {
			fmt.Println("unknown arguments")
		}
	}else {
		fmt.Println("too less arguments")
	}
}
func child() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	unix.Sethostname([]byte("container"))
	must(cmd.Run())
}
