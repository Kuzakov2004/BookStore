{{template "header" . }}

	<p class="lead" style="margin: 25px 0"><a href="/book/">Books</a> > {{ .book.Title }}</p>
      <div class="row">
        <div class="col-md-6">
          <h4>Name</h4>
          <p>{{ .publisher.Name }}</p>
          <h4>Details</h4>
          <table class="table">
            <tr>
			  <td>Country</td>
			  <td>{{ .publisher.Country }}</td>
			</tr>
			<tr>
			  <td>Phone</td>
			  <td>{{ .publisher.Phone }}</td>
			</tr>
          </table>
       	</div>
      </div>

{{template "footer" . }}
