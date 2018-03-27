import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllBattles } from '../actions/battles';
import { Link } from 'react-router-dom';
class AllBattles extends Component {
  componentDidMount() {
    this.props.fetchAllBattles();
  }
  render() {
    // fake battle data
    return (
      <div>
        <h1>All Battles</h1>
        {this.props.battle.all_battles.map(b => {
          let { id } = b;
          return (
            <div>
              <h3>#{id}</h3>
              <h3>Opponent 1</h3>
              <pre>{JSON.stringify(b.group1, true, 2)}</pre>

              <h3>Opponent 2</h3>
              <pre>{JSON.stringify(b.group2, true, 2)}</pre>
            </div>
          );
        })}
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
