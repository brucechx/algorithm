package josephus

// 约瑟夫环
/*
从零开始编号:
	f[1]=0;
	f[i]=(f[i-1]+k)%i = (f[i-1] +m%i) % i = (f[i-1] + m) % i ;  (i>1)
*/

// 常规解法
func josephus1(n int, k int) int{
	nums := 0 //当num到达N-1时，只剩一个
	flags := make([]int, n)
	i := 0
	for nums != (n - 1){
		find := 0 // find=1找到下一个让它出去,count计数是否数到
		count := 0
		for find == 0{
			tmp := i  % n
			if flags[tmp] == 0 {
				count++
			}
			if count == 3{
				nums ++
				find = 1
				flags[tmp] = 1
			}
			i++
		}
	}
	//fmt.Println(flags)
	for i, val := range flags{
		if val == 0 {
			return i
		}
	}
	return -1
}

// 递归解法
func josephus2(n int,k int) int{
	if n==1 {
		return 0
	}else {
		return (josephus2(n-1,k)+k)%n
	}
}

func josephus3(n, m int) int {
	p := 0
	for i:=2; i<=n; i++{
		p=(p+m)%i
	}
	return p
}
