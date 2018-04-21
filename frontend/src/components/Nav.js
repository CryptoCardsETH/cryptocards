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
    let isAdmin =
      this.props.user.authenticated &&
      this.props.user.me &&
      this.props.user.me.admin;
    return (
      <div className="navigation">
        <nav className="navbar navbar-expand-md fixed-top">
          <div className="navbar-brand">{APP_NAME}</div>
          <div className="collapse navbar-collapse" id="navbarCollapse">
            <ul className="navbar-nav ml-auto">
              <NavbarItem to="/" text="Home" />
              <NavbarItem to="/marketplace" text="Marketplace" />
              {this.props.user.authenticated ? (
                <NavbarItem to="/account" text="Account" />
              ) : null}
              {this.props.user.authenticated ? (
                <NavbarItem
                  to={'/user/' + (this.props.user.me && this.props.user.me.id)}
                  text="My Collection"
                />
              ) : null}
              <NavbarItem to="/users" text="All Users" />
              <NavbarItem to="/cards" text="All Cards" />
              <NavbarItem to="/debug" text="Debug" />
              <NavbarItem to="/faq" text="FAQ" />
              {isAdmin ? <NavbarItem to="/admin" text="Admin" /> : null}
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

const NavbarItem = ({ to, text }) => (
  <li className="nav-item">
    <NavLink to={to} className="nav-link">
      {text}
    </NavLink>
  </li>
);
