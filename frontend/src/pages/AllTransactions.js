import React, { Component } from 'react';
import TransactionList from '../components/TransactionList';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchMyTransactions } from '../actions/users';

class TransactionPage extends Component {
  componentDidMount() {
    console.log('comes here1');
    this.props.fetchMyTransactions();
  }
  render() {
    console.log('comes here2');
    return (
      <div class="container">
        <h2>Transactions History</h2>
        <TransactionList transaction={this.props.user.transactions} />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { user: state.user, transaction: state.transaction };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchMyTransactions }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(TransactionPage);
