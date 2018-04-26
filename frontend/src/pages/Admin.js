import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllCountStats } from '../actions/stats.js';
import { Button, Card, CardBody, CardTitle } from 'reactstrap';
import '../styles/App.scss';

class AdminPage extends Component {
  componentDidMount() {
    this.props.fetchAllCountStats();
  }
  render() {
    return (
      <div className="container">
        <h1>Admin</h1>

        {this.props.stats.count_stats_loading ? (
          <p>Loading...</p>
        ) : (
          <div>
            <div className="stats row">
              <Card className="col-lg-3 col-md-5 m-2">
                <CardBody>
                  <CardTitle>Total Users</CardTitle>
                  <h1 className="text-center">
                    {this.props.stats.count_stats.user}
                  </h1>
                </CardBody>
              </Card>
              <Card className="card col-lg-3 col-md-5 m-2">
                <CardBody>
                  <CardTitle>Total Cards</CardTitle>
                  <h1 className="text-center">
                    {this.props.stats.count_stats.card}
                  </h1>
                </CardBody>
              </Card>
              <Card className="card col-lg-3 col-md-5 m-2">
                <CardBody>
                  <CardTitle>Total Listings</CardTitle>
                  <h1 className="text-center">
                    {this.props.stats.count_stats.listing}
                  </h1>
                </CardBody>
              </Card>
            </div>
            <div className="row">
              <Link to="/report/transactions">
                <Button>View Transaction Report</Button>
              </Link>
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
