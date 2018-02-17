import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllCards } from '../actions/cards';

class CollectionPage extends Component {
  componentDidMount() {
    this.props.fetchAllCards();
  }
  render() {
    // fake card data
    return (
      <div>
        <h1>All Cards</h1>
        <CardGrid cards={this.props.card.all_cards} />
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
