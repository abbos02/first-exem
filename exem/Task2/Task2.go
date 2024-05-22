package letter

import "strings"

func Letters(str string) string{
	res := ""
	words := strings.Split(str," ")
	for _,v := range words {
		for i,v2 := range string(v) {
			if i==0 {
				res+=string(v2-32)
			}else  {
				res+=string(v2)
			}
				
		}
		res+=" "
	}
	return res
}

