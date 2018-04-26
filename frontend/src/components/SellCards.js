import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { sellCards } from '../actions/listings';
import PropTypes from 'prop-types';
import { Button } from 'reactstrap';
import CardValue from './CardValue';
class SellCards extends React.Component {
  render() {
    let { cardIds } = this.props;
    let isValidSize = cardIds.length >= 1;
    return (
      <div>
        <div class="row">
          {cardIds.map(cardId => {
            return (
              <div class="col-md-2">
                <CardValue cardId={cardId} />
              </div>
            );
          })}
        </div>
        <br />
        <Button
          disabled={!isValidSize}
          onClick={() => this.props.sellCards(cardIds)}
        >
          Sell cards
        </Button>
      </div>
    );
  }
}

SellCards.propTypes = {
  carddIds: PropTypes.array.isRequired
};

function mapStateToProps(state) {
  let { card, user } = state;
  return {
    card,
    user
  };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators(
    {
      sellCards
    },
    dispatch
  );
};
export default connect(mapStateToProps, mapDispatchToProps)(SellCards);
