import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchMyCards } from '../actions/users';
import CardFilterSort from '../components/CardFilterSort';

class CollectionPage extends Component {
  componentDidMount() {
    this.props.fetchMyCards();
  }
  render() {
    // fake card data
    return (
      <div>
        <h1>My Collection</h1>
        <CardFilterSort filterSortKey="mycards" />
        <CardGrid
          cards={this.props.user.cards}
          filter={this.props.card.filters['mycards']}
          type="collection"
        />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { user: state.user, card: state.card };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchMyCards }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(CollectionPage);
