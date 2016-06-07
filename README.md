# tumblr-mock-stream


##Usage

This is a go program.  You must have go installed to use it. On OSX you can use Homebrew to install it.
```sh 
brew install go
```

Make sure to follow the instructions from Homebrew and add the proper environment variables
https://golang.org/doc/code.html#GOPATH

Once everything is installed, download the package using `go get`
```sh 
go get "github.com/brettallred/tumblr-mock-stream"
```

Install the package
```sh
cd $GOPATH/src/github.com/brettallred/tumblr-mock-stream/
go install
```

Run the package (assuming your go bin folder is in your PATH)
```sh
tumblr-mock-stream
```

In a different window you can consume the stream

```sh
curl -i http://127.0.0.1:9999/stream
```

You should see something like

```sh
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Tue, 07 Jun 2016 04:06:54 GMT
Transfer-Encoding: chunked

{"id":"5577006791947779410","timestamp":1465272414,"version":"2.2","activity_privacy":"public","activity_type":"blog","activity":{"id":"8674665223082153551","action":"create","blog":{"created":4787198327,"updated":4787198327,"name":"Blog name","url":"http://blog-name.tumblr.com","title":"Title of blog","description":"Blog description","is_group_blog":false,"is_primary":false,"timezone":"US/EST","post_count":744,"last_post":4787198327,"language":"English"}}}
```







