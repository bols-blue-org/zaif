package main

import (
    "fmt"
    "time"
)

func main() {
    q := make(chan string, 5)

    go func() {
        time.Sleep(1 * time.Second)
        q <- "foo 1"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        q <- "foo 2"
    }()
    go func() {
        time.Sleep(3 * time.Second)
        q <- "bar 3"
    }()
    go func() {
        time.Sleep(3 * time.Second)
        q <- "foo 3"
    }()

    // ほぼ同時に2つトリガが発生するのが分かっている
    var cmds []string
    cmds = append(cmds, <-q) // まずは一つ受信する
wait_some:
    for {
        select {
        case cmd := <-q:
            // 1秒以内なら一緒に処理しちゃうよ
            cmds = append(cmds, cmd)
        case <-time.After(1 * time.Second):
            // 1秒過ぎたらもう受け付けないよ
            break wait_some
        }
    }
    for _, cmd := range cmds {
        fmt.Println(cmd)
    }
}
