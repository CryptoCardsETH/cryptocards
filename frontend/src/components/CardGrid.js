import React, { Component } from 'react';
import '../styles/App.css';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faExternalLinkAlt from '@fortawesome/fontawesome-free-solid/faExternalLinkAlt';
import 'animate.css';
import { withRouter } from 'react-router-dom';
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
    let { cards, filter } = this.props;
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
          let style = {
            animationDelay: (index % 3) / 10 + 's'
          };
          return (
            <div
              key={index}
              className="card col-lg-3 col-md-5 m-2 animated fadeInUp"
              style={style}
              onClick={() => {
                this.props.history.push('/card/' + card.id);
              }}
            >
              <div className="overlay-container">
                <img
                  className="card-img-top"
                  src={`http://via.placeholder.com/350?text=card+id+${card.id}`}
                  alt={card.name}
                />
                <div className="overlay">
                  <a href="/collection" title="More Details">
                    <FontAwesomeIcon
                      icon={faExternalLinkAlt}
                      className="details"
                      size="2x"
                    />
                  </a>
                </div>
              </div>
              <div className="card-body">
                <h5 className="card-title text-center">{card.name}</h5>
                <p className="card-text">
                  owner: {card.user ? card.user.nickname : 'n/a'}
                </p>
                <p className="card-text">
                  hidden: {card.hidden ? 'yes' : 'no'}
                </p>
              </div>
            </div>
          );
        })}
      </div>
    );
  }
}

export default withRouter(CardGrid);
