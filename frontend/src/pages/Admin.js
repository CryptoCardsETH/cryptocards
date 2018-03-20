import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllCountStats } from '../actions/stats.js';
import '../styles/App.scss';

class AdminPage extends Component {
  componentDidMount() {
    this.props.fetchAllCountStats();
  }
  render() {
    return (
      <div>
        <h1>Admin</h1>

        {this.props.stats.count_stats_loading ? (
          <p>Loading...</p>
        ) : (
          <div>
            <div className="stats row">
              <div className="card col-lg-3 col-md-5 m-2">
                <div className="card-body">
                  <h4 className="card-title">Total Users</h4>
                  <h1 className="text-center">
                    {this.props.stats.count_stats.user}
                  </h1>
                </div>
              </div>
              <div className="card col-lg-3 col-md-5 m-2">
                <div className="card-body">
                  <h4 className="card-title">Total Cards</h4>
                  <h1 className="text-center">
                    {this.props.stats.count_stats.card}
                  </h1>
                </div>
              </div>
              <div className="card col-lg-3 col-md-5 m-2">
                <div className="card-body">
                  <h4 className="card-title">Total Listings</h4>
                  <h1 className="text-center">
                    {this.props.stats.count_stats.listing}
                  </h1>
                </div>
              </div>
            </div>
          </div>
        )}
      </div>
    );
  }
}
function mapStateToProps(state) {
  return { stats: state.stats };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchAllCountStats }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(AdminPage);
