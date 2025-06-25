{{template "adminheader" . }}

	<p class="lead" style="margin: 25px 0"><a href="/admin/order">Orders</a></p>
      <div class="row">
        <div class="col-md-6">
          <h4>Details</h4>
          <table class="table">
            <tr>
			  <td>Date</td>
			  <td>{{ .order.Dt }}</td>
			</tr>
			<tr>
			  <td><a href="/admin/client/{{ .order.ClientID }}">Client fio</a></td>
			  <td><a href="/admin/client/{{ .order.ClientID }}">{{ .order.ClientFIO }}</a></td>
			</tr>
			<tr>
              <td><a href="/admin/client/{{ .order.ClientID }}">Client phone</a></td>
              <td><a href="/admin/client/{{ .order.ClientID }}">{{ .order.ClientPhone }}</a></td>
            </tr>
            <tr>
              <td>Amount</td>
              <td>{{ .order.Amount }}</td>
            </tr>
            <tr>
              <td>Qty</td>
              <td>{{ .order.Qty }}</td>
            </tr>
            <tr>
              <td>Ship name</td>
              <td>{{ .order.Ship.Name }}</td>
            </tr>
            <tr>
              <td>Ship Address</td>
              <td>{{ .order.Ship.Address }}</td>
            </tr>
            <tr>
              <td>Ship City</td>
              <td>{{ .order.Ship.City }}</td>
            </tr>
            <tr>
              <td>Ship ZipCode</td>
              <td>{{ .order.Ship.ZipCode }}</td>
            </tr>
            <tr>
              <td>Ship Country</td>
              <td>{{ .order.Ship.Country }}</td>
            </tr>
          </table>
       	</div>
      </div>
        {{ if eq .order.Status "N" }}
            <a class="btn btn-primary" href="/admin/order/{{ .order.ID }}/pay">Pay</a>
        {{ end }}
        {{ if eq .order.Status "P" }}
            <a class="btn btn-primary"  href="/admin/order/{{ .order.ID }}/send">Send</a>
        {{ end }}
      <div class="row">
          <div class="col-md-6">
            <h4>Items</h4>

            <table class="table" style="margin-top: 20px">
            <tr>
                <th>Title</th>
                <th>Author</th>
                <th>Price</th>
                <th>Qty</th>
            </tr>
            {{ range $index, $value := .order.Items }}
            <tr>
                <td>{{ .BookTitle }}</td>
                <td>{{ .BookAuthor }}</td>
                <td>{{ .Price }}</td>
                <td>{{ .Qty }}</td>
            </tr>
            {{ end }}
            </table>
       	</div>
      </div>

{{template "footer" . }}



