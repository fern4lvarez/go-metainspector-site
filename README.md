# go-metainspector-site 

## See the [go-metainspector](http://github.com/fern4lvarez/go-metainspector) package in action!

* 1: Get the `go-metainspector-site` package

```
go get github.com/fern4lvarez/go-metainspector-site
```

* 2: Run it

```
$ go-metainspector-site
```

* 3: Try it on localhost:8080

## Deploy it on dotCloud

* 1: Create app on dotCloud: `dcapp APP_NAME create custom https://github.com/cloudControl/buildpack-go.git`

* 2: Push and deploy:

```
$ dcapp APP_NAME/default push

$ dcapp APP_NAME/default deploy
```

* 3: Check it out on APP_NAME.dotcloudapp.com

## TL; DR

* http://metainspector.dotcloudapp.com/

## Contribute!
You all are welcome to take a seat and make a contribution to this repo: reviews, issues, feature suggestions, possible code or functionality enhancements... Everything is appreciated!

## TODO
* Nice error message when url is wrong

## License
go-metainspector-site is MIT licensed, see [here](https://github.com/fern4lvarez/go-metainspector-site/blob/master/LICENSE)
