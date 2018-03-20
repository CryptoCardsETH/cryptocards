import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchMe, fetchUserDetail } from '../actions/users';
import { buildProfileURL } from '../actions';
import { CARD_TYPE_COLLECTION } from '../components/Card';
import CardFilterSort, {
  FILTER_SORT_PRESET_BASE,
  FILTER_SORT_PRESET_FULL
} from '../components/CardFilterSort';

class UserDetail extends Component {
  constructor(props) {
    super(props);
    this.state = { userIdOrNickname: null };
  }
  componentDidMount() {
    this.props.fetchMe();
  }
  componentWillReceiveProps(nextProps) {
    let userIdOrNicknameFromRouter = nextProps.match.params.id;
    if (userIdOrNicknameFromRouter !== this.state.userIdOrNickname) {
      this.props.fetchUserDetail(userIdOrNicknameFromRouter);
      this.setState({ userIdOrNickname: userIdOrNicknameFromRouter });
    }
  }
  render() {
    let { user } = this.props;
    let { userIdOrNickname } = this.state;
    let userDetail = user.user_detail[userIdOrNickname];
    if (!userDetail) return <h1>Loading</h1>;
    let userId = userDetail.user.id;
    let isViewingMyProfile = user.authenticated && user.me.id === userId;
    return (
      <div>
        <h1>
          {isViewingMyProfile ? 'My Collection' : `Viewing User #${userId}`}{' '}
          &nbsp;
          <small>
            {userDetail.user.nickname ? userDetail.user.nickname : ''}
          </small>
        </h1>
        Canonical profile URL: {buildProfileURL(userDetail.user, true)}
        <hr />
        <pre>{JSON.stringify({ user, userDetail }, true, 2)}</pre>
        <hr />
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
  return bindActionCreators({ fetchMe, fetchUserDetail }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(UserDetail);
