package symbol

import (
	"fmt"
	"github.com/iovisor/gobpf/bcc"
	"os"
	"strconv"
)

func main() {
	pid, _ := strconv.Atoi(os.Args[1])
	addr, _ := strconv.ParseInt(os.Args[2], 10, 10)

	byAddr, err := bcc.BccResolveNameByAddr("", pid, addr)
	fmt.Printf("symbol: %s, err: %v", byAddr, err)
}
