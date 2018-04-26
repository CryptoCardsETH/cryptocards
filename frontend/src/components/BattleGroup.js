import React from 'react';
import Card, { CARD_TYPE_COLLECTION } from './Card';
import BattleCreator from './BattleCreator';
const BattleGroup = ({ group }) => {
  let { user, group_cards } = group;
  //todo: disabled for in progress battle and stuff
  return (
    <div style={{ border: '1px solid #d8d8d8' }}>
      USER: {user.nickname || user.id}
      <br />
      Group ID: {group.id} <br />
      <BattleCreator battleGroupId={group.token_id} />
      <div className="row">
        {group_cards.map((gc, x) => {
          let { card } = gc;
          return <Card key={x} card={card || gc} type={CARD_TYPE_COLLECTION} />;
        })}
      </div>
    </div>
  );
};
export default BattleGroup;
