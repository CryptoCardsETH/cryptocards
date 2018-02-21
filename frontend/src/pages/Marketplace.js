import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllListings } from '../actions/listings';

class MarketplacePage extends Component {
  componentDidMount() {
    this.props.fetchAllListings();
  }
  render() {
    return (
      <div className="marketplace">
        <h1>Marketplace</h1>
        <CardGrid cards={this.props.listing.all_listings} type="marketplace" />
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
