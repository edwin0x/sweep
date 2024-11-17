package clean

import (
	"time"

	"github.com/edwin0x/sweep/pkg/types"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewCommand creates a new clean command
func NewCommand(cfg *types.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clean",
		Short: "Clean system files",
		Long: `Clean unnecessary files from your system.
Safely removes caches, temporary files, and old downloads.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runClean(cfg)
		},
	}

	// Add clean-specific flags
	cmd.Flags().DurationVar(&cfg.MaxAge, "max-age", 30*24*time.Hour, "Maximum age of files to clean")
	cmd.Flags().StringSliceVar(&cfg.SkipTypes, "skip", []string{}, "File types to skip")

	return cmd
}

func runClean(cfg *types.Config) error {
	title := color.New(color.FgCyan, color.Bold)
	success := color.New(color.FgGreen)

	if cfg.DryRun {
		color.Yellow("🔍 Dry run mode - no changes will be made")
	}

	title.Println("\n🧹 Starting cleanup process...")

	// Implement cleaning logic here
	// TODO: Add actual cleaning implementation

	success.Println("✨ Cleanup completed successfully!")
	return nil
}
