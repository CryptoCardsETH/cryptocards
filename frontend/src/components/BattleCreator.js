import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { setCardFilterText, setCardSortOption } from '../actions/cards';
import PropTypes from 'prop-types';
import { CONTRACT_NAME_BATTLES } from './../contracts';
import {
  isReadyForContract,
  waitForTxToBeMined,
  getContractInstanceByName
} from '../selectors';
import { Button } from 'reactstrap';
class BattleCreator extends React.Component {
  doContract = async (senderAddress, ci, test1, test2) => {
    //todo: web3 check
    //TODO: this needs to be rewritten when the Battles.sol contract is finished, right now it just sends the specified group as both sides for the battle.
    let contractInstance = await ci;
    contractInstance
      .createBattle(test1, test2, {
        from: senderAddress
      })
      .then(txHash => {
        waitForTxToBeMined(txHash);
      })
      .catch(console.error);
  };

  render() {
    let { ready, contractInstance, battleGroupId } = this.props;
    return (
      <div>
        <Button
          disabled={!ready}
          onClick={() =>
            this.doContract(
              this.props.user.main_address,
              contractInstance,
              battleGroupId,
              battleGroupId
            )
          }
        >
          {ready ? 'send battle to contract!' : 'loading...'}
        </Button>
      </div>
    );
  }
}

BattleCreator.propTypes = {
  battleGroupId: PropTypes.number.isRequired
};

function mapStateToProps(state) {
  let { card, user } = state;
  return {
    card,
    user,
    ready: isReadyForContract(state),
    contractInstance: getContractInstanceByName(state, CONTRACT_NAME_BATTLES)
  };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators(
    {
      setCardFilterText,
      setCardSortOption
    },
    dispatch
  );
};
export default connect(mapStateToProps, mapDispatchToProps)(BattleCreator);
