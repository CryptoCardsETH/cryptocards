import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchMyCards } from '../actions/users';

class CollectionPage extends Component {
  componentDidMount() {
    this.props.fetchMyCards();
  }
  render() {
    // fake card data
    return (
      <div>
        <h1>My Collection</h1>
        <CardGrid cards={this.props.user.cards} />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { user: state.user };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchMyCards }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(CollectionPage);
