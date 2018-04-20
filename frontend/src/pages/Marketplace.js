import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllListings } from '../actions/listings';
import { CARD_TYPE_MARKETPLACE } from '../components/Card';

class MarketplacePage extends Component {
  componentDidMount() {
    this.props.fetchAllListings();
  }
  render() {
    return (
      <div className="container">
        <br />
        <h3>Marketplace</h3>
        <br />
        <CardGrid
          cards={this.props.listing.all_listings}
          type={CARD_TYPE_MARKETPLACE}
        />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { listing: state.listing };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchAllListings }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(MarketplacePage);
