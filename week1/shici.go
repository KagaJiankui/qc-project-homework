package main

import (
	// "flag"
	"os"
	"strconv"
)

func shici() {
	args := os.Args[1:]
	var rv = helpShici(args)
	switch rv {
	case 1:
		println("梅岭绿阴青子，蒲间清泉白石，怪我旧盟寒。")
	case 2:
		println("织女机丝虚夜月，石鲸鳞甲动秋风")
	case 3:
		println("Error occured: Invalid Arguments")
	default:
	}
}
func helpShici(Args []string) int {
	/*
		处理命令行输入. 状态表如下:
		len(args)=0 or >2 => 3
		args[0] | args[1]
		'h'=>0  | 1/2
		1 =>1   | 3
		2 =>2   | 3
		3
		0:help
		1/2:valid return
		3:error
	*/
	if len(Args) == 1 {
		if Args[0] == "h" {
			println("Help:\n Type 1 to print a line of Song Lyrics\n Type 2 to print a line of Tang Poetry\n Type 0 to print this help.")
			return 0
		} else {
			iv, _ := strconv.Atoi(Args[0])
			if (iv == 1) || (iv == 2) {
				return iv
			} else {
				return 3
			}
		}
	} else if len(Args) == 2 {
		if Args[0] == "h" {
			iv, _ := strconv.Atoi(Args[1])
			switch iv {
			case 1:
				println("Help 1:\n Print a line of Song lyrics")
				return 0
			case 2:
				println("Help 2:\n Print a line of Tang poetry")
				return 0
			default:
				return 3
			}
		} else {
			return 3
		}
	} else {
		return 3
	}
}
