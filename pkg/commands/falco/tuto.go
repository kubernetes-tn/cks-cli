package falco

import (
	"cks-cli/pkg/commands/utils"
	_ "embed"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

//go:embed files/tuto/step0.txt
var step0 string

//go:embed files/tuto/step1.txt
var step1 string

//go:embed files/tuto/step2.txt
var step2 string

//go:embed files/tuto/step3.txt
var step3 string

//go:embed files/tuto/step4.txt
var step4 string

//go:embed files/tuto/step5.txt
var step5 string

func Tuto(ctx *cli.Context) error {
	fmt.Println(step0)
	if utils.Confirm("# Are you done?", os.Stdin, 3) {
		fmt.Println(step1)
		if utils.Confirm("# Are you done?", os.Stdin, 3) {
			fmt.Println(step2)
			if utils.Confirm("# Did you SSH to that node?", os.Stdin, 3) {
				fmt.Println(step3)
				if utils.Confirm("# Are you done?", os.Stdin, 3) {
					fmt.Println(step4)
					if utils.Confirm("# Are you done?", os.Stdin, 3) {
						fmt.Println(step5)
						utils.Confirm("# Clean up done?", os.Stdin, 3)
						fmt.Println("END of Falco Tuto - Congrats!")
					}
				}
			}
		}
	}
	return nil
}
