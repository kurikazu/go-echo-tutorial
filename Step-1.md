# Step-1: 動かしてみよう!

## ディレクトリの設定

アプリケーションを作るために必要なディレクトリ構成を作ります。

```
$ cd $GOPATH/src
$ mkdir tutorial
```

## はじめてのgoファイル

`$GOPATH/src/tutorial`の下に`main.go`を作成し、下記のように書いてください。

```$GOPATH/src/tutorial/main.go
// go run main.go
package main

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)
```

先ほどのimport句の下に`init`関数を書いてください。

```$GOPATH/src/tutorial/main.go
...
import (
...
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
```

`init`の下に`main`関数を書きます。

```$GOPATH/src/tutorial/main.go
...
import (
...
)

func init() {
...
}

func main() {
	// instance
	e := echo.New()
	e.Debug()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.Get("/", hello)

	// start server
	e.Run(fasthttp.New(":" + port()))
}
```

`Hello World!`という文字列を返却する`hello`関数の中身を書きます。

```$GOPATH/src/tutorial/main.go
...
func main() {
...
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!\n")
}
```

アクセスするポートを指定する`port`関数の中身を書きます。  

```$GOPATH/src/tutorial/main.go
...
func main() {
...
}

func Hello(c echo.Context) error {
...
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080" // localhost:8080
	}
	return port
}
```

## Hello World!
ターミナルで、下記のように実行します。

```
go run main.go
```

ブラウザで`http://localhost:8080/`を見てみましょう。  
ターミナル側にもログが記録されています。