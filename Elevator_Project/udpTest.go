package main

// ip-adress: 129.241.187.23
import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
	"encoding/json"
	"strconv"
)


const PORT int = 40000;//setting the portnumber. Selected a random port for testing purposes
type Message struct {
	Target_Floor int
	Current_Floor_location int 
	Id int
	Timestamp int
}


/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}


func UDPListen(listenPort int) {

	/* For testing: sett addresse lik ip#255:30000*/
	//ServerAddr, err := net.ResolveUDPAddr("udp", ":40000")
	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(listenPort))
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buffer)
		fmt.Println("Received ", string(buffer[0:n]), " from ", addr)



		CheckError(err)
		time.Sleep(time.Second * 1)
		}
	

	}


 
//need to include message-sending
func UDPSend(isMaster bool,transmitPort int,masterIP string) {

	/* Dial up UDP */

	if (isMaster) {
		BroadcastAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:"+strconv.Itoa(transmitPort))
		CheckError(err)
		/* Make a connection to the server */
		Conn, err := net.DialUDP("udp", nil, BroadcastAddr)
		CheckError(err)
		fmt.Println("This is the master")
		defer Conn.Close()
		msg1:=&Message{
		 Target_Floor: 3,
		 Current_Floor_location: 2,
		 Id: 1,
		 Timestamp: 0}
			for {
			msg1.Timestamp = msg1.Timestamp+1
			buf,err := json.Marshal(msg1)
			CheckError(err)
			Conn.Write(buf)
			time.Sleep(time.Second * 5)
		}
	} else {
		BroadcastAddr, err := net.ResolveUDPAddr("udp", masterIP+strconv.Itoa(transmitPort))
		CheckError(err)
		/* Make a connection to the server */
		Conn, err := net.DialUDP("udp", nil, BroadcastAddr)
		CheckError(err)
		fmt.Println("This is a slave")
		defer Conn.Close()
		msg1:=&Message{
		 Target_Floor: 3,
		 Current_Floor_location: 2,
		 Id: 1,
		 Timestamp: 0}
			for {
			msg1.Timestamp = msg1.Timestamp+1
			buf,err := json.Marshal(msg1)
			CheckError(err)
			Conn.Write(buf)
			time.Sleep(time.Second * 5)
		}
	}
	
}

	

	

func UDP_initialize(isMaster bool, port int,masterIP string) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go UDPListen(port)
	go UDPSend(isMaster,port,masterIP)
	time.Sleep(time.Second * 60)
}
func main(){

	UDP_initialize(true,PORT,"129.241.187.150")
}