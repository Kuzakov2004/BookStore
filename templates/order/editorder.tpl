{{template "adminheader" . }}

	<p class="lead" style="margin: 25px 0"><a href="/admin/order">Orders</a></p>

	<span>
    {{ .err }}
    </span>
    <br>

    <ul class="nav nav-tabs" role="tablist">
        <li role="presentation" class="active"><a href="#client" aria-controls="client" role="tab" data-toggle="tab">Client</a></li>
        <li role="presentation"><a href="#ship" aria-controls="ship" role="tab" data-toggle="tab">Ship address</a></li>
        <li role="presentation"><a href="#books" aria-controls="books" role="tab" data-toggle="tab">Books</a></li>
    </ul>

    <!-- Tab panes -->
      <div class="tab-content">
        <div role="tabpanel" class="tab-pane active" id="client">
        {{ if eq .order.ClientID 0}}
            <br>
              <form action="/admin/order/selectclient" method="POST">
	          <input type="hidden" name="order" value="{{ .order.ID }}">
            <div class="input-group">
              <input type="text" class="form-control" name="str" aria-label="">
              <div class="input-group-btn">
                <button class="btn btn-primary" type="submit">Find client</button>
              </div>
            </div>
              </form>
        {{ else }}
        <br>
        <form action="/admin/order/selectclient" method="POST">
            <input type="hidden" name="order" value="{{ .order.ID }}">
            <button class="btn btn-primary" type="submit">Change client</button>
        </form>
        <br>
        <table class="table">
            <tr>
              <td><a href="/admin/client/{{ .order.ClientID }}">Client fio</a></td>
              <td><a href="/admin/client/{{ .order.ClientID }}">{{ .order.ClientFIO }}</a></td>
            </tr>
            <tr>
              <td><a href="/admin/client/{{ .order.ClientID }}">Client phone</a></td>
              <td><a href="/admin/client/{{ .order.ClientID }}">{{ .order.ClientPhone }}</a></td>
            </tr>
          </table>
        {{ end }}
        <div class="col-md-4">
        </div>
        <div class="col-md-4">
            <a class="btn btn-primary btnNext">Next</a>
        </div>

        <br>
        </div>
        <div role="tabpanel" class="tab-pane" id="ship">
        <br>
            <form class="form-horizontal" action="/admin/order/saveship" method="POST">
            <input type="hidden" name="order" value="{{ .order.ID }}">
            <div class="form-group">
                <label for="name" class="control-label col-md-4">Name</label>
                <div class="col-md-4">
                    <input type="text" name="name" class="form-control" value="{{ .order.Ship.Name}}">
                </div>
            </div>
            <div class="form-group">
                <label for="address" class="control-label col-md-4">Address</label>
                <div class="col-md-4">
                    <input type="text" name="address" class="form-control"  value="{{ .order.Ship.Address}}">
                </div>
            </div>
            <div class="form-group">
                <label for="city" class="control-label col-md-4">City</label>
                <div class="col-md-4">
                    <input type="text" name="city" class="form-control" value="{{ .order.Ship.City}}">
                </div>
            </div>
            <div class="form-group">
                <label for="zip" class="control-label col-md-4">ZipCode</label>
                <div class="col-md-4">
                    <input type="text" name="zip" class="form-control"  value="{{ .order.Ship.ZipCode}}">
                </div>
            </div>
            <div class="form-group">
                <label for="country" class="control-label col-md-4">Country</label>
                <div class="col-md-4">
                    <input type="text" name="country" class="form-control"  value="{{ .order.Ship.Country}}">
                </div>
            </div>
                <div class="col-md-4">
                    <a class="btn btn-primary btnPrev pull-right">Prev</a>
                </div>
                <div class="col-md-4">
                    <button class="btn btn-primary" type="submit">Save and Next</button>
                </div>
                <br>
            </form>
        </div>
        <div role="tabpanel" class="tab-pane" id="books">
            <br>
            <form action="/admin/order/selectbook" method="POST">
                <input type="hidden" name="order" value="{{ .order.ID }}">
                <button class="btn btn-primary" type="submit">Select books</button>
            </form>
            <br>
            <form class="form-horizontal" action="/admin/order/saveqty" method="POST">
            <input type="hidden" name="order" value="{{ .order.ID }}">
            <table class="table" style="margin-top: 20px">
                <tr>
                    <th>Title</th>
                    <th>Author</th>
                    <th>Price</th>
                    <th>Qty</th>
                    <th>In stock</th>
                    {{ if eq .order.Status "N" }}
                        <th>&nbsp;</th>
                    {{ end }}
                </tr>
                {{$oid := .order.ID}}
                {{$os := .order.Status}}
                {{ range $index, $value := .order.Items }}
                <tr>
                    <td>{{ .BookTitle }}</td>
                    <td>{{ .BookAuthor }}</td>
                    <td>{{ .Price }}</td>
                    <td>
                        <input type="hidden" name="book_id" value="{{ .BookID }}">
                        <input type="text" name="qty" class="form-control" value="{{ .Qty }}">
                    </td>
                    <td>{{ .InStock }}</td>
                    {{ if eq $os "N" }}
                    <td>
                        <a href="/admin/order/{{ $oid }}/delbook?book={{.BookID}}">Delete</a>
                    </td>
                    {{ end }}
                </tr>
                {{ end }}
            </table>
            <div class="col-md-4">
                <a class="btn btn-primary btnPrev pull-right">Prev</a>
            </div>
            <div class="col-md-4">
                <button class="btn btn-primary" type="submit">Save</button>
                <button class="btn btn-primary" type="submit" name="pay" value="pay">Pay</button>
            </div>
            </form>
        </div>
      </div>


<script  type="text/javascript">
$('.btnNext').click(function(){
  $('.nav-tabs > .active').next('li').find('a').trigger('click');
});
$('.btnPrev').click(function(){
  $('.nav-tabs > .active').prev('li').find('a').trigger('click');
});

$(function(){
   var url = window.location.href;
   if (url.indexOf("#") > 0) {
       var activeTab = url.substring(url.indexOf("#") + 1);

       $(".tab-pane").removeClass("active");
       $("#" + activeTab).addClass("active");
       $('a[href="#'+ activeTab +'"]').tab('show')
   }
});
</script>

{{template "footer" . }}



