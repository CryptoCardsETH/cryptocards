import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllUsers } from '../actions/users';

class AllUsers extends Component {
  componentDidMount() {
    this.props.fetchAllUsers();
  }
  render() {
    // fake user data
    return (
      <div>
        <h1>All Users</h1>
        {this.props.user.all_users.map(u => (
          <div>
            <h3>
              {u.id}: {u.address}
            </h3>
            {JSON.stringify(u, true, 2)}
          </div>
        ))}
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { user: state.user };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchAllUsers }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(AllUsers);
