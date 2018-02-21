import React, { Component } from 'react';
import '../styles/App.scss';
import { updateMe } from '../actions/users';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchMe } from '../actions/users';

//Account page for update user's info
class AccountPage extends Component {
  constructor(props) {
    super(props);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.onChangeEmail = this.onChangeEmail.bind(this);
    this.onChangeNickName = this.onChangeNickName.bind(this);
    this.state = { showresult: false };
  }

  componentDidMount() {
    this.props.fetchMe();
  }

  onChangeEmail(e) {
    const state = this.props.user;
    state[e.target.name] = e.target.value;
    this.setState(state);
  }

  onChangeNickName(e) {
    const state = this.props.user;
    state[e.target.name] = e.target.value;
    this.setState(state);
  }

  handleSubmit(e) {
    e.preventDefault();
    const data = this.state;
    this.setState({ showresult: true });
    this.props.updateMe(data);
  }

  render() {
    return (
      <div className="container">
        <div className="row">
          <div className="col-md-9">
            <h4>My Account</h4>
            <br />
            <form className="form-horizontal" onSubmit={this.handleSubmit}>
              <div className="form-group">
                <label className="col-lg-3 control-label">Nickname:</label>
                <div className="col-lg-8">
                  <input
                    className="form-control"
                    placeholder="Nickname"
                    name="nickname"
                    type="text"
                    defaultValue={
                      this.props.user.me.nickname == null
                        ? ''
                        : this.props.user.me.nickname
                    }
                    onChange={this.onChangeNickName}
                    autoFocus
                  />
                </div>
              </div>

              <div className="form-group">
                <label className="col-lg-3 control-label">Email Address:</label>
                <div className="col-lg-8">
                  <input
                    className="form-control"
                    placeholder="E-mail"
                    name="email"
                    type="email"
                    defaultValue={
                      this.props.user.me.email == null
                        ? ''
                        : this.props.user.me.email
                    }
                    onChange={this.onChangeEmail}
                  />
                </div>
              </div>
              <div className="form-group">
                <label className="col-md-3 control-label" />
                <div className="col-md-8">
                  <button> Save Changes</button>
                  <span className="tab-space">
                    {' '}
                    {this.state.showresult ? 'Changes Saved!' : ''}{' '}
                  </span>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { user: state.user };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchMe, updateMe }, dispatch);
};

export default connect(mapStateToProps, mapDispatchToProps)(AccountPage);
