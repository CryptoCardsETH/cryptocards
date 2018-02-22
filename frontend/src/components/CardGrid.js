import React, { Component } from 'react';
import '../styles/App.css';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faExternalLinkAlt from '@fortawesome/fontawesome-free-solid/faExternalLinkAlt';
import 'animate.css';
import { withRouter } from 'react-router-dom';
import { Button } from 'reactstrap';

class CardGrid extends Component {
  render() {
    let { cards, filter } = this.props;
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
                <div class="float-right">
                  <Button
                    onClick={() => {
                      //call functions from card in actions
                    }}
                  >
                    Buy
                  </Button>
                </div>
              </div>
            </div>
          );
        })}
      </div>
    );
  }
}

export default withRouter(CardGrid);
