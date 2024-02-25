/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// フラグバインド変数
var season string

// sCmd represents the s command
var sCmd = &cobra.Command{
	Use:   "s",
	Short: "Output a list of selected seasons",
	Long:  `Output a list of selected seasons of The Big Bang Theory`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// pathの取得
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cannot get user home dir path:%s", err.Error())
		}

		// ファイルが取得ができない場合はヘッダーを出力させる
		file := fmt.Sprintf("%s/.bigbang/title.csv", home)
		if s, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Println("index,season,episodes,title")
			return nil
		} else if s.IsDir() {
			return fmt.Errorf("%s is directory", file)
		}

		// csvファイルを取得するためのポインタを取得
		filePath, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("cannot open file:%s", err.Error())
		}

		// 関数の終了時にファイルを閉じる
		defer filePath.Close()

		// // csvリーダーを作成
		reader := csv.NewReader(filePath)

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("cannot read file :%s", err.Error())
			}

			if len(record) >= 3 && record[1] == season {
				fmt.Printf("season%s episodes:%s title:%s\n", record[1], record[2], record[3])
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sCmd)

	sCmd.Flags().StringVarP(&season, "season", "s", "", "select season")

	//必須のフラグに指定
	sCmd.MarkFlagRequired("season")
}
