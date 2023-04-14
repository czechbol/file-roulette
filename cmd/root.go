/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"errors"
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	cmdUse   = "file-roulette"
	cmdShort = "file-roulette - Recursively selecting random files in a given folder"
	cmdLong  = `file-roulette - Recursively selecting random files in a given folder

Give it a folder and it will pick a random file inside. 
Give it folder and index and it will return a file with that index.

I have no idea what am I even doing with my time.`
)

type CLIOptions struct {
	SourcePath   string `json:"source"`
	Index        int    `json:"index"`
	LogVerbosity int    `json:"-"`
}

func (cliOpts *CLIOptions) Complete(args []string) error {
	switch {
	case len(args) != 1:
		return fmt.Errorf("Invalid Arguments: Got %v positional args. Need one argument containing the source path", len(args))
	case args[0] == "":
		return errors.New("Invalid Arguments: source path empty")
	case cliOpts.Index < 0:
		return errors.New("Invalid Arguments: Index must be greater than or equal to 0")
	}

	cliOpts.SourcePath = args[0]
	return nil
}

func (cliOpts *CLIOptions) Run(
	cmd *cobra.Command,
	out io.Writer,
) error {
	log.SetLevel(log.Level(3 + cliOpts.LogVerbosity))

	log.Debug("Getting list of files")
	list, err := GetFileList(cliOpts.SourcePath)
	if err != nil {
		return err
	}
	fmt.Printf("Total number of files in provided folder: %d\n", len(*list))

	if len(*list) < cliOpts.Index {
		return fmt.Errorf("Index must be lower than or equal to %d", len(*list))
	} else if cliOpts.Index == 0 {
		fmt.Println("Choosing a random file:")
		ShuffleFileList(list)
		fmt.Println((*list)[0])
	}

	fmt.Println((*list)[cliOpts.Index-1])

	return nil
}

func (cliOpts *CLIOptions) CobraCommand() *cobra.Command {
	version := fmt.Sprintf(`%s %s (built on %s)`, ProgName, ProgVersion, BuildDate)

	cmd := &cobra.Command{
		Use:           cmdUse,
		Short:         cmdShort,
		Long:          cmdLong,
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version,
	}
	f := cmd.Flags()

	f.IntVarP(
		&cliOpts.Index,
		"index",
		"i",
		0,
		"File index. Zero to disable",
	)
	f.CountVarP(
		&cliOpts.LogVerbosity,
		"verbose",
		"v",
		"Logging verbosity. Specify multiple times for higher verbosity",
	)

	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		if err := cliOpts.Complete(args); err != nil {
			return err
		}
		return cliOpts.Run(cmd, cmd.OutOrStderr())
	}

	return cmd
}
