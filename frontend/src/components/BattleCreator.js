import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { setCardFilterText, setCardSortOption } from '../actions/cards';
import PropTypes from 'prop-types';
import { CONTRACT_NAME_BATTLEQUEUE } from './../contracts';
import {
  isReadyForContract,
  waitForTxToBeMined,
  getContractInstanceByName
} from '../selectors';
import { Button } from 'reactstrap';
class BattleCreator extends React.Component {
  doContract = async (senderAddress, ci, battleGroupId) => {
    //todo: web3 check
    let contractInstance = await ci;
    contractInstance
      .joinQueue(battleGroupId, {
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
              battleGroupId
            )
          }
        >
          {ready ? 'send battle group to queue!' : 'loading...'} {battleGroupId}
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
    contractInstance: getContractInstanceByName(
      state,
      CONTRACT_NAME_BATTLEQUEUE
    )
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
