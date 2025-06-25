{{ define "admin/selectclient.tpl" }}
{{template "adminheader" . }}
	<p class="lead">Select client</p>
    <form action="/admin/order/selectclient" method="POST">
	<div class="input-group">
	  <input type="hidden" name="order" value="{{ .orderID }}">
      <input type="text" class="form-control" name="str" value="{{ .str }}" aria-label="">
      <div class="input-group-btn">
        <button class="btn btn-default" type="submit">Find</button>
      </div>
    </div>
    </form>

    <form action="/admin/order/selectclientdone" method="POST">
	  <input type="hidden" name="order" value="{{ .orderID }}">
    	<table class="table" style="margin-top: 20px">
    		<tr>
    			<th>&nbsp;</th>
    			<th>FirstName</th>
    			<th>LastName</th>
    			<th>MiddleName</th>
    			<th>Login</th>
    			<th>Phone</th>
    		</tr>

    		{{ range $index, $value := .clients }}
    		<tr>
    		    <td><input type="radio" value="{{ .ID }}" name="client_id"></td>
    			<td>{{ .FirstName }}</td>
    			<td>{{ .LastName }}</td>
    			<td>{{ .MiddleName }}</td>
    			<td>{{ .Login }}</td>
    			<td>{{ .Phone }}</td>
    		</tr>
    		{{ end }}
    	</table>
    	<button class="btn btn-default" type="submit">Select</button>
    	<a class="btn btn-default" href="/admin/order/{{ .orderID }}/edit#client">Cancel</a>
    	</form>


{{template "footer" . }}

{{ end }}