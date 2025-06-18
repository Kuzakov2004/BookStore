{{ define "admin/books.tpl" }}
{{template "header" . }}

	<p class="lead"><a href="/admin/book/create">Add new book</a></p>
    	<table class="table" style="margin-top: 20px">
    		<tr>
    			<th>ISBN</th>
    			<th>Title</th>
    			<th>Genre</th>
    			<th>Author</th>
    			<th>Price</th>
    			<th>&nbsp;</th>
    			<th>&nbsp;</th>
    		</tr>
    		{{ range $index, $value := .books }}
    		<tr>
    			<td>{{ .ISBN }}</td>
    			<td>{{ .Title }}</td>
    			<td>{{ .Genre }}</td>
    			<td>{{ .Author }}</td>
    			<td>{{ .Price }}</td>
    			<td><a href="/admin/book/{{ .ID }}/edit">Edit</a></td>
    			<td><a href="/admin/book/{{ .ID }}/delete">Delete</a></td>
    		</tr>
    		{{ end }}
    	</table>


{{template "footer" . }}

{{ end }}