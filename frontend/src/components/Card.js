import React, { Component } from 'react';
import { removeCard } from '../actions/cards';
import '../styles/App.css';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faExternalLinkAlt from '@fortawesome/fontawesome-free-solid/faExternalLinkAlt';
import 'animate.css';
import EtherPrice from './EtherPrice.js';
import { Link } from 'react-router-dom';
import classNames from 'classnames';
import {
  Card as BootstrapCard,
  CardBody,
  CardTitle,
  CardText,
  UncontrolledTooltip,
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter
} from 'reactstrap';
class Card extends Component {
  constructor(props) {
    super(props);
    this.state = {
      modal: false,
      cardId: null
    };

    this.toggle = this.toggle.bind(this);
  }

  toggle() {
    this.setState({
      modal: !this.state.modal
    });
  }

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
    let cardClass = classNames('col-lg-3 col-md-5 m-2 animated fadeInUp', {
      'bg-secondary': isSelected
    });

    return (
      <BootstrapCard key={index} className={cardClass} style={style}>
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
        <CardBody>
          <CardTitle className="text-center">
            {card.name}
            <div className="float-right">
              <button
                type="button"
                class="close"
                aria-label="Close"
                id={`popover_${index}`}
                onClick={this.toggle}
              >
                <span aria-hidden="true">&times;</span>
              </button>
              <UncontrolledTooltip placement="top" target={`popover_${index}`}>
                Delete Card
              </UncontrolledTooltip>
              <Modal
                isOpen={this.state.modal}
                toggle={this.toggle}
                className={this.props.className}
              >
                <ModalHeader toggle={this.toggle}>My Collection</ModalHeader>
                <ModalBody>
                  Are you sure you want to remove card {card.id} from your
                  collection?
                </ModalBody>
                <ModalFooter>
                  <button color="primary" onClick={removeCard(card.id)}>
                    Yes
                  </button>{' '}
                  <button color="secondary" onClick={this.toggle}>
                    Cancel
                  </button>
                </ModalFooter>
              </Modal>
            </div>
          </CardTitle>
          {type === CARD_TYPE_COLLECTION ? (
            <CardText>
              <p>owner: {card.user ? card.user.nickname : 'n/a'}</p>
              <p>hidden: {card.hidden ? 'yes' : 'no'}</p>
              <div className="text-center">
                <button
                  onClick={() => toggleCardSelection(card.id)}
                  className="btn-blue m-*-auto"
                >
                  {isSelected ? 'deselect' : 'select'}
                </button>
              </div>
            </CardText>
          ) : null}
        </CardBody>
      </BootstrapCard>
    );
  }
}
Card.defaultProps = {
  toggleCardSelection: null
};

export const CARD_TYPE_MARKETPLACE = 'marketplace';
export const CARD_TYPE_COLLECTION = 'collection';

export default Card;
