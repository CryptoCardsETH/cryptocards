import React from 'react';
import { Table } from 'reactstrap';
import BattleGroup from './BattleGroup';
const BattleTable = ({ battles }) => {
  return (
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
        {battles.map(b => {
          let { id, groupwinner } = b;
          return (
            <tr key={id}>
              <td>{id}</td>
              <td>
                {groupwinner
                  ? `winning group id: ${groupwinner.id}`
                  : 'battle in progress'}
              </td>
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
  );
};
export default BattleTable;
