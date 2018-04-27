import CompiledContracts from './compiled_contracts';

export const ABIs = {};

export const BattleGroupABI =
  CompiledContracts['battles/BattleGroups.sol:BattleGroups']['interface'];
export const BattleABI =
  CompiledContracts['battles/Battles.sol:Battles']['interface'];
export const BattleQueueABI =
  CompiledContracts['battles/BattleQueue.sol:BattleQueue']['interface'];

export const CoreABI =
  CompiledContracts['CryptoCardsCore.sol:CryptoCardsCore']['interface'];

export const CONTRACT_NAME_BATTLEGROUPS = 'BattleGroups';
export const CONTRACT_NAME_BATTLEQUEUE = 'BattleQueue';
export const CONTRACT_NAME_BATTLES = 'Battles';
export const CONTRACT_NAME_CORE = 'CryptoCardsCore';

ABIs[CONTRACT_NAME_BATTLEGROUPS] = BattleGroupABI;
ABIs[CONTRACT_NAME_BATTLEQUEUE] = BattleQueueABI;
ABIs[CONTRACT_NAME_BATTLES] = BattleABI;
ABIs[CONTRACT_NAME_CORE] = CoreABI;
