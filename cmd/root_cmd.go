package cmd

import (
	"fmt"
	"holiday/router"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/spf13/cobra"
)

const (
	Project = "dictionary"
)

var rootCmd = &cobra.Command{
	Use: Project,
	Run: func(cmd *cobra.Command, args []string) {
		routers := router.Router{}
		g := gin.Default()
		routers.InitRouter(g)
		err := g.Run(":8089")
		if err != nil {
			panic("kill server")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
