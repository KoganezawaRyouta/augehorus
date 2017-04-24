package cmd

import (
	"github.com/KoganezawaRyouta/augehorus/batches"
	"github.com/spf13/cobra"
)

var importerCmd = &cobra.Command{
	Use:   "importer",
	Short: "import bitcoin rate info to database",
	Long:  `import bitcoin rate info to database`,
	Run: func(cmd *cobra.Command, args []string) {
		config := LoadConfig(configName)
		batches.NewVacuummer(config).Run()
	},
}
