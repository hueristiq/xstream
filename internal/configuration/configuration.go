package configuration

import "github.com/logrusorgru/aurora/v3"

const (
	NAME    string = "xstream"
	VERSION string = "0.1.0"
)

var BANNER = aurora.Sprintf(
	aurora.BrightBlue(` 
          _
__  _____| |_ _ __ ___  __ _ _ __ ___
\ \/ / __| __| '__/ _ \/ _`+"`"+` | '_ `+"`"+` _ \
 >  <\__ \ |_| | |  __/ (_| | | | | | |
/_/\_\___/\__|_|  \___|\__,_|_| |_| |_|
                                 %s

       %s`).Bold(),
	aurora.BrightRed("v"+VERSION).Bold(),
	aurora.BrightYellow("with <3 by Hueristiq Open Source").Italic(),
)
