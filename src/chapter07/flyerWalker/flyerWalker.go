package main

import "fmt"

// 飞行物接口
type Flyer interface {
	Fly()
}

// 行走物接口
type Walker interface {
	Walk()
}

// 能飞能走的组合接口
type FlyerWalker interface {
	Flyer
	Walker
}

// 鸟类
type bird struct {
}

// 实现飞行接口
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// 实现行走接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

// 猪
type pig struct {
}

// 实现行走接口
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

func main() {

	// 动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	// 遍历映射
	for name, obj := range animals {

		fmt.Printf("name: %s\n", name)

		// 断言对象是否为飞行物
		f, isFlyer := obj.(Flyer)
		// 飞行物调用飞行接口
		if isFlyer {
			f.Fly()
		}

		// 断言对象是否为行走物
		w, isWalker := obj.(Walker)
		// 行走物调用行走接口
		if isWalker {
			w.Walk()
		}

		// 断言是否实现组合接口
		fw, isFlyerWalker := obj.(FlyerWalker)
		// 调用组合接口
		if isFlyerWalker {
			fw.Fly()
			fw.Walk()
		}
		
		fmt.Printf("%s isFlyer: %v isWalker: %v isFlyerWalker: %v\n", name, isFlyer, isWalker, isFlyerWalker)

		/* 
		//尝试用type-switch写，不能实现，因为type-switch只能走一个分支
		//cannot fallthrough in type switch
		switch obj.(type) {
			case Flyer:
				fmt.Println("is Flyer")
				f := obj.(Flyer)
				f.Fly()
				fallthrough
			case Walker:
				fmt.Println("is Walker")
				w := obj.(Walker)
				w.Walk()
				fallthrough
			case FlyerWalker:
				fmt.Println("is FlyerWalker")
				fw := obj.(FlyerWalker)
				fw.Fly()
				fw.Walk()
		}
		*/
	}

}
