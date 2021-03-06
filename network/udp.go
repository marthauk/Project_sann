package main

// ip-adress: 129.241.187.23
import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
	"encoding/json"
)



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

func UDPListen(Master bool) {

	/* For testing: sett addresse lik ip#255:30000*/
	
	if(Master){
	
		ServerAddr, err := net.ResolveUDPAddr("udp", ":40000")
		CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buffer)
		fmt.Println("Received ", string(buffer[0:n]), " from ", addr)



		if err != nil {
			fmt.Println("Error: ", err)
		}
		time.Sleep(time.Second * 1)
		}
	}

	}
	



}

func UDPSend() {

	/* Dial up UDP */
	BroadcastAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:40000")
	CheckError(err)
	/* Make a connection to the server */
	Conn, err := net.DialUDP("udp", nil, BroadcastAddr)
	CheckError(err)

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

	

	

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go UDPListen()
	go UDPSend()
	time.Sleep(time.Second * 30)
}