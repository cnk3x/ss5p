package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	socks5 "github.com/armon/go-socks5"
	flag "gopkg.in/alecthomas/kingpin.v2"
)

var version = "master"

var (
	ip   net.IP
	port int
	usr  string
	pwd  string
)

func main() {
	l := log.New(os.Stdout, "ss5p", log.Ltime)
	flag.CommandLine.Help = "socks5代理服务器"
	flag.Version(version).VersionFlag.Short('V').Hidden()
	flag.HelpFlag.Hidden()

	flag.Flag("ip", "如果本机没有网卡没有外网ip，需要设置外网ip").Short('i').IPVar(&ip)
	flag.Flag("port", "监听的端口").Short('l').Default("8080").IntVar(&port)
	flag.Flag("usr", "账号").Short('u').Default("").StringVar(&usr)
	flag.Flag("pwd", "密码").Short('p').Default("").StringVar(&pwd)

	flag.Parse()

	conf := &socks5.Config{Logger: l}

	if usr != "" && pwd != "" {
		conf.Credentials = socks5.StaticCredentials{usr: pwd}
	}

	if len(ip) > 0 {
		conf.BindIP = ip
	} else {
		ip = getV4IP()
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	if port == 0 {
		port = 8080
	}

	go func() {
		if err := server.ListenAndServe("tcp", fmt.Sprint("0.0.0.0:", port)); err != nil {
			l.Fatal(err)
		}
	}()

	time.Sleep(time.Second * 1)

	status := fmt.Sprintf("已启动\n----------\n端口: %d\n", port)
	if usr != "" || pwd != "" {
		status += fmt.Sprintf("账号: %s\n密码: %s\n链接: socks5://%s:%s@%s:%d\n", usr, pwd, usr, pwd, ip, port)
	} else {
		status += fmt.Sprintf("链接: socks5://%s:%d\n", ip, port)
	}
	fmt.Println(status)

	q := make(chan int)
	<-q
}

//本机 ipv4 列表
func getV4IP() net.IP {
	iAddrs, _ := net.InterfaceAddrs()
	for _, address := range iAddrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() { //过滤回环地址
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return net.ParseIP("0.0.0.0")
}
