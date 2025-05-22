<!doctype html>

<html>
    <head>
        <title>{{.title}}</title>
        {{include "layouts/header"}}
    </head>

    <body>
        <a href="/"><- Back home!</a>
        {{template "booklist" .}}
        <hr>
        {{include "layouts/footer"}}
    </body>
</html>
