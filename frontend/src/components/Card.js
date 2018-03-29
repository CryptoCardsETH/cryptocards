import React, { Component } from 'react';
import '../styles/App.css';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faExternalLinkAlt from '@fortawesome/fontawesome-free-solid/faExternalLinkAlt';
import 'animate.css';
import EtherPrice from './EtherPrice.js';
import { Link } from 'react-router-dom';
import classNames from 'classnames';
import { Button } from 'reactstrap';
class Card extends Component {
  render() {
    let { card, index, type, toggleCardSelection } = this.props;
    let listing;
    if (type === CARD_TYPE_MARKETPLACE) {
      listing = card;
      card = card.cards;
    }

    let { isSelected } = card;
    let style = {
      animationDelay: (index % 3) / 10 + 's'
    };
    let cardClass = classNames('card col-lg-3 col-md-5 m-2 animated fadeInUp', {
      'bg-secondary': isSelected
    });

    return (
      <div key={index} className={cardClass} style={style}>
        <div className="overlay-container">
          <img
            className="card-img-top"
            src={`http://via.placeholder.com/350?text=card+id+${card.id}`}
            alt={card.name}
          />
          <div className="overlay overlay-price">
            {type === CARD_TYPE_MARKETPLACE ? (
              <EtherPrice price={listing.price} />
            ) : null}
          </div>
          <div className="overlay overlay-background">
            <Link to={'/card/' + card.id} title="More Details">
              <FontAwesomeIcon
                icon={faExternalLinkAlt}
                className="details"
                size="2x"
              />
            </Link>
          </div>
        </div>
        <div className="card-body">
          <h5 className="card-title text-center">{card.name}</h5>
          {type === CARD_TYPE_COLLECTION ? (
            <div>
              <p className="card-text">
                <Button onClick={() => toggleCardSelection(card.id)}>
                  {isSelected ? 'deselect' : 'select'}
                </Button>
                <br />
                owner: {card.user ? card.user.nickname : 'n/a'}
              </p>
              <p className="card-text">hidden: {card.hidden ? 'yes' : 'no'}</p>
            </div>
          ) : null}
        </div>
      </div>
    );
  }
}

export const CARD_TYPE_MARKETPLACE = 'marketplace';
export const CARD_TYPE_COLLECTION = 'collection';

export default Card;
