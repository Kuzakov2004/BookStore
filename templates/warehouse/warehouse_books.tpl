{{template "adminheader" .}}

<p class="lead" style="margin: 25px 0">
    <a href="/admin/warehouse">Склады</a> &gt; {{ .warehouse.Address }}
</p>

<table class="table" style="margin-top: 20px">
    <tr>
        <th>ISBN</th>
        <th>Название</th>
        <th>Автор</th>
        <th>Жанр</th>
        <th>Издательство</th>
        <th>Цена</th>
        <th>Количество</th>
    </tr>

    {{ range .books }}
    <tr>
        <td>{{ .ISBN }}</td>
        <td>{{ .Title }}</td>
        <td>{{ .Author }}</td>
        <td>{{ .Genre }}</td>
        <td>{{ .Publisher }}</td>
        <td>{{ printf "%.2f" .Price }}</td>
        <td>{{ .QuantityOnStock }}</td>
    </tr>
    {{ else }}
    <tr>
        <td colspan="7">Нет книг на складе</td>
    </tr>
    {{ end }}
</table>

{{template "footer" .}}