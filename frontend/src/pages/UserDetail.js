import React, { Component } from 'react';
import CardGrid from '../components/CardGrid';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchMe, follow, fetchUserDetail } from '../actions/users';
import { toggleCardSelection } from '../actions/cards';
import { buildProfileURL } from '../actions';
import { CARD_TYPE_COLLECTION } from '../components/Card';
import BattleGroup from '../components/BattleGroup';
import BattleTable from '../components/BattleTable';
import { Redirect } from 'react-router';
import CardFilterSort, {
  FILTER_SORT_PRESET_BASE,
  FILTER_SORT_PRESET_FULL
} from '../components/CardFilterSort';
import BattleGroupCreator from '../components/BattleGroupCreator';
import UserDetailFollowing from '../components/UserDetailFollowing';
import CopyToClipboardButton from '../components/CopyToClipboardButton';
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
    let { user, toggleCardSelection, card } = this.props;
    let { userIdOrNickname } = this.state;
    let userDetail = user.user_detail[userIdOrNickname];
    // show loading message if user data hasn't loaded yet.
    if (!userDetail) return <h1>Loading</h1>;

    let userCardsWithSelection = userDetail.cards.map(x => ({
      ...x,
      isSelected: card.selectedCardIDs.includes(x.id)
    }));

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
        <hr />
        {!isViewingMyProfile ? (
          <div className="float-right">
            <UserDetailFollowing
              follow={() => this.props.follow(userId)}
              isFollowing={userDetail.isFollowing}
            />
            <CopyToClipboardButton
              copyText={buildProfileURL(userDetail.user, true)}
              buttonText="Copy Profile URL"
            />
          </div>
        ) : (
          <div>
            <h4>make a battle group!</h4>
            <BattleGroupCreator cardIds={card.selectedCardIDs} />
          </div>
        )}
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
          toggleCardSelection={toggleCardSelection}
          cards={userCardsWithSelection}
          filter={this.props.card.filters['mycards']}
          type={CARD_TYPE_COLLECTION}
        />
        <hr />
        <h1>Battle Groups</h1>
        {userDetail.battleGroups.map(bg => (
          <BattleGroup key={bg.id} group={bg} />
        ))}
        <h2>Battles</h2>
        <BattleTable battles={userDetail.battles} />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { user: state.user, card: state.card };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators(
    { fetchMe, fetchUserDetail, follow, toggleCardSelection },
    dispatch
  );
};
export default connect(mapStateToProps, mapDispatchToProps)(UserDetail);
