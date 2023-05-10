package cli

type Command struct {
	usage string
	desc  string
}

func NewCommand(usage, desc string) *Command {
	return &Command{
		usage: usage,
		desc:  desc,
	}
}
