package main


import (
	"fmt",
)

type CLI struct {
	bc *Blockchain
}


func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
			err := addBlockCmd.Parse(os.Args[2:])
	case "printchain":
			err := printChainCmd.Parse(os.Args[2:])
	default:
			cli.printUsage()
			os.Exit(1)
	}

	if addBlockCmd.Parsed() {
			if *addBlockData == "" {
					addBlockCmd.Usage()
					os.Exit(1)
			}
			cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
			cli.printChain()
	}
}

func main() {
	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}