package telnet

import "fmt"

const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func banner() string {
	b :=
		`
$$$$$$$\                       $$$$$$\                      

$$  __$$\                     $$  __$$\                     

$$ |  $$ | $$$$$$\ $$\    $$\ $$ /  $$ | $$$$$$\   $$$$$$$\ 

$$ |  $$ |$$  __$$\\$$\  $$  |$$ |  $$ |$$  __$$\ $$  _____|

$$ |  $$ |$$$$$$$$ |\$$\$$  / $$ |  $$ |$$ /  $$ |\$$$$$$\  

$$ |  $$ |$$   ____| \$$$  /  $$ |  $$ |$$ |  $$ | \____$$\ 

$$$$$$$  |\$$$$$$$\   \$  /    $$$$$$  |$$$$$$$  |$$$$$$$  |

\_______/  \_______|   \_/     \______/ $$  ____/ \_______/ 

                                        $$ |                

                                        $$ |                

                                        \__| 
`
	return fmt.Sprintf("%s%s%s", colorRed, b, colorReset)
}
