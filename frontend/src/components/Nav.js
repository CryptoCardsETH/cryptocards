import React from 'react';
import { NavLink } from 'react-router-dom';
import { APP_NAME } from '../config';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { logout } from '../actions/users';

class Nav extends React.Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.state = {
      isOpen: false
    };
  }
  toggle() {
    this.setState({
      isOpen: !this.state.isOpen
    });
  }
  render() {
    return (
      <div>
        <nav className="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
          <div className="navbar-brand">{APP_NAME}</div>
          <div className="collapse navbar-collapse" id="navbarCollapse">
            <ul className="navbar-nav mr-auto">
              <li className="nav-item active">
                <NavLink exact to="/" className="nav-link">
                  Home
                </NavLink>
              </li>
              {this.props.user.authenticated ? (
                <li className="nav-item">
                  <NavLink to="/collection" className="nav-link">
                    My Collection
                  </NavLink>
                </li>
              ) : null}
              <li className="nav-item">
                <NavLink to="/cards" className="nav-link">
                  All Cards
                </NavLink>
              </li>
              <li className="nav-item">
                <NavLink to="/debug" className="nav-link">
                  Debug
                </NavLink>
              </li>
              {this.props.user.authenticated ? (
                <li className="nav-item">
                  <div onClick={() => this.props.logout()} className="nav-link">
                    Logout
                  </div>
                </li>
              ) : null}
            </ul>
          </div>
        </nav>
      </div>
    );
  }
}

function mapStateToProps(state) {
  let { user } = state;
  return { user };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators(
    {
      logout
    },
    dispatch
  );
};

export default connect(mapStateToProps, mapDispatchToProps)(Nav);
