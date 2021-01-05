package main

import (
	"fmt"
)

//�ۼ���
func Accumulate(value int) func() int {

	//����һ����������������value���ⲿ������Accumulate�����ı�������˷��ص��Ǳհ�
	return func() int {
		value++
		return value
	}
}

func main() {

	//����һ���ۼ��������ص��Ǳհ�
	//�հ�����main�����ģ������������ͱ���(value)���������ں�main��ͬ
	accumulator := Accumulate(0)

	//�Ահ��б����ۼ�
	fmt.Println(accumulator())
	fmt.Println(accumulator())

	//��ӡ�ۼ����ĵ�ַ
	fmt.Printf("%p\n", accumulator)

	//�ִ���һ���ۼ����հ�
	accumulator2 := Accumulate(1)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", accumulator2)

	//ֱ�Ӵ�����ʹ�ñհ�
	fmt.Printf("%p\n", Accumulate(2))
}
