//sinimage.go

package main

import(
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"math"
)

func main(){
	//图片大小
	const size = 300;

	//根据给定大小创建灰度图
	//image.Rect描述一个方形的两个定位点（xl,yl ）和（x2，y2） 
	pic := image.NewGray(image.Rect(0, 0, size, size))

	//遍历每个像素, 初始化成白色
	for x := 0; x < size; x++ {
		for y :=0; y < size; y++ {
			//填充为臼色
			//灰度图颜色由8位组成，灰度范围为0～255，0表示黑色，255表示白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	//从 0 到最大像素生成 x 坐标
	for x := 0; x < size; x++ {
		//让像素显示的sin横坐标的值，映射在范围0～2Pi之间
		s := float64(x) * 2 * math.Pi / size

		//sin的纵坐标映射, 注意纵坐标范围是 0 到 size
		y := size/2 - math.Sin(s) * size / 2

		//绘制 sin 轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}

	//写入文件
	//创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	//使用 PNG 格式将数据写入文件
	png.Encode(file, pic)
	//关闭文件
	file.Close()
}