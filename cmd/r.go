/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rCmd represents the r command
var rCmd = &cobra.Command{
	Use:   "r",
	Short: "Rondom output",
	Long:  `Random output of Big Bang Theory titles`,
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

		// csvリーダーを作成
		reader := csv.NewReader(filePath)

		// 行数を取得
		rows, err := reader.ReadAll()
		if err != nil {
			fmt.Printf("Error reading CSV rows: %s\n", err)
		}

		// ランダムなインデックスを生成
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		randomIndex := random.Intn(len(rows))

		// ランダムな行を表示
		randomRow := rows[randomIndex]
		fmt.Printf("season%s episodes:%s title:%s\n", randomRow[1], randomRow[2], randomRow[3])

		return nil
	},
}

func init() {
	rootCmd.AddCommand(rCmd)
}
