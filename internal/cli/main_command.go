package cli

import (
	"bufio"
	"errors"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jahkeup/smashing/ini"
	"github.com/jahkeup/smashing/internal/build"
	"github.com/jahkeup/smashing/internal/log"
)

// MainCommand is essentially "main.main".
func MainCommand() *cobra.Command {
	var (
		selectedParser string
	)
	log.Logger.SetLevel(logrus.TraceLevel)

	cmd := &cobra.Command{
		Use:     "smashing",
		Version: build.Version,
	}

	cmd.Flags().StringVarP(&selectedParser, "parser", "p", "ini", "Parser to read input with")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		var proc processor
		switch selectedParser {
		case "ini":
			proc = ini.NewReadWriter(nil)
		default:
			return errors.New("invalid selection")
		}

		ctx := log.WithLogger(
			cmd.Context(),
			log.Logger.WithField("proc", selectedParser))

		if len(args) == 0 {
			log.G(ctx).Trace("reading stdin")
			if err := proc.Read(ctx, bufio.NewReader(os.Stdin)); err != nil {
				log.G(ctx).WithError(err).Error("cannot read input")
				return err
			}
		}

		for _, f := range args {
			l := log.G(ctx).WithField("file", f)
			l.Trace("reading file")

			fd, err := os.OpenFile(f, os.O_RDONLY, 0)
			if err != nil {
				l.WithError(err).Error("cannot open file")
				return err
			}

			ctx := log.WithLogger(ctx, l)

			err = proc.Read(ctx, fd)
			defer fd.Close() // yeah... stack em up.

			if err != nil {
				l.WithError(err).Error("cannot read data")
				return err
			}
		}

		err := proc.Write(ctx, os.Stdout)

		return err
	}

	return cmd
}
