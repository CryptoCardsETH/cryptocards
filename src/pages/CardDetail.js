import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { fetchCardDetail } from '../actions/cards';
import { fetchMe } from '../actions/users';

class CardDetail extends Component {
  constructor(props) {
    super(props);
    this.state = { cardId: null };
  }
  componentDidMount() {
    this.props.fetchMe();
  }
  componentWillReceiveProps(nextProps) {
    let cardIdFromRouter = nextProps.match.params.id;
    if (cardIdFromRouter !== this.state.cardId) {
      this.props.fetchCardDetail(cardIdFromRouter);
      this.setState({ cardId: cardIdFromRouter });
    }
  }
  render() {
    let { user, card } = this.props;
    let cardDetail = card.card_detail[this.state.cardId];
    if (!cardDetail) return <h1>Loading</h1>;

    let doesCurrentUserOwnCard =
      user.authenticated && user.me.id === cardDetail.user.id;
    return (
      <div>
        <h1>Card Detail: #{this.state.cardId}</h1>
        <pre>{JSON.stringify(cardDetail, true, 2)}</pre>
        <div>my user: {user.me.id}</div>
        <div>{doesCurrentUserOwnCard ? 'You own this card!' : null}</div>
      </div>
    );
  }
}
function mapStateToProps(state) {
  let { user, card } = state;
  return { user, card };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchCardDetail, fetchMe }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(CardDetail);
