//httpserver.go

//标记当前文件为 main 包， main 包也是 Go 程序的入口包。
package main

//导入 net/http 包，这个包的作用是 HTTP 的基础封装和访问。
import(
	"net/http"
)

func main(){
	//使用 http.FileServer 文件服务器将 当前目录作为根目录"/" 的处理器，
	//访问根目录，就会进入当前目录。
	http.Handle("/", http.FileServer(http.Dir(".")))
	//默认的 HTTP 服务侦昕在本机 8080 端口 
	http.ListenAndServe(":9088", nil)
}