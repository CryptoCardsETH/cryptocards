import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import {
  editCardDetail,
  fetchCardDetail,
  saveCardDetail,
  updateUserCards
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
        <div>
          <div className="row">
            <div className="col-md-6">
              <img
                className="card-img-top"
                src={cardDetail.imageURL}
                alt={cardDetail.name}
              />
            </div>
            <div className="col-md-6">
              <h1>{cardDetail.name}</h1>
              <hr />
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
                    <br />
                    <Button
                      onClick={() => {
                        this.props.saveCardDetail(this.state.cardId);
                      }}
                    >
                      Save Card Preferences
                    </Button>
                  </div>
                ) : (
                  <div>
                    <br />
                    <Button
                      onClick={e => {
                        this.props.updateUserCards(this.state.cardId);
                        e.preventDefault();
                      }}
                    >
                      Buy
                    </Button>
                  </div>
                )}
              </div>
              <br />
              <div>
                {cardDetail.attributes.length > 0 ? (
                  <div>
                    <h4>Attributes</h4>
                    <table className="table">
                      <tbody>
                        {cardDetail.attributes.map((attribute, index) => {
                          return (
                            <tr>
                              <td>{attribute.name}</td>
                              <td>{attribute.pivot.value}</td>
                            </tr>
                          );
                        })}
                      </tbody>
                    </table>
                  </div>
                ) : null}
              </div>
            </div>
          </div>
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
    {
      fetchCardDetail,
      fetchMe,
      editCardDetail,
      saveCardDetail,
      updateUserCards
    },
    dispatch
  );
};
export default connect(mapStateToProps, mapDispatchToProps)(CardDetail);
