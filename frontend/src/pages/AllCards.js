import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllCards } from '../actions/cards';
import CardFilterSort from '../components/CardFilterSort';
import { CARD_TYPE_COLLECTION } from '../components/Card';

class CollectionPage extends Component {
  componentDidMount() {
    this.props.fetchAllCards();
  }
  render() {
    // fake card data
    return (
      <div>
        <h1>All Cards</h1>
        <CardFilterSort filterSortKey="allcards" />
        <CardGrid
          cards={this.props.card.all_cards}
          filter={this.props.card.filters['allcards']}
          type={CARD_TYPE_COLLECTION}
        />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { card: state.card };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchAllCards }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(CollectionPage);
