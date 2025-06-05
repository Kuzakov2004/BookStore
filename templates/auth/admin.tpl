{{template "header" . }}

	<form class="form-horizontal" method="post" action="admin/login">
		<div class="form-group">
			<label for="name" class="control-label col-md-4">Login</label>
			<div class="col-md-4">
				<input type="text" name="login" class="form-control">
			</div>
		</div>
		<div class="form-group">
			<label for="pass" class="control-label col-md-4">Pass</label>
			<div class="col-md-4">
				<input type="password" name="pass" class="form-control">
			</div>
		</div>
		<input type="submit" name="submit" class="btn btn-primary">
	</form>

{{template "footer" . }}
