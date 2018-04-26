import React, { Component } from 'react';
import { Card, CardBody, CardTitle, CardDeck } from 'reactstrap';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchEntranceFeeReport } from '../actions/stats.js';
import TimeGraph from '../components/TimeGraph.js';

class EntranceFeeReportPage extends Component {
  componentDidMount() {
    this.props.fetchEntranceFeeReport();
  }
  render() {
    let report = this.props.stats.entrance_fee_report;
    let intervalLength = 1000 * 60 * 60 * 24; // one day in milliseconds
    return (
      <div className="container-fluid transaction-report">
        <h1>Entrance Fee Report</h1>
        <hr />
        <CardDeck>
          <Card className="col-lg-4 col-md-6 ">
            <CardBody>
              <CardTitle>Volume</CardTitle>
              <p className="text-center price">{report.volume}</p>
            </CardBody>
          </Card>
          <Card className="card col-lg-4 col-md-6 ">
            <CardBody>
              <CardTitle>Lowest Entrance Fee</CardTitle>
              <p className="text-center price">{report.min_entrance_fee}</p>
            </CardBody>
          </Card>
          <Card className="card col-lg-4 col-md-6 ">
            <CardBody>
              <CardTitle>Average Entrance Fee</CardTitle>
              <p className="text-center price">{report.avg_entrance_fee}</p>
            </CardBody>
          </Card>
          <Card className="card col-lg-4 col-md-6 ">
            <CardBody>
              <CardTitle>Highest Entrance Fee</CardTitle>
              <p className="text-center price">{report.max_entrance_fee}</p>
            </CardBody>
          </Card>
        </CardDeck>
        <div className="row">
          <div className="col-xl-12 chart">
            <h3 className="text-center">Entrance Fee Volume</h3>
            <TimeGraph
              data={report.entrance_fees}
              intervalLength={intervalLength}
              xAxisKey={'day'}
              yAxisKey={'fee_sum'}
              yAxisLabel={'Ethereum'}
            />
          </div>
        </div>
      </div>
    );
  }
}
function mapStateToProps(state) {
  return { stats: state.stats };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchEntranceFeeReport }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(
  EntranceFeeReportPage
);
