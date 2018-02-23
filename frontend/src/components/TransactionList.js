import React, { Component } from 'react';
import 'animate.css';
import { withRouter } from 'react-router-dom';

class TransactionList extends Component {
  render() {
    let { transactions } = this.props;
    return (
      <div>
        <div className="transaction">
          <div className="table-responsive">
            <table className="table table-condensed table-bordered">
              <thead>
                <tr>
                  <th>Transaction Id</th>
                  <th>Transaction Time</th>
                  <th>Card Id</th>
                  <th>Transaction Price</th>
                </tr>
              </thead>
              {transactions.map((transaction, index) => {
                return (
                  <tbody key={index}>
                    <tr>
                      <td
                        class="transaction-id"
                        class="active"
                        rowspan="index + 1"
                      >
                        <div>{transaction.id}</div>
                      </td>
                      <td class="transaction-time">
                        <div>{transaction.created_at}</div>
                      </td>
                      <td class="transaction-cardId">
                        <div>{transaction.card_id}</div>
                      </td>
                      <td class="transaction-price">
                        <div>{transaction.price}</div>
                      </td>
                    </tr>
                  </tbody>
                );
              })}
            </table>
          </div>
        </div>
      </div>
    );
  }
}

export default withRouter(TransactionList);
