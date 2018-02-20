import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import {
  editCardDetail,
  fetchCardDetail,
  saveCardDetail
} from '../actions/cards';
import { fetchMe } from '../actions/users';
import { Button } from 'reactstrap';

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
      user.authenticated &&
      cardDetail.user &&
      user.me.id === cardDetail.user.id;
    return (
      <div>
        <h1>Card Detail: #{this.state.cardId}</h1>
        <pre>{JSON.stringify(cardDetail, true, 2)}</pre>
        <div>my user: {user.me.id}</div>
        <div>
          {doesCurrentUserOwnCard ? (
            <div>
              You own this card!
              <br />
              Hide card from public lists and your profile?
              <input
                name="isHidden"
                type="checkbox"
                checked={cardDetail.hidden}
                onChange={e => {
                  this.props.editCardDetail(
                    this.state.cardId,
                    'hidden',
                    e.target.checked
                  );
                }}
              />
              <Button
                onClick={() => {
                  this.props.saveCardDetail(this.state.cardId);
                }}
              >
                Save Card Preferences
              </Button>
            </div>
          ) : null}
        </div>
      </div>
    );
  }
}
function mapStateToProps(state) {
  let { user, card } = state;
  return { user, card };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators(
    { fetchCardDetail, fetchMe, editCardDetail, saveCardDetail },
    dispatch
  );
};
export default connect(mapStateToProps, mapDispatchToProps)(CardDetail);
