{{ define "admin/publisheredit.tpl" }}
{{template "adminheader" . }}

    <span>
    {{ .err }}
    </span>
    <br>
    {{ if .isCreate }}
    <form method="post" action="/admin/publisher/create" enctype="multipart/form-data">
    {{ else }}
	<form method="post" action="/admin/publisher/{{ .publisher.ID }}/edit" enctype="multipart/form-data">
    {{ end }}
        <table class="table">
            <tr>
                <th>Name</th>
                <td><input type="text" name="name" value="{{ .publisher.Name }}" required></td>
            </tr>
            <tr>
                <th>Country</th>
                <td><input type="text" name="country" value="{{ .publisher.Country }}" required></td>
            </tr>
            <tr>
                <th>Phone</th>
                <td><input type="text" name="phone" value="{{ .publisher.Phone }}" required></td>
            </tr>
        </table>
        <input type="submit" name="save_change" value="Change" class="btn btn-primary">
        <a href="/admin/publisher" class="btn btn-default">Cancel</a>
    </form>
    <br/>


{{template "footer" . }}

{{ end }}

