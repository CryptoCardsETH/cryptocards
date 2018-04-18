import React from 'react';
import Card, { CARD_TYPE_COLLECTION } from './Card';
const BattleGroup = ({ group }) => {
  let { user, group_cards } = group;
  return (
    <div>
      USER: {user.nickname || user.id}
      <br />
      Group ID: {group.id} <br />
      <div className="row">
        {group_cards.map((gc, x) => (
          <Card key={x} card={gc} type={CARD_TYPE_COLLECTION} />
        ))}
      </div>
    </div>
  );
};
export default BattleGroup;
