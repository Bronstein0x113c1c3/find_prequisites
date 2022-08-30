package main

import "fmt"

var prequisites map[string][]string = map[string][]string{
	"mth1125": []string{"mth1114"},
	"mth1126": []string{"mth1125"},
	"cs2250":  []string{"mth1112"},
	"cs2255":  []string{"cs2250"},
	"cs3310":  []string{"mth1112"},
	"cs3323":  []string{"cs2255", "mth1125"},
	"cs3329":  []string{"cs3323"},
	"cs3332":  []string{"cs3323"},
	"cs3360":  []string{"cs2255"},
	"cs3365":  []string{"cs3310"},
	"cs3370":  []string{"cs3323"},
	"cs3372":  []string{"cs3323"},
	"cs4420":  []string{"cs3323"},
	"cs4445":  []string{"cs3323"},
	"cs4448":  []string{"cs3323"},
	"mth2210": []string{"mth1112"},
	"mth2215": []string{"mth1112"},
	"mth1114": []string{"mth1112"},
	"mth1112": []string{""},
}

func find_prequisite(x string) {
	fmt.Printf("prequisite for %v: \n", x)
	if prequisites[x][0] == "" {
		fmt.Println("no prequisites")
	} else {

		for i := 0; i < len(prequisites[x]); i++ {
			fmt.Println(prequisites[x][i])
		}
		for i := 0; i < len(prequisites[x]); i++ {
			find_prequisite(prequisites[x][i])
		}
	}

}
func main() {
	find_prequisite("cs4448")
}
