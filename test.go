package http_test

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"

)

func HomeWork(res http.ResponseWriter, req *http.Request) {

	//查看全部数据
	fmt.Fprintln(res,"Header全部数据:",req.Header)
	header := req.Header
	//for循环查看获取
	for k,v :=range header{
		fmt.Println("key is",k,"value is",v)
		res.Header().Set(k,v[0])
	}

}

func HomeWorkSec(res http.ResponseWriter,req *http.Request)  {

	//无版本 set一个
	os.Setenv("VERSION", "v1")
	osversion := os.Getenv("VERSION")
	if osversion != "" {
		fmt.Printf("VERSION_ENV is %s.\n", osversion)
		res.Header().Set("VERSION",osversion)
	}else {
		fmt.Printf("VERSION_ENV is not found \n")
	}

}

func HomeWorkThr(res http.ResponseWriter, req *http.Request)  {
	remoteAddr := req.RemoteAddr
	//fmt.Println(remoteAddr)
	if ip := req.Header.Get("Remote_addr"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	//fmt.Println(remoteAddr)

	host,port, err := net.SplitHostPort(req.RemoteAddr)
	fmt.Println(host,port,err)
	if err == nil {
		if net.ParseIP(host) != nil{
			fmt.Printf("host is:%s\nport is:%s\n", host, port)
		}
	}else{
		fmt.Printf("error :%s\n", err.Error())
	}
	fmt.Printf("status is :%v\n", http.StatusOK)

}

func HomeHealthz (res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("test_healthz"))
	res.WriteHeader(200)

}


func TestHttpReque(t *testing.T) {
	http.HandleFunc("/homefir", HomeWork)
	http.HandleFunc("/homesec", HomeWorkSec)
	http.HandleFunc("/homethr", HomeWorkThr)
	http.HandleFunc("/healthz", HomeHealthz)
	http.ListenAndServe(":8080", nil)
}

