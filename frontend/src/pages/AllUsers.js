import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllUsers } from '../actions/users';
import { buildProfileURL } from '../actions';
import { Link } from 'react-router-dom';
class AllUsers extends Component {
  componentDidMount() {
    this.props.fetchAllUsers();
  }
  render() {
    // fake user data
    return (
      <div className="container">
        <h1>All Users</h1>
        {this.props.user.all_users.map(u => {
          let { cards, nickname, email, address } = u;
          return (
            <div>
              <h3>
                #{u.id}: {address}
              </h3>
              <Link to={buildProfileURL(u)}>view user</Link>
              <pre>
                {JSON.stringify({ cards, nickname, email, address }, true, 2)}
              </pre>
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
  return bindActionCreators({ fetchAllUsers }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(AllUsers);
