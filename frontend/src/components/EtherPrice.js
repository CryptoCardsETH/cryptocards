import React, { Component } from 'react';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faTag from '@fortawesome/fontawesome-free-solid/faTag';
import PropTypes from 'prop-types';

class EtherPrice extends Component {
  render() {
    return (
      <div className="ether-price">
        <span>{this.props.price} ETH</span>
        <FontAwesomeIcon className="price-icon" icon={faTag} />
      </div>
    );
  }
}

EtherPrice.propTypes = {
  price: PropTypes.number.isRequired
};

export default EtherPrice;
