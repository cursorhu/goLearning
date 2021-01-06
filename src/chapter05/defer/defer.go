package main

import (
	"fmt"
	"sync"
	"os"
)

var (
	//创建map
	m = make(map[string]int)
	//为了并发安全的互斥锁
	mGuard sync.Mutex
)

/* 读取map */

//常规写法，lock-unlock包围临界区
func readMap(k string) int {
	//加锁
	mGuard.Lock()
	//取值
	v := m[k]
	//解锁
	mGuard.Unlock()

	return v
}

//defer写法，申请资源和释放资源一起写，释放资源延迟到函数返回后执行
func readMapUseDefer(k string) int {
	//加锁
	mGuard.Lock()
	//声明延迟执行解锁
	defer mGuard.Unlock()

	return m[k]
}

/* 读取文件 */
//常规写法，多处都要考虑close文件
func getFileSize(file string) int64 {
	//根据文件名打开文件，返回文件句柄和错误码
	f, err := os.Open(file)
	//打开出错，即未获取文件资源
	if err != nil {
		return 0
	}

	//打开成功，读取文件状态
	info, err := f.Stat()
	//状态出错，需要关闭文件再返回
	if err != nil {
		f.Close()
		return 0
	} 

	//获取信息成功，读出文件大小
	size := info.Size()

	//关闭文件，释放资源
	f.Close()

	//正常返回
	return size
}

//defer写法，单处close文件
func getFileSizeUseDefer(file string) int64 {

	f, err := os.Open(file)
	//打开出错，没有获取文件资源，因此不需要close
	if err != nil {
		return 0
	}

	//打开成功，已经获取文件资源，此处定义defer-Close语句
	//该语句现在不执行，此后任何地方函数返回，都会调用此语句
	defer f.Close()
	
	info, err := f.Stat()
	//状态出错，需要关闭文件再返回
	if err != nil {
		//函数返回，并执行defer语句，释放文件资源
		return 0
	} 

	//获取信息成功，读出文件大小
	size := info.Size()

	//正常返回，并执行defer语句，释放文件资源
	return size
}

func main() {


}