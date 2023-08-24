package configuration

import "github.com/logrusorgru/aurora/v3"

const (
	NAME    string = "xt"
	VERSION string = "0.0.0"
)

var (
	BANNER = aurora.Sprintf(
		aurora.BrightBlue(`
       _   
__  _ | |_ 
\ \/ /  __|
 >  < | |_ 
/_/\_\ \__| %s
`).Bold(),
		aurora.BrightYellow("v"+VERSION).Bold(),
	)
)
