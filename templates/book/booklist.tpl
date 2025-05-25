{{ define "books/booklist" }}


<p class="lead text-center text-muted">Full Catalogs of Books</p>

	<div class="row">
{{ range $index, $value := .books }}
		<div class="col-md-3">
			<a href="/book/{{ .ID }}">
				<img class="img-responsive img-thumbnail" src="/images/{{ .ID }}.jpg">
			</a>
		</div>

	{{ if each1 $index 4 }}
	</div>
	<div class="row">
	{{ end }}

{{ end }}
	</div>
{{ end }}
