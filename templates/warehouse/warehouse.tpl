{{template "adminheader" . }}

	<p class="lead" style="margin: 25px 0"><a href="/admin/warehouse">Warehouse</a> > {{ .warehouse.Address }}</p>
      <div class="row">
        <div class="col-md-6">
          <h4>Address</h4>
          <p>{{ .warehouse.Address }}</p>
          <p>{{ .warehouse.Capacity }}</p>
       	</div>
      </div>

{{template "footer" . }}
