import CompiledContracts from './compiled_contracts';

export const ABIs = {};

export const BattleGroupABI =
  CompiledContracts['BattleGroups.sol:BattleGroups']['interface'];
export const BattleABI = CompiledContracts['Battles.sol:Battles']['interface'];

export const CONTRACT_NAME_BATTLEGROUPS = 'BattleGroups';
export const CONTRACT_NAME_BATTLES = 'Battles';

ABIs[CONTRACT_NAME_BATTLEGROUPS] = BattleGroupABI;
ABIs[CONTRACT_NAME_BATTLES] = BattleABI;
