package configuration

import "github.com/logrusorgru/aurora/v3"

const (
	NAME    string = "xtee"
	VERSION string = "0.0.0"
)

var (
	BANNER = aurora.Sprintf(
		aurora.BrightBlue(`
       _            
__  _ | |_ ___  ___ 
\ \/ /  __/ _ \/ _ \
 >  < | |_  __/  __/
/_/\_\ \__\___|\___| %s
`).Bold(),
		aurora.BrightYellow("v"+VERSION).Bold(),
	)
)
