{{template "header" . }}

    <p class="lead text-center text-muted">Publishers</p>

	<table class="table" style="margin-top: 20px">
        <tr>
            <th>Name</th>
            <th>Country</th>
            <th>Phone</th>
        </tr>
        {{ range $index, $value := .publishers }}
        <tr>
            <td>{{ .Name }}</td>
            <td>{{ .Country }}</td>
            <td>{{ .Phone }}</td>
        </tr>
        {{ end }}
    </table>

{{template "footer" . }}
