import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchMyNotifications } from '../actions/users';
import { buildProfileURL } from '../actions';
import { Link } from 'react-router-dom';
class Notifications extends Component {
  componentDidMount() {
    this.props.fetchMyNotifications();
  }
  render() {
    // fake user data
    return (
      <div>
        <h1>All Notifications</h1>
        {this.props.user.notifications.map(n => {
          return (
            <div>
              <pre>{JSON.stringify({ n }, true, 2)}</pre>
            </div>
          );
        })}
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { user: state.user };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchMyNotifications }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(Notifications);
