// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"crypto/tls"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.SetConfigName("s3clientconfig")
		viper.SetConfigType("json")
		viper.AddConfigPath("/etc")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Not able to load config: %s \n", err))
		}
		fmt.Println(viper.GetString("endpoint.address"))
		fmt.Println(viper.GetString("credentials.key"))
		fmt.Println(viper.GetString("credentials.secret"))

		endpoint := viper.GetString("endpoint.address")
		key := viper.GetString("credentials.key")
		secret := viper.GetString("credentials.secret")

		webscaleConfig := aws.NewConfig().WithRegion("a").WithCredentials(credentials.NewStaticCredentials(key, secret, "")).WithEndpoint(endpoint).WithHTTPClient(&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}).WithS3ForcePathStyle(true)
		s3client := s3.New(session.New(webscaleConfig))
		var params *s3.ListBucketsInput
		resp, err := s3client.ListBuckets(params)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
