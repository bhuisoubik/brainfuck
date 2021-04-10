package bf

import (
	"bufio"
	"os"
)

func input() byte {
	inp, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return []byte(inp)[0]
}
