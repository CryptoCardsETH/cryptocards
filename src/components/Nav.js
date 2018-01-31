import React from 'react';
import { NavLink } from 'react-router-dom';
import { APP_NAME } from '../config';

export default class Nav extends React.Component {
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
          <a className="navbar-brand">{APP_NAME}</a>
          <div className="collapse navbar-collapse" id="navbarCollapse">
            <ul className="navbar-nav mr-auto">
              <li className="nav-item active">
                <NavLink exact to="/" className="nav-link">
                  Home
                </NavLink>
              </li>
              <li className="nav-item">
                <NavLink to="/debug" className="nav-link">
                  Debug
                </NavLink>
              </li>
            </ul>
          </div>
        </nav>
      </div>
    );
  }
}
