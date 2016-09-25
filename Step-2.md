# Step-2: 寝る前に一行妄想を登録するツールを作ってみよう！

## データベースの設定

```fantasies.sql
$ mysql -uroot

create database fantasy;
use fantasy;

create table `fantasies` (
    `id` int(11) not null auto_increment,
    `fantasy` text not null,
    `created_at` datetime not null default now(),
    primary key (`id`)
) engine=InnoDB default charset=utf8;

```

## ディレクトリの追加
アプリケーションを作るために必要なディレクトリ構成を作ります。

```
$ cd $GOPATH/src
$ mkdir step-2
```

アプリケーションの下に、下記ディレクトリを作成します。

- 設定用ディレクトリ
- コントローラーディレクトリ
- モデルディレクトリ
- テンプレートディレクトリ

```
$ cd $GOPATH/src/step-2
$ mkdir config
$ mkdir controllers
$ mkdir models
$ mkdir views
```

## 妄想を追加してみよう！
`$GOPATH/src/step-2`の下に`main.go`を作成し、下記のように書いてください。

```$GOPATH/src/step-2/main.go
// go run main.go
package main

import (
	"os"
	"io"
	"html/template"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"

	"step-2/controllers"
)

// template
type Template struct {
	templates *template.Template
}

// template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	return t.templates.ExecuteTemplate(w, name, data)
}

func init() {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	// instance
	e := echo.New()
	e.Debug()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// template
	t := &Template {
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.SetRenderer(t)

	// routes
	e.Get("/fantasy",  controllers.InputFantasy())

	// start server
	e.Run(fasthttp.New(":" + port()))
}

func port() string {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080" // localhost:8080
	}
	return port
}
```

ここでは、コントローラの追加と、テンプレートを読み込む部分を書いてあります。
`step-1`で定義した`hello()`にあたる`InputFantasy()`をコントローラディレクトリに移動しています。

`$GOPATH/src/step-2/controllers`の下に`fantasy.go`を作成し、下記のように書いてください。

```$GOPATH/src/step-2/controllers/fantasy.go
package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

func InputFantasy() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		return c.Render(http.StatusOK, "input", nil)
	}
}
```

`$GOPATH/src/step-2/views`の下に`input.html`を作成し、下記のように書いてください。

```$GOPATH/src/step-2/views/input.html
{{define "input"}}
<!DOCTYPE html>
<html lang="ja">
  <head>
    <title>一行妄想</title>
  </head>
<body>
  <h1>一行妄想</h1>
  <h2>登録</h2>
  <form name="form" method="POST" action="/fantasy/post">
    <input type="text" maxlength="100" name="fantasy">
    <input type="submit" value="登録">
  </form>
</body>
</html>
{{end}}
```

`go run main.go`を実行し、ブラウザで`http://localhost:8080/fantasy`を見てみましょう。  

![Demo](https://{T.B.D})

`$GOPATH/src/step-2/views`の下に`conplete.html`を作成し、下記のように書いてください。
これは、完了画面のテンプレートです。

```$GOPATH/src/step-2/views/input.html
{{define "complete"}}
<!DOCTYPE html>
<html lang="ja">
<head>
  <title>一行妄想</title>
</head>
<body>
  <h1>登録完了!</h1>
  <a href="/">まだまだ登録するよ!</a>
</body>
</html>
{{end}}
```

`$GOPATH/src/step-2/config`の下に`mysql.go`を作成し、下記のように書いてください。
設定値はインストールしたmysqlの設定次第となります。

```$GOPATH/src/step-2/config/mysql.go
package config

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

const (
	User     string = "root"
	Password string = ""
	Db       string = "fantasy"
	Host     string = "localhost"
	Port     string = "3306"
)

func GetSession() *dbr.Session {

	db, err := dbr.Open("mysql", User+":"+Password+"@tcp("+Host+":"+Port+")/"+Db+"?parseTime=true&loc=Asia%2FTokyo", nil)

	if err != nil {
		logrus.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}
```

`$GOPATH/src/step-2/models`の下に`fantasy.go`を作成し、下記のように書いてください。

```$GOPATH/src/step-2/models/fantasy.go
package models

import (
	"github.com/gocraft/dbr"
)

type (
	Fantasy struct {
		ID        int64     `db:"id"`
		Fantasy   string    `db:"name"`
		Createdat time.Time `db:"createdAt"`
	}
)

func New(fantasy string) *Fantasy {

	return &Fantasy {
		Fantasy:   fantasy,
	}
}

func (f *Fantasy) Post(sess *dbr.Session) error {

	_, err := sess.InsertInto("fantasies").Columns("fantasy").Record(f).Exec()
	return err
}
```

`fantasy.go`コントローラにimportとメソッドを追加します。

```$GOPATH/src/step-2/controllers/fantasy.go
package controllers

import (
...
	"step-2/models"
	"step-2/config"
)

func InputFantasy() echo.HandlerFunc {
...
}

func PostFantasy() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		fantasy := models.New(c.FormValue("fantasy"))
		session := config.GetSession()

		if err := fantasy.Post(session); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}

		return c.Render(http.StatusOK, "complete", nil)
	}
}
```

`main.go`コントローラにrouterを追加します。

```$GOPATH/src/step-2/main.go
...
func main() {
...
	// routes
	e.Get("/fantasy",  controllers.InputFantasy())
	e.Post("/fantasy/post", controllers.PostFantasy())
...
}
...
```

## 追加した妄想を見てみよう！


