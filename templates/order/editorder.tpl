{{template "adminheader" . }}

	<p class="lead" style="margin: 25px 0"><a href="/admin/order">Orders</a></p>

    <ul class="nav nav-tabs" role="tablist">
        <li role="presentation" class="active"><a href="#client" aria-controls="client" role="tab" data-toggle="tab">Client</a></li>
        <li role="presentation"><a href="#profile" aria-controls="profile" role="tab" data-toggle="tab">Profile</a></li>
        <li role="presentation"><a href="#messages" aria-controls="messages" role="tab" data-toggle="tab">Messages</a></li>
        <li role="presentation"><a href="#settings" aria-controls="settings" role="tab" data-toggle="tab">Settings</a></li>
    </ul>

    <!-- Tab panes -->
      <div class="tab-content">
        <div role="tabpanel" class="tab-pane active" id="client">
              <form action="/admin/order/selectclient" method="POST">
	          <input type="hidden" name="order" value="{{ .order.ID }}">
            <div class="input-group">
              <input type="text" class="form-control" name="str" aria-label="">
              <div class="input-group-btn">
                <button class="btn btn-default" type="submit">Find</button>
              </div>
            </div>
              </form>
        </div>
        <div role="tabpanel" class="tab-pane" id="profile">222222222222222222</div>
        <div role="tabpanel" class="tab-pane" id="messages">333333333333333333</div>
        <div role="tabpanel" class="tab-pane" id="settings">44444444444444444444444</div>
      </div>

{{template "footer" . }}



