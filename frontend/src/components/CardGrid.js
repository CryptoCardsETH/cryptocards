import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';
import Card from './Card.js';

class CardGrid extends Component {
  render() {
    let { cards, filter, type } = this.props;
    if (filter && filter.text && filter.text !== '') {
      cards = cards.filter(c => c.name.includes(filter.text));
    }
    if (filter && filter.sort) {
      if (filter.sort === 'id')
        cards.sort((a, b) => (a.id > b.id ? 1 : b.id > a.id ? -1 : 0));
      if (filter.sort === 'az')
        cards.sort((a, b) => (a.name > b.name ? 1 : b.name > a.name ? -1 : 0));
      if (filter.sort === 'za')
        cards.sort((a, b) => (a.name < b.name ? 1 : b.name < a.name ? -1 : 0));
    }
    return (
      <div className="row">
        {cards.map((card, index) => {
          return <Card key={index} card={card} index={index} type={type} />;
        })}
      </div>
    );
  }
}

export default withRouter(CardGrid);
