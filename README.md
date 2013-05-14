# go-metainspector-site 

## See the [go-metainspector](http://github.com/fern4lvarez/go-metainspector-site) package in action!

* 1: Get the `go-metainspector-site` package

```
go get github.com/fern4lvarez/go-metainspector-site
```

* 2: Run it

```
$ go-metainspector-site
```

* 3: Try it on localhost:8080

## Deploy it on cloudControl

* 1: Create app on cloudControl: `cctrlapp APP_NAME create custom https://github.com/kr/heroku-buildpack-go.git`

* 2: Push and deploy:

```
$ cctrlapp APP_NAME/default push

$ cctrlapp APP_NAME/default deploy
```

* 3: Check it out on APP_NAME.cloudcontrolapp.com

## TL; DR

* http://gometainspector.cloudcontrolled.com/

## Contribute!
You all are welcome to take a seat and make a contribution to this repo: reviews, issues, feature suggestions, possible code or functionality enhancements... Everything is appreciated!

## TODO
* Nice error message when url is wrong

## License
go-metainspector-site is MIT licensed, see [here](https://github.com/fern4lvarez/go-metainspector-site/blob/master/LICENSE)