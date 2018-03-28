import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllBattles } from '../actions/battles';
import BattleTable from '../components/BattleTable';
class AllBattles extends Component {
  componentDidMount() {
    this.props.fetchAllBattles();
  }
  render() {
    return (
      <div>
        <h1>All Battles</h1>
        <BattleTable battles={this.props.battle.all_battles} />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { battle: state.battle };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchAllBattles }, dispatch);
};
export default connect(mapStateToProps, mapDispatchToProps)(AllBattles);
