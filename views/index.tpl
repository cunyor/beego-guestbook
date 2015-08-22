<!DOCTYPE html>
<html>
  <head>
    <title>Beego GuestBook</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/static/css/main.css" type="text/css" />
  </head>
  <body>
    <a href="https://github.com/yuhei0718/beego-guestbook"><img style="position: absolute; top: 0; right: 0; border: 0;" src="https://s3.amazonaws.com/github/ribbons/forkme_right_white_ffffff.png" alt="Fork me on GitHub"></a>

    <div id="main">
      <h1>Ziyaretci Defteri</h1>
      <div id="form-area">
        <p>Lütfen yazınız</p>
        <form action="/post" method="post">
          <table>
            <tr>
              <th>İsim</th>
              <td>
                <input type="text" size="20" name="name" />
              </td>
            </tr>
            <tr>
              <th>Yorum</th>
              <td>
                <textarea rows="5" cols="40" name="comment"></textarea>
              </td>
            </tr>
          </table>
          <p><button type="submit">Gönder</button></p>
        </form>
      </div>
      <div id="entries-area">
        <h2>Önceki yorumlar</h2>
        {{range $key, $val := .greetings}}
        <div class="entry">
          <h3>{{$val.Name}} yazdı: ({{$val.CreateAt|dateformat}}):</h3>
          <p>{{$val.Comment|htmlquote|nl2br|str2html}}</p>
        </div>
        {{end}}
      </div>
    </div>
  </body>
</html>
