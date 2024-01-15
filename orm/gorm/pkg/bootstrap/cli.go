package bootstrap

// FIX: missing go.sum entry for module providing package XXXXXXXXXXXXXXXXXXXXX

import (
	// wire
	_ "github.com/google/subcommands"
	_ "golang.org/x/tools/go/ast/astutil"
	_ "golang.org/x/tools/go/packages"

	// ent
	_ "github.com/olekukonko/tablewriter"
	_ "github.com/spf13/cobra"
)
