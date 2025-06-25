{{ define "admin/selectbook.tpl" }}
{{template "adminheader" . }}
	<p class="lead">Select book</p>

    <form action="/admin/order/selectbookdone" method="POST">
	  <input type="hidden" name="order" value="{{ .orderID }}">
    	<table class="table" style="margin-top: 20px">
            <tr>
                <th>&nbsp;</th>
                <th>ISBN</th>
                <th>Title</th>
                <th>Genre</th>
                <th>Author</th>
                <th>Price</th>
            </tr>
            {{ range $index, $value := .books }}
            <tr>
                <td><input type="checkbox" value="{{ .ID }}" name="book_id"></td>
                <td>{{ .ISBN }}</td>
                <td>{{ .Title }}</td>
                <td>{{ .Genre }}</td>
                <td>{{ .Author }}</td>
                <td>{{ .Price }}</td>
            </tr>
            {{ end }}
        </table>
    	<button class="btn btn-default" type="submit">Select</button>
    	<a class="btn btn-default" href="/admin/order/{{ .orderID }}/edit#books">Cancel</a>
    	</form>


{{template "footer" . }}

{{ end }}