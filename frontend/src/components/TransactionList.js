import PropTypes from 'prop-types';
import React, { Component } from 'react';
import 'animate.css';
import { withRouter } from 'react-router-dom';

class TransactionList extends Component {
  render() {
    let { transactions } = this.props;
    let showCardId = this.props.showCardId;
    return (
      <div>
        <div className="transaction">
          <div className="table-responsive">
            <table className="table table-condensed table-bordered">
              <thead>
                <tr>
                  <th>Transaction Id</th>
                  <th>Transaction Time</th>
                  {showCardId ? <th>Card Id</th> : null}
                  <th>Transaction Price</th>
                </tr>
              </thead>
              {transactions.map((transaction, index) => {
                return (
                  <tbody key={index}>
                    <tr>
                      <td className="transaction-id" rowspan="index + 1">
                        <div>{transaction.id}</div>
                      </td>
                      <td className="transaction-time">
                        <div>{transaction.created_at}</div>
                      </td>
                      {showCardId ? (
                        <td className="transaction-cardId">
                          <div>{transaction.card_id}</div>
                        </td>
                      ) : null}
                      <td className="transaction-price">
                        <div>{transaction.price}</div>
                      </td>
                    </tr>
                  </tbody>
                );
              })}
            </table>
            {transactions.length < 1 ? (
              <p className="text-center">
                <i>No Transactions to display</i>
              </p>
            ) : null}
          </div>
        </div>
      </div>
    );
  }
}

TransactionList.propTypes = {
  showCardId: PropTypes.bool.isRequired
};

export default withRouter(TransactionList);
