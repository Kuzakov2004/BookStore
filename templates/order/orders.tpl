{{template "adminheader" . }}

<p class="lead"><a href="/admin/order/create">Add new order</a></p>

<div>
  <nav>
    <div class="container">

        <div class="navbar-header">
            <ul class="nav  nav-pills nav-fill">
              <li class="nav-item {{ .activeActive }}">
                <a class="nav-link" href="/admin/order/?status=N">Active</a>
              </li>
              <li class="nav-item {{ .activePayed }}">
                <a class="nav-link" href="/admin/order/?status=P">Payed</a>
              </li>
              <li class="nav-item  {{ .activeCompleted }}">
                <a class="nav-link" href="/admin/order/?status=C">Completed</a>
              </li>
            </ul>
        </div>
      </div>
  </nav>
</div>

	<table class="table" style="margin-top: 20px">
        <tr>
            <th>Dt</th>
            <th>Client</th>
            <th>Client Phone</th>
            <th>Amount</th>
            <th>Qty</th>
            <th>Ship address</th>
            <th>&nbsp;</th>
            <th>&nbsp;</th>
        </tr>
        {{ range $index, $value := .orders }}
        <tr>
            <td>{{ .Dt }}</td>
            <td><a href="/admin/client/{{ .ClientID }}">{{ .ClientFIO }}</a></td>
            <td><a href="/admin/client/{{ .ClientID }}">{{ .ClientPhone }}</a></td>
            <td>{{ .Amount }}</td>
            <td>{{ .Qty }}</td>
            <td>{{ .Ship.Address }}</td>
            <td><a href="/admin/order/{{ .ID }}">Detail</a></td>
            <td><a href="/admin/order/{{ .ID }}/edit">Edit</a></td>
        </tr>
        {{ end }}
    </table>

{{template "footer" . }}
