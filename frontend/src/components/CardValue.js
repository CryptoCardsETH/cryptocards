import PropTypes from 'prop-types';
import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchCardValue } from '../actions/cards';
import EtherPrice from './EtherPrice';

class CardValue extends Component {
  componentDidMount() {
    this.props.fetchCardValue(this.props.cardId);
  }
  render() {
    return (
      <span>
        Recomended Price:{' '}
        {!this.props.card.card_value_loading ? (
          <EtherPrice price={this.props.card.card_value} />
        ) : null}
      </span>
    );
  }
}

function mapStateToProps(state) {
  return { card: state.card };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchCardValue }, dispatch);
};

CardValue.propTypes = {
  cardId: PropTypes.number.isRequired
};
export default connect(mapStateToProps, mapDispatchToProps)(CardValue);
