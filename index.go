package handler
 
import (
	"fmt"
	"net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>NO CONTENT</h1>")
	OutputCSV()
}
func (self *Sample) OutputCSV(c web.C, w http.ResponseWriter, r *http.Request) {
    // 保存させたい文字列をbyte配列にする
    out := []byte("test,test2,test3")

    // いろんな文字列情報
    // ファイル名
    w.Header().Set("Content-Disposition", "attachment; filename=test.csv")
    // コンテントタイプ
    w.Header().Set("Content-Type", "text/csv")
    // ファイルの長さ
    w.Header().Set("Content-Length", string(len(out)))
    // bodyに書き込み
    w.Write(out)
}