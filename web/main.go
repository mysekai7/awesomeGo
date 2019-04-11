package main

import (
	"github.com/kataras/iris"
	"mime/multipart"
	"strings"
)

const maxSize = 5 << 20 // 5MB

func main() {
	app := iris.Default()
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})
	app.Get("/users/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		ctx.Writef("User with ID: %d", id)
	})

	// However, this one will match /user/john/ and also /user/john/send.
	app.Get("/user/{name:string}/{action:path}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		ctx.WriteString(message)
	})

	// This handler will match /user/john but will not match neither /user/ or /user.
	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})

	app.Post("/upload", iris.LimitRequestBodySize(maxSize), func(ctx iris.Context) {
		//
		// UploadFormFiles
		// uploads any number of incoming files ("multiple" property on the form input).
		//

		// The second, optional, argument
		// can be used to change a file's name based on the request,
		// at this example we will showcase how to use it
		// by prefixing the uploaded file with the current user's ip.
		ctx.UploadFormFiles("./uploads", beforeSave)
	})

	// Simple group: v1.
	v1 := app.Party("/v1")
	{
		v1.Get("/login", func(ctx iris.Context) {
			ctx.Writef("login")
		})
	}

	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8081"))
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	// make sure you format the ip in a way
	// that can be used for a file name (simple case):
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	// you can use the time.Now, to prefix or suffix the files
	// based on the current time as well, as an exercise.
	// i.e unixTime :=	time.Now().Unix()
	// prefix the Filename with the $IP-
	// no need for more actions, internal uploader will use this
	// name to save the file into the "./uploads" folder.
	file.Filename = ip + "-" + file.Filename
}
