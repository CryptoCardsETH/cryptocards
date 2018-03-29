import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';
import Card from './Card.js';
import {
  FILTER_SORT_ID,
  FILTER_SORT_AZ,
  FILTER_SORT_ZA,
  FILTER_SORT_PUBLIC,
  FILTER_SORT_HIDDEN,
  FILTER_SORT_POPULARITY,
  FILTER_SORT_PRICE,
  FILTER_SORT_ATTRIBUTES
} from './CardFilterSort';

class CardGrid extends Component {
  render() {
    let { cards, filter, type, toggleCardSelection } = this.props;
    if (filter && filter.text && filter.text !== '') {
      cards = cards.filter(c => c.name.includes(filter.text));
    }
    if (filter && filter.sort) {
      if (filter.sort === FILTER_SORT_ID.value)
        cards.sort((a, b) => (a.id > b.id ? 1 : b.id > a.id ? -1 : 0));
      if (filter.sort === FILTER_SORT_AZ.value)
        cards.sort((a, b) => (a.name > b.name ? 1 : b.name > a.name ? -1 : 0));
      if (filter.sort === FILTER_SORT_ZA.value)
        cards.sort((a, b) => (a.name < b.name ? 1 : b.name < a.name ? -1 : 0));
      if (filter.sort === FILTER_SORT_PUBLIC.value)
        cards = cards.filter(c => !c.hidden);
      if (filter.sort === FILTER_SORT_HIDDEN.value)
        cards = cards.filter(c => c.hidden);
      if (filter.sort === FILTER_SORT_POPULARITY.value)
        console.log('TODO:IMPLEMENT');
      if (filter.sort === FILTER_SORT_PRICE.value)
        console.log('TODO:IMPLEMENT');
      if (filter.sort === FILTER_SORT_ATTRIBUTES.value)
        console.log('TODO:IMPLEMENT');
    }

    return (
      <div className="row">
        {cards.map((card, index) => {
          return (
            <Card
              key={index}
              card={card}
              index={index}
              type={type}
              toggleCardSelection={toggleCardSelection}
            />
          );
        })}
      </div>
    );
  }
}

export default withRouter(CardGrid);
