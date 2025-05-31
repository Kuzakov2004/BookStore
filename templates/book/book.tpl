{{template "header" . }}

	<p class="lead" style="margin: 25px 0"><a href="/book/">Books</a> > {{ .book.Title }}</p>
      <div class="row">
        <div class="col-md-3 text-center">
          <img class="img-responsive img-thumbnail" src="/images/{{ .book.ID}}.jpg">
        </div>
        <div class="col-md-6">
          <h4>Book Description</h4>
          <p>{{ .book.Descr }}</p>
          <h4>Book Details</h4>
          <table class="table">
            <tr>
			  <td>Title</td>
			  <td>{{ .book.Title }}</td>
			</tr>
			<tr>
			  <td>Genre</td>
			  <td>{{ .book.Genre }}</td>
			</tr>
			<tr>
			  <td>Author</td>
			  <td>{{ .book.Author }}</td>
			</tr>
			<tr>
			  <td>Price</td>
			  <td>{{ .book.Price }}</td>
			</tr>
            <tr>
              <td>ISBN</td>
              <td>{{ .book.ISBN }}</td>
            </tr>
          </table>
          <form method="post" action="/cart/">
            <input type="hidden" name="id" value="{{ .book.ID }}">
            <input type="submit" value="Purchase / Add to cart" name="cart" class="btn btn-primary">
          </form>
       	</div>
      </div>

{{template "footer" . }}
