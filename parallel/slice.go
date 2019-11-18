package main

func main(){
	onesSlice := make([]int, 1000000000)
	for i:= 0; i < len(onesSlice); i++ {
		onesSlice[i] = i
	}
	var sem = make(chan int, 10)
	for i:= 0; i < len(onesSlice); i++ {
		sem <- 1
		temp:=i
		func(){
			onesSlice[temp] = onesSlice[temp] * 2
			<-sem
		}()
	}
}