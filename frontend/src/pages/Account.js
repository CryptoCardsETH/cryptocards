import React, { Component } from 'react';
import '../styles/App.scss';
import { connect } from 'react-redux';
import { Button } from 'reactstrap';
import { bindActionCreators } from 'redux';
import FollowList from '../components/FollowList';
import {
  fetchMe,
  editMeDetails,
  updateMe,
  fetchMyFollowers,
  fetchUserDetail,
  fetchMyFollowings
} from '../actions/users';

//Account page for update user's info
class AccountPage extends Component {
  constructor(props) {
    super(props);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  componentDidMount() {
    this.props.fetchMe();
    this.props.fetchMyFollowers();
    this.props.fetchMyFollowings();
  }

  handleSubmit(e) {
    e.preventDefault();
    this.props.updateMe();
  }

  render() {
    let { me } = this.props.user;

    return (
      <div className="container">
        <div className="row">
          <div className="col-md-9">
            <h1>My Account</h1>
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
                    value={me.nickname}
                    onChange={e => {
                      this.props.editMeDetails('nickname', e.target.value);
                    }}
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
                    value={me.email}
                    onChange={e => {
                      this.props.editMeDetails('email', e.target.value);
                    }}
                  />
                </div>
              </div>
              <div className="form-group">
                <br />
                <div className="col-md-8">
                  <Button type="submit" value="submit">
                    Save Changes
                  </Button>
                </div>
              </div>
            </form>
          </div>
          <br />
          <div className="col-md-9">
            <div className="float-left col-lg-4">
              <h4> Followers </h4>
              <FollowList follow={this.props.user.followers} />
            </div>
            <div className="float-right col-lg-4">
              <h4> Followings </h4>
              <FollowList follow={this.props.user.followings} />
            </div>
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
  return bindActionCreators(
    {
      fetchMe,
      updateMe,
      editMeDetails,
      fetchMyFollowers,
      fetchUserDetail,
      fetchMyFollowings
    },
    dispatch
  );
};

export default connect(mapStateToProps, mapDispatchToProps)(AccountPage);
