package main

import (
	"fmt"
)

const LINES int = 10//设定杨辉三角10行，同时也相当于10列

func ShowYangHui(){
	var yh [LINES][LINES]int
	for i := 0; i < LINES; i++ {
		for j := 0; j < i + 1; j++ {
			if i < 2 {//两行以内三角中的数字都是1
				yh[i][j] = 1
			}else{//第三行开始，正式计算数值写入数组
				if j == 0 || j == i {
					yh[i][j] = 1//所有行的第一列和最后一列都是1
				}else{
					yh[i][j] = yh[i-1][j-1] + yh[i-1][j]//当前数组元素是上一行的前一个元素加上上一行的当前列元素
				}
			}
			fmt.Printf("%d\t", yh[i][j] )//格式化输出一行
		}
		fmt.Print("\n")//换行
	}
}

func main(){
	ShowYangHui()
}
