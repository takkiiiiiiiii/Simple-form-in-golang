package Req_handler // 独自のHTTPリクエストハンドラパッケージ

import (
    "fmt"
    "html/template"
    "net/http"
)

// 入力内容の確認画面
func HandlerUserConfirm(w http.ResponseWriter, req *http.Request) {
    // テンプレートをパースする
    tpl := template.Must(template.ParseFiles("/Users/yudai/GoProjects/Req_handler/user-confirm.gtpl"))

    // テンプレートに出力する値をマップにセット
    values := map[string]string{
        "account": req.FormValue("account"),
        "name":    req.FormValue("name"),
        "passwd":  req.FormValue("passwd"),
        "hid_account": req.FormValue("account"),
        "hid_name":    req.FormValue("name"),
        "hid_passwd":  req.FormValue("passwd"),
    }
   // Form.Value() 
   // http.RequestからGET/POSTされたパラメータを取得する際に、キー名を指定する

    // マップを展開してテンプレートを出力する
    if err := tpl.ExecuteTemplate(w, "user-confirm.gtpl", values); err != nil {
        fmt.Println(err)
    }
}
