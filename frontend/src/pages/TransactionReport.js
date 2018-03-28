import React, { Component } from 'react';
import { Card, CardBody, CardTitle, CardDeck } from 'reactstrap';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchTransactionReport } from '../actions/stats.js';
import TimeGraph from '../components/TimeGraph.js';

class TransactionReportPage extends Component {
  componentDidMount() {
    this.props.fetchTransactionReport();
  }
  render() {
    let report = this.props.stats.transaction_report;
    let intervalLength = 1000 * 60 * 60 * 24; // one day in milliseconds
    return (
      <div className="transaction-report">
        <h1>Transaction Report</h1>
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
              <CardTitle>Lowest Price</CardTitle>
              <p className="text-center price">{report.min_price}</p>
            </CardBody>
          </Card>
          <Card className="card col-lg-4 col-md-6 ">
            <CardBody>
              <CardTitle>Average Price</CardTitle>
              <p className="text-center price">{report.avg_price}</p>
            </CardBody>
          </Card>
          <Card className="card col-lg-4 col-md-6 ">
            <CardBody>
              <CardTitle>Highest Price</CardTitle>
              <p className="text-center price">{report.max_price}</p>
            </CardBody>
          </Card>
        </CardDeck>
        <div className="row">
          <div className="col-xl-12 chart">
            <h3 className="text-center">Transaction Volume</h3>
            <TimeGraph
              data={report.transactions}
              intervalLength={intervalLength}
              xAxisKey={'day'}
              yAxisKey={'price_sum'}
              yAxisLabel={'Ethereum'}
            />
          </div>
          <div className="col-xl-12 chart">
            <h3 className="text-center">Card Creation</h3>
            <TimeGraph
              data={report.cards}
              intervalLength={intervalLength}
              xAxisKey={'day'}
              yAxisKey={'num_cards'}
              yAxisLabel={'Card Counts'}
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
  return bindActionCreators({ fetchTransactionReport }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(
  TransactionReportPage
);
