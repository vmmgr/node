package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmmgr/node/pkg/core/request"
	"github.com/vmmgr/node/pkg/core/storage"
	"github.com/vmmgr/node/pkg/core/tool"
	"log"
)

// copyCmd represents the start command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		src, err := cmd.Flags().GetString("src")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		dest, err := cmd.Flags().GetString("dest")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		controller, err := cmd.Flags().GetString("controller")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		uuid, err := cmd.Flags().GetString("uuid")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		isDebug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		tool.ChangeDebugMode(isDebug)
		log.Println(src, dest, controller, isDebug)

		req := storage.Base{
			Controller: controller,
			UUID:       uuid,
			IsDebug:    isDebug,
		}
		err = req.FileCopy(src, dest)
		if err != nil {
			request.SendServer(controller, uuid, 0, "", err)
		} else {
			request.SendServer(controller, uuid, 100, "success", nil)
		}

		log.Println("end")
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.PersistentFlags().StringP("src", "s", "", "src filepath")
	copyCmd.PersistentFlags().StringP("dest", "d", "", "dest filepath")
	copyCmd.PersistentFlags().StringP("controller", "c", "", "controller")
	copyCmd.PersistentFlags().StringP("uuid", "u", "", "uuid")
	copyCmd.PersistentFlags().BoolP("debug", "b", false, "debug")
}
