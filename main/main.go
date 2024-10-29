package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go-webssh/proto"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var (
	grpcIpAndPort = "172.22.0.5:6565"
)

/*
*
监听ws请求
*/
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Sec-WebSocket-Protocol")
	if token != "" {
		if authTk(token) == 1 {
			id := r.FormValue("id")
			if id != "" {
				header := w.Header()
				header.Set("Sec-WebSocket-Protocol", token)
				//升级请求
				ws, err := upgrader.Upgrade(w, r, header)
				if err != nil {
					log.Println(err)
				}
				sshInfo := getConInfo(id)
				if sshInfo != nil {
					go reader(ws, sshInfo)
				}

			}
		}
	}
}

// 查询信息
func getConInfo(id string) *proto.SshInfo {
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(grpcIpAndPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewConServiceClient(conn)
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//获取连接信息
	r, err := c.GetSshInfo(ctx, &proto.SendId{
		Id: id})
	if err != nil {
		log.Println(err)
	}
	return &proto.SshInfo{
		Ip:       r.Ip,
		Password: r.Password,
	}
}

// token校验
func authTk(token string) int32 {
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(grpcIpAndPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewVerifyAuthClient(conn)
	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AuthToken(ctx, &proto.Request{JwtText: token})
	if err != nil {
		log.Println(err)
	}
	return r.GetResult()
}

// 读取客户端消息
func reader(conn *websocket.Conn, sshinfo *proto.SshInfo) {
	session := ConSsh(&SshInfo{
		address:  sshinfo.Ip,
		port:     22,
		password: sshinfo.Password})
	//session := ConSsh(&SshInfo{
	//	address:  "192.168.25.135",
	//	port:     22,
	//	password: "qwe@123"})
	//获取输入管道
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	//获取输出管道
	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
	}
	if err = session.RequestPty("shell", 50, 400, modes); err != nil {
		log.Fatalf("request pty error: %s", err.Error())
	}
	if err = session.Shell(); err != nil {
		log.Fatalf("start shell error: %s", err.Error())
	}
	//执行命令
	go func() {
		for {
			// read in a message
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			// 去除字符串首尾空白字符
			//执行命令
			fmt.Println("发送：", p)
			_, err = stdin.Write(p)
			if err != nil {
				log.Fatal(err)
				return
			} else if err == io.EOF {
				err := stdin.Close()
				if err != nil {
					return
				}
				return
			}
		}
	}()
	//返回输出
	go func() {
		for {
			buf := make([]byte, 2048)
			n, err := stdout.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("接收到：", buf[:n])
			err = conn.WriteMessage(1, buf[:n])
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}()
}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("接收ws请求")
	//添加接口
	setupRoutes()
	//监听8080端口
	log.Fatal(http.ListenAndServe(":8080", nil))

}
