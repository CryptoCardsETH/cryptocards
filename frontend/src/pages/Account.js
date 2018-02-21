import React, { Component } from 'react';
import '../styles/App.scss';

class AccountPage extends Component {
  render() {
    return (
      <div class="container">
        <div class="row">
          <div class="col-md-9">
            <h4>My Account</h4>
            <br />
            <form
              class="form-horizontal"
              role="form"
              method="post"
              action="Account.php"
            >
              <div class="form-group">
                <label class="col-lg-3 control-label">Nickname:</label>
                <div class="col-lg-8">
                  <input
                    class="form-control"
                    placeholder="Nickname"
                    name="Nickname"
                    type="text"
                    autofocus
                  />
                </div>
              </div>

              <div class="form-group">
                <label class="col-lg-3 control-label">Email Address:</label>
                <div class="col-lg-8">
                  <input
                    class="form-control"
                    placeholder="E-mail"
                    name="email"
                    type="email"
                    autofocus
                  />
                </div>
              </div>
              <div class="form-group">
                <label class="col-md-3 control-label" />
                <div class="col-md-8">
                  <input
                    type="button"
                    class="btn btn-primary"
                    value="Save Changes"
                  />
                  <span class="tab-space"> </span>
                  <input type="reset" class="btn btn-default" value="Cancel" />
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    );
  }
}

export default AccountPage;
