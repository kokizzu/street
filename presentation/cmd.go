package presentation

import (
	"street/domain"
)

type CLI struct {
	*domain.Domain
}

func (c *CLI) Run(args []string) {
	if len(args) < 1 {
		c.Log.Print(`must start with one of: `, allCommands)
		return
	}
	if len(args) < 2 {
		c.Log.Print(`must provide json payload`)
		return
	}

	c.InitTimedBuffer()
	cmdRun(c.Domain, args[0], []byte(args[1]))
	c.Domain.WaitTimedBufferFinalFlush()

}
