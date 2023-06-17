package controllers

import (
	"strings"
)

/*
	 func main(){
		s := "01-03-2000"
		str:= strings.Split(s,"-")
		fmt.Println(Transform(str))
	}
*/
func DateFormat(s []string) []string {
	i := 0
	for i < len(s) {
		for j, val := range s {
			val = strings.ReplaceAll(val, "*", "")
			str := strings.Split(val, "-")
			s[j] = Transform(str)
		}
		i++
	}
	return s
}
func Transform(str []string) string {
	tabint := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	tabmonth := []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	for i, val := range tabint {
		if val == str[1] {
			str[1] = tabmonth[i]
		}
	}
	return strings.Join(str, " - ")
}
