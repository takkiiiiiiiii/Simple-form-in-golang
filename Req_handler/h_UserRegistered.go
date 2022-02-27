package Req_handler // 独自のHTTPリクエストハンドラパッケージ

import (
    "database/sql"
    "fmt"
    "html/template"
    "net/http"

    "Templates/conf"    // 実装した設定パッケージの読み込み
    "Templates/query"   // 実装したクエリパッケージの読み込み
    _ "github.com/go-sql-driver/mysql"
)

// 登録結果の確認画面
func HandlerUserRegistered(w http.ResponseWriter, req *http.Request) {

    // POSTデータINSERT関数を実行
   result := insertPostedUser(req)

    // テンプレートをパースする
    tpl := template.Must(template.ParseFiles("/Users/yudai/Go/Project/Templates/Req_handler/user-registered.gtpl"))

    // テンプレートに出力する値をマップにセット
    values := map[string]string{
        "result": result,
    }

    // マップを展開してテンプレートを出力する
    if err := tpl.ExecuteTemplate(w, "user-registered.gtpl", values); err != nil {
        fmt.Println(err)
    }
}
// POSTデータINSERT関数
func insertPostedUser(req *http.Request) string {

    // 正常終了時のreturn値
    result := "ユーザ情報の登録に成功しました。"

    // 設定ファイルを読み込む
    confDB, err := conf.ReadConfDB()
    if err != nil {
        result = "設定ファイルの読み込みに失敗しました。"
    }

    // 設定値から接続文字列を生成
    conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", confDB.User, confDB.Pass, confDB.Host, confDB.Port, confDB.DbName, confDB.Charset)

    fmt.Println(conStr)
    // データベース接続
    db, err := sql.Open("mysql", conStr)
    if err != nil {
        result = "データベースへの接続に失敗しました。"
    }
    // deferで処理終了前に必ず接続をクローズする
    defer db.Close()

    // POST値を渡してINSERT処理を実行
    _, err = query.InsertUser(req.FormValue("account"), req.FormValue("name"), req.FormValue("passwd"), db)
    if err != nil {
        result = "ユーザ情報の登録に失敗しました。"
    }

    // 結果をreturnする
    return result
}
