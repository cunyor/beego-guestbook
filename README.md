# Beego Guestbook

This is sample guestbook app with beego.
Forked from: https://github.com/yuhei0718/beego-guestbook

## upload to heroku

```
$ heroku create -b https://github.com/kr/heroku-buildpack-go.git
$ heroku addons:create cleardb:ignite
$ git push heroku master
$ heroku open
```
