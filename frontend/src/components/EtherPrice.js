import React, { Component } from 'react';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faTag from '@fortawesome/fontawesome-free-solid/faTag';

class EtherPrice extends Component {
  convertToEth(value) {
    return value / Math.pow(10, 13);
  }
  render() {
    return (
      <div className="ether-price">
        <p className="">{this.convertToEth(this.props.price)} Eth</p>
        <FontAwesomeIcon className="price-icon" icon={faTag} />
      </div>
    );
  }
}

export default EtherPrice;
