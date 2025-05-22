<!doctype html>

<html>
    <head>
        <title>{{.title}}</title>
        {{include "layouts/header"}}
    </head>
    <body>
        {{template "content" .}}
        <hr>
        {{template "ad" .}}
        <hr>
        {{include "layouts/footer"}}
    </body>
</html>
