import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faUserPlus from '@fortawesome/fontawesome-free-solid/faUserPlus';
import faCheck from '@fortawesome/fontawesome-free-solid/faCheck';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchMe, follow, fetchUserDetail } from '../actions/users';
import { buildProfileURL } from '../actions';
import { CARD_TYPE_COLLECTION } from '../components/Card';
import { Redirect } from 'react-router';
import { CopyToClipboard } from 'react-copy-to-clipboard';
import CardFilterSort, {
  FILTER_SORT_PRESET_BASE,
  FILTER_SORT_PRESET_FULL
} from '../components/CardFilterSort';
import { Button } from 'reactstrap';
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
    // show loading message if user data hasn't loaded yet.
    if (!userDetail) return <h1>Loading</h1>;
    let userId = userDetail.user.id;
    let isViewingMyProfile = user.authenticated && user.me.id === userId;
    // redirect to username if accessed via ID
    if (parseInt(userIdOrNickname, 10) === userId && userDetail.user.nickname)
      return <Redirect to={buildProfileURL(userDetail.user)} />;
    return (
      <div>
        <h1>
          {isViewingMyProfile ? 'My Collection' : `Viewing User #${userId}`}{' '}
          &nbsp;
          <small>
            {userDetail.user.nickname ? userDetail.user.nickname : ''}
          </small>
        </h1>
        {!isViewingMyProfile ? (
          <div className="float-right">
            {!userDetail.isFollowing ? (
              <button
                className="btn btn-success"
                onClick={() => {
                  this.props.follow(userId);
                }}
              >
                <FontAwesomeIcon icon={faUserPlus} /> Follow
              </button>
            ) : (
              <div className="btn bg-primary text-white">
                <FontAwesomeIcon icon={faCheck} /> Following
              </div>
            )}{' '}
            <CopyToClipboard text={buildProfileURL(userDetail.user, true)}>
              <Button color="primary">copy profile URL</Button>
            </CopyToClipboard>
          </div>
        ) : null}
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
  return bindActionCreators({ fetchMe, fetchUserDetail, follow }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(UserDetail);
