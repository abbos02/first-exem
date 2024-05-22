package summa

func Add(slc []int) int {
	sum:=0
	for _,v := range slc {
		sum+=v
	}
	return sum
}