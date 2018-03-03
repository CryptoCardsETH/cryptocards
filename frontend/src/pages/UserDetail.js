import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faUserPlus from '@fortawesome/fontawesome-free-solid/faUserPlus';
import faCheck from '@fortawesome/fontawesome-free-solid/faCheck';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { addFriend, fetchMe, fetchUserDetail } from '../actions/users';
import { CARD_TYPE_COLLECTION } from '../components/Card';
import CardFilterSort, {
  FILTER_SORT_PRESET_BASE,
  FILTER_SORT_PRESET_FULL
} from '../components/CardFilterSort';
import '../styles/App.css';

class UserDetail extends Component {
  constructor(props) {
    super(props);
    this.state = { userId: null };
  }
  componentDidMount() {
    this.props.fetchMe();
  }
  componentWillReceiveProps(nextProps) {
    let userIdFromRouter = parseInt(nextProps.match.params.id, 10);
    if (userIdFromRouter !== this.state.userId) {
      this.props.fetchUserDetail(userIdFromRouter);
      this.setState({ userId: userIdFromRouter });
    }
  }
  render() {
    let { user } = this.props;
    let { userId } = this.state;
    let userDetail = user.user_detail[userId];
    if (!userDetail) return <h1>Loading</h1>;
    let isViewingMyProfile = user.authenticated && user.me.id === userId;
    return (
      <div>
        <h1>
          {isViewingMyProfile ? 'My Collection' : `Viewing User #${userId}`}
        </h1>
        {!isViewingMyProfile ? (
          <div className="float-right">
            {!userDetail.isFriend ? (
              <button
                className="btn btn-success"
                onClick={() => {
                  this.props.addFriend(userId);
                }}
              >
                <FontAwesomeIcon icon={faUserPlus} /> Add Friend
              </button>
            ) : (
              <div className="btn bg-primary text-white">
                <FontAwesomeIcon icon={faCheck} /> Friends
              </div>
            )}
          </div>
        ) : null}
        <CardFilterSort
          filterSortKey="mycards"
          sortTypes={
            isViewingMyProfile
              ? FILTER_SORT_PRESET_FULL
              : FILTER_SORT_PRESET_BASE
          }
        />
        <CardGrid
          cards={userDetail.cards}
          filter={this.props.card.filters['mycards']}
          type={CARD_TYPE_COLLECTION}
        />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { user: state.user, card: state.card };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchMe, fetchUserDetail, addFriend }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(UserDetail);
