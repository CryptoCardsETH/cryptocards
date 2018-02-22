import React, { Component } from 'react';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faTag from '@fortawesome/fontawesome-free-solid/faTag';
import PropTypes from 'prop-types';

class EtherPrice extends Component {
  convertToEth(value) {
    return value / Math.pow(10, 13);
  }
  render() {
    return (
      <div className="ether-price">
        <span>{this.convertToEth(this.props.price)} Eth</span>
        <FontAwesomeIcon className="price-icon" icon={faTag} />
      </div>
    );
  }
}

EtherPrice.propTypes = {
  price: PropTypes.number.isRequired
};

export default EtherPrice;
