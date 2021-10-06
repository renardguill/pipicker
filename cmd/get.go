// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func get(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Vous devez spécifier une url de dépo gitlab !")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("Vous avez spécifié trop d'arguments, une seule url de dépo gitlab est accéptée !")
		os.Exit(1)
	}
	url := args[0]
	//validate url
	//todo
	//

	token := viper.GetString("token")
	if len(token) == 0 {
		fmt.Println("Vous n'avez pas spécifié de token ! (utilisez --token ou le fichier de conf)")
		os.Exit(1)
	}
	fmt.Println("get called", url)
	fmt.Println("get called", token)

	git, err := gitlab.NewClient(token, gitlab.WithBaseURL(fmt.Sprintf("%s/api/v4", url)))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	gf := &gitlab.GetFileOptions{
		Ref: gitlab.String("main"),
	}
	f, _, err := git.RepositoryFiles.GetFile("30235323", "README.md", gf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File contains: %s", f.Content)
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get dependencies",
	Long:  `Cette commande récupère tous les fichiers dont cette pipeline dépend.`,
	Run:   get,
}

func init() {
	RootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getCmd.PersistentFlags().StringP("token", "t", "", "Token d'accès à GitLab")
	viper.BindPFlag("token", getCmd.PersistentFlags().Lookup("token"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
