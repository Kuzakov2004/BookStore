{{template "adminheader" . }}

    <p class="lead text-center text-muted">Warehouses</p>

	<table class="table" style="margin-top: 20px">
        <tr>
            <th>Address</th>
            <th>Capacity</th>
        </tr>
        {{ range $index, $value := .warehouses }}
        <tr>
            <td><a href="/admin/warehouse/{{ .ID }}/books">{{ .Address }}</a></td>
            <td>{{ .Capacity }}</td>
        </tr>
        {{ end }}
    </table>

{{template "footer" . }}
