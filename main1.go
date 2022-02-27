package main

import (
    "net/http"
    "GoProjects/Req_handler"
)

func main() {

    // "user-form"へのリクエストを関数で処理する
    http.HandleFunc("/user-form", Req_handler.HandlerUserForm)

    // "user-confirm"へのリクエストを関数で処理する
    http.HandleFunc("/user-confirm", Req_handler.HandlerUserConfirm)

    //"user-registered"へのリクエストを関数で処理する
    http.HandleFunc("/user-registered", Req_handler.HandlerUserRegistered)
    // サーバーを起動
    http.ListenAndServe(":9090", nil)
}
