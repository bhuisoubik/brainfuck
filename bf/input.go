package bf

import (
	"os"
	"bufio"
)

func input() byte {
	inp, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return []byte(inp)[0]
}
