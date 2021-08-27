package cmd

import (
	"fmt"
	"log"

	"github.com/sadasystems/gcsb/pkg/generator/data"
	"github.com/spf13/cobra"
)

func init() {
	genRandStrFlags := generateRandomString.Flags()
	genRandStrFlags.IntVarP(&rndStrLen, "length", "l", 1, "string length")
	genRandStrFlags.IntVarP(&strAmount, "amount", "a", 1, "amount of strings")

	generateCmd.AddCommand(generateRandomString)
	rootCmd.AddCommand(generateCmd)
}

var (
	rndStrLen int
	strAmount int

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate data",
	}

	generateRandomString = &cobra.Command{
		Use:   "randomstr",
		Short: "Generate random string",
		Run:   doGenerateRandomString,
	}
)

func doGenerateRandomString(cmd *cobra.Command, args []string) {
	if rndStrLen <= 0 {
		log.Fatal("String length is required")
	}
	if strAmount <= 0 {
		strAmount = 1
	}
	sg, err := data.NewStringGenerator(data.StringGeneratorConfig{
		Length: rndStrLen,
	})
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	for i < strAmount {
		fmt.Println(sg.Next())
		i++
	}
}
