{{ define "adminheader" }}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>{{.title}}</title>

    <link href="/bootstrap/css/bootstrap.min.css" rel="stylesheet">
    <link href="/bootstrap/css/bootstrap-theme.min.css" rel="stylesheet">
    <link href="/bootstrap/css/jumbotron.css" rel="stylesheet">
  </head>

  <body>

    <nav class="navbar navbar-light bg-light navbar-fixed-top">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="/">IT Bookstore</a>
          <span class="navbar-brand">Administration</span>
        </div>

        <!--/.navbar-collapse -->
        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav navbar-right">
              <li><a href="/admin/publisher"><span class="glyphicon glyphicon-paperclip"></span>&nbsp; Publishers</a></li>
              <li><a href="/admin/book"><span class="glyphicon glyphicon-book"></span>&nbsp; Books</a></li>
              <li><a href="/admin/order"><span class="glyphicon glyphicon-shopping-cart"></span>&nbsp; Orders</a></li>
              <li><a href="/admin/warehouse"><span class="glyphicon glyphicon-home"></span>&nbsp; Warehouses</a></li>
              <li><a href="/admin/contact"><span class="glyphicon glyphicon-phone-alt"></span>&nbsp; Contact</a></li>
              <li><a href="/admin/logout" class="btn btn-primary">Sign out!</a></li>
           </ul>
        </div>
      </div>
    </nav>

    <div class="container" id="main">
{{ end }}
