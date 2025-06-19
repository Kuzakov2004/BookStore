{{ define "admin/publishers.tpl" }}
{{template "adminheader" . }}

	<p class="lead"><a href="/admin/publisher/create">Add new publisher</a></p>
	    <table class="table" style="margin-top: 20px">
            <tr>
                <th>Name</th>
                <th>Country</th>
                <th>Phone</th>
    			<th>&nbsp;</th>
    			<th>&nbsp;</th>
            </tr>
            {{ range $index, $value := .publishers }}
            <tr>
                <td>{{ .Name }}</td>
                <td>{{ .Country }}</td>
                <td>{{ .Phone }}</td>

    			<td><a href="/admin/publisher/{{ .ID }}/edit">Edit</a></td>
    			<td><a href="/admin/publisher/{{ .ID }}/delete">Delete</a></td>
            </tr>
            {{ end }}
        </table>

{{template "footer" . }}
{{ end }}