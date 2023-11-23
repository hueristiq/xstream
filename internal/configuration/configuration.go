package configuration

import "github.com/logrusorgru/aurora/v3"

const (
	NAME    string = "xbridge"
	VERSION string = "0.0.0"
)

var BANNER = aurora.Sprintf(
	aurora.BrightBlue(`
      _          _     _
__  _| |__  _ __(_) __| | __ _  ___
\ \/ / '_ \| '__| |/ _`+"`"+` |/ _`+"`"+` |/ _ \
 >  <| |_) | |  | | (_| | (_| |  __/
/_/\_\_.__/|_|  |_|\__,_|\__, |\___|
                         |___/ 
                              %s

    %s`).Bold(),
	aurora.BrightRed("v"+VERSION).Bold(),
	aurora.BrightYellow("with <3 by Hueristiq Open Source").Italic(),
)
