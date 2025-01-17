package client

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/x/data"
)

// QueryCmd returns the parent command for all x/data CLI query commands
func QueryCmd(name string) *cobra.Command {
	queryByIRICmd := QueryByIRICmd()

	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Use:   fmt.Sprintf("%s [iri]", name),
		Short: "Querying commands for the data module",
		Long: strings.TrimSpace(`Querying commands for the data module.

If an IRI is passed as first argument, then this command will query timestamp,
attestors and content (if available) for the given IRI. Otherwise, this command
will run the given subcommand.

Example (the two following commands are equivalent):
$ regen query data regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf
$ regen query data by-iri regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf`),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE: func(cmd *cobra.Command, args []string) error {
			// If 1st arg is NOT an IRI, parse subcommands as usual.
			if !strings.Contains(args[0], "regen:") {
				return client.ValidateCmd(cmd, args)
			}

			// Or else, we call QueryByIRICmd.
			return queryByIRICmd.RunE(cmd, args)
		},
	}

	cmd.AddCommand(
		queryByIRICmd,
		QueryByAttestorCmd(),
		QueryHashByIRICmd(),
		QueryAttestorsCmd(),
		QueryResolverInfoCmd(),
		QueryResolversCmd(),
	)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryByIRICmd creates a CLI command for Query/Data.
func QueryByIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "by-iri [iri]",
		Short: "Query for anchored data based on IRI",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.ByIRI(cmd.Context(), &data.QueryByIRIRequest{
				Iri: args[0],
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryByAttestorCmd creates a CLI command for Query/ByAttestor.
func QueryByAttestorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "by-attestor [attestor]",
		Short: "Query for anchored data based on an attestor",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.ByAttestor(cmd.Context(), &data.QueryByAttestorRequest{
				Attestor:   args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryHashByIRICmd creates a CLI command for Query/HashByIRI.
func QueryHashByIRICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content-hash [iri]",
		Short: "Query for content hash based on IRI",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.HashByIRI(cmd.Context(), &data.QueryHashByIRIRequest{
				Iri: args[0],
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryAttestorsCmd creates a CLI command for Query/Attestors.
func QueryAttestorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attestors [iri]",
		Short: "Query for attestors based on IRI",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.AttestorsByIRI(cmd.Context(), &data.QueryAttestorsByIRIRequest{
				Iri:        args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryResolverInfoCmd creates a CLI command for Query/ResolverInfo.
func QueryResolverInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolver-info [url]",
		Short: "Query for resolver information",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.ResolverInfo(cmd.Context(), &data.QueryResolverInfoRequest{
				Url: args[0],
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryResolversCmd creates a CLI command for Query/Resolvers.
func QueryResolversCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers [iri]",
		Short: "Query for registered resolver URLs for the data IRI",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.ResolversByIRI(cmd.Context(), &data.QueryResolversByIRIRequest{
				Iri:        args[0],
				Pagination: pagination,
			})

			return print(ctx, res, err)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "resolvers")

	return cmd
}
