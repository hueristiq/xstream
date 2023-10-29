package configuration

import "github.com/logrusorgru/aurora/v3"

const (
	NAME    string = "xstreamin"
	VERSION string = "0.0.0"
)

var BANNER = aurora.Sprintf(
	aurora.BrightBlue(`
          _                            _
__  _____| |_ _ __ ___  __ _ _ __ ___ (_)_ __
\ \/ / __| __| '__/ _ \/ _`+"`"+` | '_ `+"`"+` _ \| | '_ \
 >  <\__ \ |_| | |  __/ (_| | | | | | | | | | |
/_/\_\___/\__|_|  \___|\__,_|_| |_| |_|_|_| |_|
                                         %s

               %s`).Bold(),
	aurora.BrightRed("v"+VERSION).Bold(),
	aurora.BrightYellow("with <3 by Hueristiq Open Source").Italic(),
)
