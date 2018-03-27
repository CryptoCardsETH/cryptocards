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

  /**
   * Creates a new Array of data, but inserts missing days with a price of 0.
   */
  addMissingDays(data) {
    let allDays = this.getAllDays(data);
    let i = 0,
      j = 0;
    while (i < allDays.length) {
      let tmpDay = new Date(data[j].day);
      if (j < data.length && allDays[i].date.getTime() === tmpDay.getTime()) {
        allDays[i].price_sum = data[j].price_sum;
        j++;
      } else {
        allDays[i].price_sum = 0;
      }
      allDays[i].day = `${allDays[i].date.getFullYear()}-${this.pad(
        allDays[i].date.getMonth() + 1
      )}-${this.pad(allDays[i].date.getUTCDate())}`;
      i++;
    }
    return allDays;
  }

  /**
   * Gets an Array of every date in the range of the data.
   */
  getAllDays(data) {
    if (!data) return [];
    let earliestDate = new Date(data[0].day);
    let latestDate = new Date(data[data.length - 1].day);
    let range = latestDate - earliestDate;
    let lengthOfDay = 1000 * 60 * 60 * 24;
    let numDays = range / lengthOfDay;
    return Array.from({ length: numDays + 1 }, (v, i) => {
      return { date: new Date(earliestDate.valueOf() + lengthOfDay * i) };
    });
  }

  pad(number) {
    return ('0' + number).slice(-2);
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
