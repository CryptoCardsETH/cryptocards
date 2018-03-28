import React, { Component } from 'react';
import TransactionList from '../components/TransactionList';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchMyTransactions } from '../actions/users';

class TransactionPage extends Component {
  componentDidMount() {
    this.props.fetchMyTransactions();
  }
  render() {
    return (
      <div className="container">
        <h2>Transactions History</h2>
        <br />
        <TransactionList transactions={this.props.user.transactions} />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { user: state.user };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchMyTransactions }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(TransactionPage);
