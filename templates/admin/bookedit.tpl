{{ define "admin/bookedit.tpl" }}
{{template "adminheader" . }}

    <span>
    {{ .err }}
    </span>
    <br>
    {{ if .isCreate }}
    <form method="post" action="/admin/book/create" enctype="multipart/form-data">
            <table class="table">
                <tr>
                    <th>ISBN</th>
                    <td><input type="text" name="isbn" value="{{ .book.ISBN }}"></td>
                </tr>
    {{ else }}
	<form method="post" action="/admin/book/{{ .book.ID }}/edit" enctype="multipart/form-data">
        <table class="table">
            <tr>
                <th>ISBN</th>
                <td><input type="text" name="isbn" value="{{ .book.ISBN }}" readOnly="true"></td>
            </tr>
    {{ end }}
            <tr>
                <th>Title</th>
                <td><input type="text" name="title" value="{{ .book.Title }}" required></td>
            </tr>
            <tr>
                <th>Author</th>
                <td>
                {{$a := .book.AuthorID}}
                <select class="form-select" aria-label="Author" name="author_id" value="{{ .book.AuthorID }}">
                {{ range $index, $value := .authors }}
                  <option value="{{ .ID }}" {{ if eq $a .ID }} selected {{ end }}>{{ .FIO }}</option>
                {{ end }}
                </select>
                </td>
            </tr>
            <tr>
                <th>Publisher</th>
                <td>
                {{$p := .book.PublisherID}}
                <select class="form-select" aria-label="Author" name="publisher_id" value="{{ .book.PublisherID }}">
                {{ range $index, $value := .publishers }}
                  <option value="{{ .ID }}" {{ if eq $p .ID }} selected {{ end }}>{{ .Name }}</option>
                {{ end }}
                </select>
                </td>
            </tr>
            <tr>
                <th>Image</th>
                <td><input type="file" name="image"></td>
            </tr>
            <tr>
                <th>Description</th>
                <td><textarea name="descr" cols="40" rows="5">{{ .book.Descr }}</textarea>
            </tr>
            <tr>
                <th>Price</th>
                <td><input type="text" name="price" value="{{ .book.Price }}" required></td>
            </tr>
            <tr>
                <th>Publication Year</th>
                <td><input type="text" name="publication_year" value="{{ .book.PublicationYear }}" required></td>
            </tr>
            <tr>
                <th>Genre</th>
                <td><input type="text" name="genre" value="{{ .book.Genre }}" required></td>
            </tr>
        </table>
        <input type="submit" name="save_change" value="Change" class="btn btn-primary">
        <a href="/admin/book" class="btn btn-default">Cancel</a>
    </form>
    <br/>


{{template "footer" . }}

{{ end }}

