import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchAllBattles } from '../actions/battles';
import BattleGroup from '../components/BattleGroup';
import { Table } from 'reactstrap';
class AllBattles extends Component {
  componentDidMount() {
    this.props.fetchAllBattles();
  }
  render() {
    return (
      <div>
        <h1>All Battles</h1>
        <Table>
          <thead>
            <tr>
              <th>#</th>
              <th>state</th>
              <th>Opponent 1</th>
              <th>Opponent 2</th>
            </tr>
          </thead>
          <tbody>
            {this.props.battle.all_battles.map(b => {
              let { id, state } = b;
              return (
                <tr key={id}>
                  <td>{id}</td>
                  <td>{state}</td>
                  <td>
                    <BattleGroup group={b.group1} />
                  </td>
                  <td>
                    <BattleGroup group={b.group2} />
                  </td>
                </tr>
              );
            })}
          </tbody>
        </Table>
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
