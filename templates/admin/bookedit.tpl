{{ define "admin/bookedit.tpl" }}
{{template "header" . }}

	<form method="post" action="/admin/books/{{ .book.ID }}/edit" enctype="multipart/form-data">
        <table class="table">
            <tr>
                <th>ISBN</th>
                <td><input type="text" name="isbn" value="{{ .book.ISBN }}" readOnly="true"></td>
            </tr>
            <tr>
                <th>Title</th>
                <td><input type="text" name="title" value="{{ .book.Title }}" required></td>
            </tr>
            <tr>
                <th>Author</th>
                <td><input type="text" name="author" value="{{ .book.Author }}" required></td>
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
        </table>
        <input type="submit" name="save_change" value="Change" class="btn btn-primary">
        <a href="/admin/books" class="btn btn-default">Cancel</a>
    </form>
    <br/>


{{template "footer" . }}

{{ end }}

