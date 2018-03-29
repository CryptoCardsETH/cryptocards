import CompiledContracts from './compiled_contracts';

export const ABIs = {};

export const BattleGroupABI =
  CompiledContracts['BattleGroups.sol:BattleGroups']['interface'];
export const CONTRACT_NAME_BATTLEGROUPS = 'BattleGroups';
ABIs[CONTRACT_NAME_BATTLEGROUPS] = BattleGroupABI;
